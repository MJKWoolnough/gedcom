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
		{ // 1
			Input: "",
			Output: []parser.Token{
				{Type: parser.TokenDone, Data: ""},
			},
			Options: options{},
		},
		{ // 2
			Input: " \t\r\n",
			Output: []parser.Token{
				{Type: parser.TokenDone, Data: ""},
			},
			Options: options{},
		},
		{ // 3
			Input: "a",
			Output: []parser.Token{
				{Type: parser.TokenError, Data: ErrInvalidLevel.Error()},
			},
			Options: options{},
		},
		{ // 4
			Input: "12a",
			Output: []parser.Token{
				{Type: parser.TokenError, Data: ErrMissingDelim.Error()},
			},
			Options: options{},
		},
		{ // 5
			Input: " 12 ",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: parser.TokenError, Data: ErrInvalidTag.Error()},
			},
			Options: options{},
		},
		{ // 6
			Input: "12 @",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: parser.TokenError, Data: ErrInvalidPointer.Error()},
			},
			Options: options{},
		},
		{ // 7
			Input: "12 @a123",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: parser.TokenError, Data: ErrInvalidPointer.Error()},
			},
			Options: options{},
		},
		{ // 8
			Input: "12 @a123@ ",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: tokenXref, Data: "a123"},
				{Type: parser.TokenError, Data: ErrInvalidTag.Error()},
			},
			Options: options{},
		},
		{ // 9
			Input: "12 @a123@",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: parser.TokenError, Data: ErrMissingDelim.Error()},
			},
			Options: options{},
		},
		{ // 10
			Input: "12 a",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: tokenTag, Data: "a"},
				{Type: parser.TokenDone, Data: ""},
			},
			Options: options{},
		},
		{ // 11
			Input: "12 a\n",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: tokenTag, Data: "a"},
				{Type: tokenEndLine, Data: ""},
				{Type: parser.TokenDone, Data: ""},
			},
			Options: options{},
		},
		{ // 12
			Input: "12 a|",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: parser.TokenError, Data: ErrInvalidTag.Error()},
			},
			Options: options{},
		},
		{ // 13
			Input: "12 @abc@ a\n",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: tokenXref, Data: "abc"},
				{Type: tokenTag, Data: "a"},
				{Type: tokenEndLine, Data: ""},
				{Type: parser.TokenDone, Data: ""},
			},
			Options: options{},
		},
		{ // 14
			Input: "12 @abc@ def @ghi",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: tokenXref, Data: "abc"},
				{Type: tokenTag, Data: "def"},
				{Type: parser.TokenError, Data: ErrInvalidPointer.Error()},
			},
			Options: options{},
		},
		{ // 15
			Input: "12 @abc@ def @ghi",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: tokenXref, Data: "abc"},
				{Type: tokenTag, Data: "def"},
				{Type: tokenLine, Data: "@ghi"},
				{Type: parser.TokenDone, Data: ""},
			},
			Options: options{allowInvalidEscape: true},
		},
		{ // 16
			Input: "12 @abc@ def @ghi@",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: tokenXref, Data: "abc"},
				{Type: tokenTag, Data: "def"},
				{Type: tokenPointer, Data: "ghi"},
				{Type: parser.TokenDone, Data: ""},
			},
			Options: options{},
		},
		{ // 17
			Input: "12 @abc@ def ghi jkl",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: tokenXref, Data: "abc"},
				{Type: tokenTag, Data: "def"},
				{Type: tokenLine, Data: "ghi jkl"},
				{Type: parser.TokenDone, Data: ""},
			},
			Options: options{},
		},
		{ // 18
			Input: "12 @abc@ def ghi jkl\n13 zyx @@ wvu",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: tokenXref, Data: "abc"},
				{Type: tokenTag, Data: "def"},
				{Type: tokenLine, Data: "ghi jkl"},
				{Type: tokenLevel, Data: "13"},
				{Type: tokenTag, Data: "zyx"},
				{Type: tokenLine, Data: "@@ wvu"},
				{Type: parser.TokenDone, Data: ""},
			},
			Options: options{},
		},
		{ // 19
			Input: "12 @abc@ def £",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: tokenXref, Data: "abc"},
				{Type: tokenTag, Data: "def"},
				{Type: parser.TokenError, Data: ErrBadChar.Error()},
			},
			Options: options{},
		},
		{ // 20
			Input: "12 @abc@ def £",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: tokenXref, Data: "abc"},
				{Type: tokenTag, Data: "def"},
				{Type: tokenLine, Data: "£"},
				{Type: parser.TokenDone, Data: ""},
			},
			Options: options{allowUnknownCharset: true},
		},
		{ // 21
			Input: "12 @abc@ def line @#98@",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: tokenXref, Data: "abc"},
				{Type: tokenTag, Data: "def"},
				{Type: tokenLine, Data: "line @#98@"},
				{Type: parser.TokenDone, Data: ""},
			},
			Options: options{},
		},
		{ // 22
			Input: "12 @abc@ def line @98@",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: tokenXref, Data: "abc"},
				{Type: tokenTag, Data: "def"},
				{Type: parser.TokenError, Data: ErrBadEscape.Error()},
			},
			Options: options{},
		},
		{ // 23
			Input: "12 @abc@ def @#98@",
			Output: []parser.Token{
				{Type: tokenLevel, Data: "12"},
				{Type: tokenXref, Data: "abc"},
				{Type: tokenTag, Data: "def"},
				{Type: tokenLine, Data: "@#98@"},
				{Type: parser.TokenDone, Data: ""},
			},
			Options: options{allowInvalidEscape: true},
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
