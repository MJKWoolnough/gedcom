package gedcom

import (
	"errors"
	"io"
	"strings"

	"vimagination.zapto.org/parser"
)

const (
	alpha       = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_"
	delim       = " "
	digit       = "0123456789"
	alphanum    = alpha + digit
	otherchar   = "!\"$%&'()*+,-./:;<=>?[\\]^`{|}~" // + "\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8a\x8b\x8c\x8d\x8e\x8f\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9a\x9b\x9c\x9d\x9e\x9f\xa0\xa1\xa2\xa3\xa4\xa5\xa6\xa7\xa8\xa9\xaa\xab\xac\xad\xae\xaf\xb0\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe\xbf\xc0\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf\xd0\xd1\xd2\xd3\xd4\xd5\xd6\xd7\xd8\xd9\xda\xdb\xdc\xdd\xde\xdf\xe0\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xeb\xec\xed\xee\xef\xf0\xf1\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xfb\xfc\xfd\xfe" // ANSEL
	nonAt       = alpha + digit + otherchar + " #"
	terminators = "\r\n"
	//	anyChar     = alpha + digit + otherchar + "# @"
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
		Parser:  parser.New(parser.NewReaderTokeniser(r)),
		options: o,
	}

	t.TokeniserState(t.level)

	return t
}

func (t *tokeniser) level(p *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	p.AcceptRun(levelIgnore)
	p.Get()

	if p.Peek() == -1 {
		return p.Done()
	} else if !p.Accept(digit) {
		p.Err = ErrInvalidLevel

		return p.Error()
	}

	p.AcceptRun(digit)

	if !p.Accept(delim) {
		p.Err = ErrMissingDelim

		return p.Error()
	}

	return parser.Token{
		Type: tokenLevel,
		Data: strings.TrimSpace(p.Get()),
	}, t.optionalXrefID
}

func (t *tokeniser) optionalXrefID(p *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	if p.Peek() == '@' {
		return t.xrefID(p)
	}

	return t.tag(p)
}

func (t *tokeniser) readPointer(p *parser.Tokeniser) (string, error) {
	if !p.Accept(alphanum) {
		return "", ErrInvalidPointer
	}

	p.AcceptRun(nonAt)

	if !p.Accept("@") {
		return "", ErrInvalidPointer
	}

	pointer := p.Get()

	return pointer[1 : len(pointer)-1], nil
}

func (t *tokeniser) xrefID(p *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	p.Accept("@")

	pointer, err := t.readPointer(p)
	if err != nil {
		p.Err = err

		return p.Error()
	} else if !p.Accept(delim) {
		p.Err = ErrMissingDelim

		return p.Error()
	}

	p.Get()

	return parser.Token{
		Type: tokenXref,
		Data: strings.Trim(pointer, "@"),
	}, t.tag
}

func (t *tokeniser) tag(p *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	if !p.Accept(alphanum) {
		p.Err = ErrInvalidTag
		return p.Error()
	}

	p.AcceptRun(alphanum)

	tag := p.Get()

	if p.Accept(delim) {
		p.Get()

		return parser.Token{
			Type: tokenTag,
			Data: tag,
		}, t.lineValue
	}

	next := t.endLine

	if p.Peek() == -1 {
		next = (*parser.Tokeniser).Done
	} else {
		if !p.Accept(terminators) {
			p.Err = ErrInvalidTag

			return p.Error()
		}

		p.AcceptRun(terminators)
		p.Get()
	}

	return parser.Token{
		Type: tokenTag,
		Data: tag,
	}, next
}

func (t *tokeniser) endLine(p *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	return parser.Token{
		Type: tokenEndLine,
		Data: "",
	}, t.level
}

func (t *tokeniser) lineValue(p *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	if p.Peek() == '@' {
		p.Accept("@")

		if p.Peek() != '@' {
			if pointer, err := t.readPointer(p); err != nil {
				if !t.allowInvalidEscape {
					p.Err = err

					return p.Error()
				}
			} else {
				p.AcceptRun(terminators)
				p.Get()

				return parser.Token{
					Type: tokenPointer,
					Data: pointer,
				}, t.level
			}
		}
	}

	next := t.level

	for {
		for {
			if t.allowUnknownCharset && p.Except(invalidchar) {
				p.ExceptRun(invalidchar)
			} else if p.Accept(nonAt) {
				p.AcceptRun(nonAt)
			} else if p.Accept("@") {
				if t.allowInvalidEscape || p.Accept("@") {
					continue
				}

				if !p.Accept("#") {
					p.Err = ErrBadEscape

					return p.Error()
				} else if t.allowUnknownCharset {
					p.ExceptRun(invalidchar)
				} else {
					p.AcceptRun(nonAt)
				}

				if !p.Accept("@") {
					p.Err = ErrBadEscape

					return p.Error()
				}
			} else {
				break
			}
		}

		pe := p.Peek()
		if pe == -1 {
			next = (*parser.Tokeniser).Done

			break
		}

		if strings.ContainsRune(terminators, pe) {
			p.AcceptRun(terminators)

			if !t.allowTerminatorsInValue {
				break
			}

			if pe = p.Peek(); pe == -1 {
				next = (*parser.Tokeniser).Done

				break
			} else if strings.ContainsRune(digit, pe) {
				break
			}
		} else {
			if !t.allowInvalidChars {
				p.Err = ErrBadChar

				return p.Error()
			}

			p.Except("")
		}
	}

	return parser.Token{
		Type: tokenLine,
		Data: strings.TrimSpace(p.Get()),
	}, next
}

// Errors.
var (
	ErrInvalidLevel   = errors.New("invalid level num")
	ErrMissingDelim   = errors.New("missing delimiter")
	ErrInvalidPointer = errors.New("invalid pointer string")
	ErrInvalidTag     = errors.New("invalid tag")
	ErrBadEscape      = errors.New("bad escape sequence")
	ErrBadChar        = errors.New("bad character")
)
