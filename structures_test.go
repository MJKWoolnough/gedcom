package gedcom

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestHeader(t *testing.T) {
	for n, test := range [...]struct {
		Input   string
		Options []Option
		Output  Header
		Err     error
	}{
		{ // 1
			Input: "0 HEADER\n0 TRLR",
			Err:   ErrContext{"Header", "Source", ErrRequiredMissing},
		},
		{ // 2
			Input: "0 HEADER\n1 SOUR id\n0 TRLR",
			Err:   ErrContext{"Header", "Submitter", ErrRequiredMissing},
		},
		{ // 3
			Input: "0 HEADER\n1 SOUR id\n1 SUBM submitter\n0 TRLR",
			Err:   ErrContext{"Header", "Version", ErrRequiredMissing},
		},
		{ // 4
			Input: "0 HEADER\n1 SOUR id\n1 SUBM submitter\n\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n0 TRLR",
			Err:   ErrContext{"Header", "CharacterSet", ErrRequiredMissing},
		},
		{ // 5
			Input: "0 HEADER\n1 SOUR id\n1 SUBM submitter\n\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL\n0 TRLR",
			Output: Header{
				Source: HeaderSource{
					SystemID: ApprovedSystemID("id"),
				},
				Submitter: "submitter",
				Version: Version{
					VersionNumber: "5.5",
					Form:          "LINEAGE-LINKED",
				},
				CharacterSet: CharacterSetStructure{
					CharacterSet: "ANSEL",
				},
			},
		},
		{ // 6
			Input: "0 HEADER\n1 SOUR\n1 SUBM submitter\n\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL\n0 TRLR",
			Err:   ErrContext{"Header", cSOUR, ErrContext{"HeaderSource", "line_value", ErrInvalidLength{"ApprovedSystemID", "", 1, 20}}},
		},
		{ // 7
			Input: "0 HEADER\n1 SOUR id\n1 SOUR other\n1 SUBM submitter\n\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL\n0 TRLR",
			Err:   ErrContext{"Header", cSOUR, ErrSingleMultiple},
		},
		{ // 8
			Input:   "0 HEADER\n1 SOUR id\n1 SOUR other\n1 SUBM submitter\n\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL\n0 TRLR",
			Options: []Option{AllowMoreThanAllowed},
			Output: Header{
				Source: HeaderSource{
					SystemID: ApprovedSystemID("id"),
				},
				Submitter: "submitter",
				Version: Version{
					VersionNumber: "5.5",
					Form:          "LINEAGE-LINKED",
				},
				CharacterSet: CharacterSetStructure{
					CharacterSet: "ANSEL",
				},
			},
		},
		{ // 9
			Input: "0 HEADER\n1 SOUR id\n1 DEST destination\n1 SUBM submitter\n\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL\n0 TRLR",
			Output: Header{
				Source: HeaderSource{
					SystemID: ApprovedSystemID("id"),
				},
				Submitter: "submitter",
				Version: Version{
					VersionNumber: "5.5",
					Form:          "LINEAGE-LINKED",
				},
				CharacterSet: CharacterSetStructure{
					CharacterSet: "ANSEL",
				},
				ReceivingSystemName: "destination",
			},
		},
		{ // 10
			Input: "0 HEADER\n1 SOUR id\n1 DEST\n1 SUBM submitter\n\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL\n0 TRLR",
			Err:   ErrContext{"Header", cDEST, ErrInvalidLength{"ReceivingSystemName", "", 1, 20}},
		},
		{ // 11
			Input: "0 HEADER\n1 SOUR id\n1 DEST destination\n1 DEST destination2\n1 SUBM submitter\n\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL\n0 TRLR",
			Err:   ErrContext{"Header", cDEST, ErrSingleMultiple},
		},
		{ // 12
			Input:   "0 HEADER\n1 SOUR id\n1 DEST destination\n1 SUBM submitter\n\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL\n0 TRLR",
			Options: []Option{AllowMoreThanAllowed},
			Output: Header{
				Source: HeaderSource{
					SystemID: ApprovedSystemID("id"),
				},
				Submitter: "submitter",
				Version: Version{
					VersionNumber: "5.5",
					Form:          "LINEAGE-LINKED",
				},
				CharacterSet: CharacterSetStructure{
					CharacterSet: "ANSEL",
				},
				ReceivingSystemName: "destination",
			},
		},
	} {
		var s Header

		r := NewReader(strings.NewReader(test.Input), test.Options...)

		if lines, err := readLines(r); err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
		} else {
			l := parseLines(lines)

			if err = s.parse(&l, r.options); !errors.Is(err, test.Err) {
				t.Errorf("test %d: expecting error %v, got %v", n+1, test.Err, err)
			} else if test.Err == nil && !reflect.DeepEqual(test.Output, s) {
				t.Errorf("test %d: expecting %#v to equal %#v", n+1, test.Output, s)
			}
		}
	}
}
