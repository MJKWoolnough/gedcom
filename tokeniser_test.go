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
