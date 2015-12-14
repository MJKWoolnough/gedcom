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
	otherchar   = "!\"$%&'{}*+,-./:;<=>?[\\]^`{|}~"
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
	next := t.level
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
