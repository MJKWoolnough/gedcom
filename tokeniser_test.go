package gedcom

import (
	"strings"
	"testing"

	"vimagination.zapto.org/parser"
)

func TestTokeniser(t *testing.T) {
	for n, test := range [...]struct {
		Input   string
		Output  []parser.Token
		Options options
	}{
		{
			Input: "",
			Output: []parser.Token{
				{Type: parser.TokenDone, Data: ""},
			},
			Options: options{},
		},
		{
			Input: " \t\r\n",
			Output: []parser.Token{
				{Type: parser.TokenDone, Data: ""},
			},
			Options: options{},
		},
		{
			Input: "a",
			Output: []parser.Token{
				{Type: parser.TokenError, Data: ErrInvalidLevel.Error()},
			},
			Options: options{},
		},
		{
			Input: "12a",
			Output: []parser.Token{
				{Type: parser.TokenError, Data: ErrMissingDelim.Error()},
			},
			Options: options{},
		},
		{
			Input: " 12 ",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: parser.TokenError, Data: ErrInvalidTag.Error()},
			},
			Options: options{},
		},
		{
			Input: "12 @",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: parser.TokenError, Data: ErrInvalidPointer.Error()},
			},
			Options: options{},
		},
		{
			Input: "12 @a123",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: parser.TokenError, Data: ErrInvalidPointer.Error()},
			},
			Options: options{},
		},
		{
			Input: "12 @a123@ ",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: tokenXref, Data: "a123"},
				{Type: parser.TokenError, Data: ErrInvalidTag.Error()},
			},
			Options: options{},
		},
	} {
		tks := newTokeniser(strings.NewReader(test.Input), test.Options)

		for m, tkn := range test.Output {
			if tk, _ := tks.GetToken(); tk.Type != tkn.Type {
				if tk.Type == parser.TokenError {
					t.Errorf("test %d.%d: unexpected error: %s", n+1, m+1, tk.Data)
				} else {
					t.Errorf("test %d.%d: Incorrect type, expecting %d, got %d", n+1, m+1, tkn.Type, tk.Type)
				}

				break
			} else if tk.Data != tkn.Data {
				t.Errorf("test %d.%d: Incorrect data, expecting %q, got %q", n+1, m+1, tkn.Data, tk.Data)

				break
			}
		}
	}
}
