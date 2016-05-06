package gedcom

import (
	"errors"
	"io"
	"strings"

	"github.com/MJKWoolnough/parser"
)

const (
	alpha       = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_"
	delim       = " "
	digit       = "0123456789"
	alphanum    = alpha + digit
	otherchar   = "!\"$%&'()*+,-./:;<=>?[\\]^`{|}~" // + "\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8a\x8b\x8c\x8d\x8e\x8f\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9a\x9b\x9c\x9d\x9e\x9f\xa0\xa1\xa2\xa3\xa4\xa5\xa6\xa7\xa8\xa9\xaa\xab\xac\xad\xae\xaf\xb0\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe\xbf\xc0\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf\xd0\xd1\xd2\xd3\xd4\xd5\xd6\xd7\xd8\xd9\xda\xdb\xdc\xdd\xde\xdf\xe0\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xeb\xec\xed\xee\xef\xf0\xf1\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xfb\xfc\xfd\xfe" // ANSEL
	nonAt       = alpha + digit + otherchar + " #"
	terminators = "\r\n"
	anyChar     = alpha + digit + otherchar + "# @"
	levelIgnore = terminators + delim + "	"
	invalidchar = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\xab\xac\xad\x1e\x1f\x7f" + terminators + "@"
)

const (
	tokenLevel parser.TokenType = iota
	tokenXref
	tokenTag
	tokenPointer
	tokenLine
	tokenEndLine
)

type tokeniser struct {
	parser.Parser
	options
}

func newTokeniser(r io.Reader, o options) *tokeniser {
	t := &tokeniser{
		Parser:  parser.NewReaderParser(r),
		options: o,
	}
	t.State = t.level
	return t
}

func (t *tokeniser) level() (parser.Token, parser.StateFn) {
	t.AcceptRun(levelIgnore)
	t.Get()
	if t.Peek() == -1 {
		return t.Done()
	}
	if !t.Accept(digit) {
		t.Err = ErrInvalidLevel
		return t.Error()
	}
	t.AcceptRun(digit)
	if !t.Accept(delim) {
		t.Err = ErrMissingDelim
		return t.Error()
	}
	return parser.Token{
		tokenLevel,
		strings.TrimSpace(t.Get()),
	}, t.optionalXrefID
}

func (t *tokeniser) optionalXrefID() (parser.Token, parser.StateFn) {
	if t.Peek() == '@' {
		return t.xrefID()
	}
	return t.tag()
}

func (t *tokeniser) readPointer() (string, error) {
	if !t.Accept(alphanum) {
		return "", ErrInvalidPointer
	}
	t.AcceptRun(nonAt)
	if !t.Accept("@") {
		return "", ErrInvalidPointer
	}
	pointer := t.Get()
	return pointer[1 : len(pointer)-1], nil
}

func (t *tokeniser) xrefID() (parser.Token, parser.StateFn) {
	t.Accept("@")
	pointer, err := t.readPointer()
	if err != nil {
		t.Err = err
		return t.Error()
	}
	if !t.Accept(delim) {
		t.Err = ErrMissingDelim
		return t.Error()
	}
	t.Get()
	return parser.Token{
		tokenXref,
		strings.Trim(pointer, "@"),
	}, t.tag
}

func (t *tokeniser) tag() (parser.Token, parser.StateFn) {
	if !t.Accept(alphanum) {
		t.Err = ErrInvalidTag
		return t.Error()
	}
	t.AcceptRun(alphanum)
	tag := t.Get()
	if t.Accept(delim) {
		t.Get()
		return parser.Token{
			tokenTag,
			tag,
		}, t.lineValue
	}
	next := t.endLine
	if t.Peek() == -1 {
		next = t.Done
	} else {
		if !t.Accept(terminators) {
			t.Err = ErrInvalidTag
			return t.Error()
		}
		t.AcceptRun(terminators)
		t.Get()
	}
	return parser.Token{
		tokenTag,
		tag,
	}, next
}

func (t *tokeniser) endLine() (parser.Token, parser.StateFn) {
	return parser.Token{
		tokenEndLine,
		"",
	}, t.level
}

func (t *tokeniser) lineValue() (parser.Token, parser.StateFn) {
	if t.Peek() == '@' {
		t.Accept("@")
		if t.Peek() != '@' {
			pointer, err := t.readPointer()
			if err != nil {
				if !t.allowInvalidEscape {
					t.Err = err
					return t.Error()
				}
			} else {
				t.AcceptRun(terminators)
				t.Get()
				return parser.Token{
					tokenPointer,
					pointer,
				}, t.level
			}
		}
	}
	next := t.level
	for {
		for {
			if t.allowUnknownCharset && t.Except(invalidchar) {
				t.ExceptRun(invalidchar)
			} else if t.Accept(nonAt) {
				t.AcceptRun(nonAt)
			} else if t.Accept("@") {
				if t.allowInvalidEscape || t.Accept("@") {
					continue
				}
				if !t.Accept("#") {
					t.Err = ErrBadEscape
					return t.Error()
				}
				if t.allowUnknownCharset {
					t.ExceptRun(invalidchar)
				} else {
					t.AcceptRun(nonAt)
				}
				if !t.Accept("@") {
					t.Err = ErrBadEscape
					return t.Error()
				}
			} else {
				break
			}
		}
		p := t.Peek()
		if p == -1 {
			next = t.Done
			break
		}
		if strings.ContainsRune(terminators, p) {
			t.AcceptRun(terminators)
			if !t.allowTerminatorsInValue {
				break
			}
			p = t.Peek()
			if p == -1 {
				next = t.Done
				break
			}
			if strings.ContainsRune(digit, p) {
				break
			}
		} else {
			if !t.allowInvalidChars {
				t.Err = ErrBadChar
				return t.Error()
			}
			t.Except("")
		}
	}
	return parser.Token{
		tokenLine,
		strings.TrimSpace(t.Get()),
	}, next
}

// Errors
var (
	ErrInvalidLevel   = errors.New("invalid level num")
	ErrMissingDelim   = errors.New("missing delminitator")
	ErrInvalidPointer = errors.New("invalid pointer string")
	ErrInvalidTag     = errors.New("invalid tag")
	ErrBadEscape      = errors.New("bad escape sequence")
	ErrBadChar        = errors.New("bad character")
)
