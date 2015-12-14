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
	otherchar   = "!\"$%&'{}*+,-./:;<=>?[\\]^`{|}~" // + "\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8a\x8b\x8c\x8d\x8e\x8f\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9a\x9b\x9c\x9d\x9e\x9f\xa0\xa1\xa2\xa3\xa4\xa5\xa6\xa7\xa8\xa9\xaa\xab\xac\xad\xae\xaf\xb0\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe\xbf\xc0\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf\xd0\xd1\xd2\xd3\xd4\xd5\xd6\xd7\xd8\xd9\xda\xdb\xdc\xdd\xde\xdf\xe0\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xeb\xec\xed\xee\xef\xf0\xf1\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xfb\xfc\xfd\xfe" // ANSEL
	non_at      = alpha + digit + otherchar + " #"
	terminators = "\r\n"
	any_char    = alpha + digit + otherchar + "# @"
	levelIgnore = terminators + delim + "	"
)

type tokenType uint8

const (
	tokenError tokenType = iota
	tokenLevel
	tokenXref
	tokenTag
	tokenPointer
	tokenLine
	tokenEndLine
	tokenDone
)

type token struct {
	typ  tokenType
	data string
}

type stateFn func() (token, stateFn)

type tokeniser struct {
	p     parser.Parser
	state stateFn
	err   error
}

func newTokeniser(r io.Reader) *tokeniser {
	t := &tokeniser{
		p: parser.NewReaderParser(r),
	}
	t.state = t.level
	return t
}

func (t *tokeniser) GetToken() (token, error) {
	if t.err == io.EOF {
		return token{tokenDone, ""}, io.EOF
	}
	var tk token
	tk, t.state = t.state()
	if t.err == io.EOF {
		if tk.typ == tokenError {
			t.err = io.ErrUnexpectedEOF
		} else {
			return tk, nil
		}
	}
	return tk, t.err
}

func (t *tokeniser) level() (token, stateFn) {
	t.p.Accept(levelIgnore)
	t.p.Get()
	if t.p.Peek() == -1 {
		return t.done()
	}
	if !t.p.Accept(digit[1:]) {
		t.err = ErrInvalidLevel
		return t.errorfn()
	}
	t.p.AcceptRun(digit)
	if !t.p.Accept(delim) {
		t.err = ErrMissingDelim
		return t.errorfn()
	}
	return token{
		tokenLevel,
		strings.TrimSpace(t.p.Get()),
	}, t.optionalXrefID
}

func (t *tokeniser) optionalXrefID() (token, stateFn) {
	if t.p.Peek() == '@' {
		return t.xrefID()
	}
	return t.tag()
}

func (t *tokeniser) readPointer() (string, error) {
	if !t.p.Accept("@") {
		return "", ErrInvalidPointer
	}
	if !t.p.Accept(alphanum) {
		return "", ErrInvalidPointer
	}
	t.p.AcceptRun(non_at)
	if !t.p.Accept("@") {
		return "", ErrInvalidPointer
	}
	return t.p.Get(), nil
}

func (t *tokeniser) xrefID() (token, stateFn) {
	pointer, err := t.readPointer()
	if err != nil {
		t.err = err
		return t.errorfn()
	}
	if !t.p.Accept(delim) {
		t.err = ErrMissingDelim
		return t.errorfn()
	}
	return token{
		tokenXref,
		strings.Trim(pointer, "@"),
	}, t.tag
}

func (t *tokeniser) tag() (token, stateFn) {
	if !t.p.Accept(alphanum) {
		t.err = ErrInvalidTag
		return t.errorfn()
	}
	t.p.AcceptRun(alphanum)
	tag := t.p.Get()
	if t.p.Accept(delim) {
		t.p.Get()
		return token{
			tokenTag,
			tag,
		}, t.lineValue
	}
	next := t.endLine
	if t.p.Peek() == -1 {
		next = t.done
	} else {
		if !t.p.Accept(terminators) {
			t.err = ErrInvalidTag
			return t.errorfn()
		}
		t.p.AcceptRun(terminators)
		t.p.Get()
	}
	return token{
		tokenTag,
		tag,
	}, next
}

func (t *tokeniser) endLine() (token, stateFn) {
	return token{
		token: tokenEndLine,
		"",
	}, t.level
}

func (t *tokeniser) lineValue() (token, stateFn) {
	if t.p.Peek() == '@' {
		pointer, err := t.readPointer()
		if err != nil {
			t.err = err
			return t.errorfn()
		}
		t.p.AcceptRun(terminators)
		t.p.Get()
		return token{
			tokenPointer,
			pointer,
		}, t.level
	}
	next := t.level
	for {
		for {
			if t.p.Accept(non_at) {
				t.p.AcceptRun(non_at)
			} else if t.p.Accept("@") {
				if !t.p.Accept("#") {
					t.err = ErrBadEscape
					return t.errorfn()
				}
				t.p.AcceptRun(non_at)
				if !t.p.Accept("@") {
					t.err = ErrBadEscape
					return t.errorfn()
				}
			} else {
				break
			}
		}
		p := t.p.Peek()
		if p == -1 {
			next = t.done
			break
		}
		// What follows is a hack for broken gedcom generators that insert newlines where they shouldn't go.
		// CONT & CONC are there for a reason!
		if strings.ContainsRune(terminators, p) {
			t.p.AcceptRun(terminators)
			p = t.p.Peek()
			if p == -1 {
				next = t.done
				break
			}
			if strings.ContainsRune(digit[1:], p) {
				break
			}
		}
	}
	return token{
		tokenLine,
		strings.TrimSpace(t.p.Get()),
	}, next
}

func (t *tokeniser) done() (token, stateFn) {
	t.err = io.EOF
	return token{
		tokenDone,
		"",
	}, t.done
}

func (t *tokeniser) errorfn() (token, stateFn) {
	return token{
		tokenError,
		t.err.Error(),
	}, t.errorfn
}

// Errors
var (
	ErrInvalidLevel   = errors.New("invalid level num")
	ErrMissingDelim   = errors.New("missing delminitator")
	ErrInvalidPointer = errors.New("invalid pointer string")
	ErrInvalidTag     = errors.New("invalid tag")
	ErrBadEscape      = errors.New("bad escape sequence")
)
