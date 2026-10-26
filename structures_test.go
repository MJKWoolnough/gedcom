package gedcom

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

type testData[T any] struct {
	Input   string
	Options []Option
	Output  T
	Err     error
}

func test[T any, I interface {
	*T
	parse(l *Line, o options) error
}](t *testing.T, tests []testData[T]) {
	for n, test := range tests {
		var s T

		var si I = &s

		r := NewReader(strings.NewReader(test.Input), test.Options...)

		if lines, err := readLines(r); err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
		} else {
			l := parseLines(lines)

			if err = si.parse(&l, r.options); !errors.Is(err, test.Err) {
				t.Errorf("test %d: expecting error %v, got %v", n+1, test.Err, err)
			} else if test.Err == nil && !reflect.DeepEqual(test.Output, s) {
				t.Errorf("test %d: expecting %#v, got %#v", n+1, test.Output, s)
			}
		}
	}
}

func TestHeader(t *testing.T) {
	test(t, []testData[Header]{
		{ // 1
			Input: "0 HEADER\n",
			Err:   ErrContext{"Header", cSUBM, ErrRequiredMissing},
		},
		{ // 2
			Input:   "0 HEADER\n",
			Options: []Option{AllowMissingRequired},
		},
		{ // 3
			Input: "0 HEADER\n",
			Err:   ErrContext{"Header", cSUBM, ErrRequiredMissing},
		},
		{ // 4
			Input: "0 HEADER\n1 SUBM submitter",
			Err:   ErrContext{"Header", cGEDC, ErrRequiredMissing},
		},
		{ // 5
			Input: "0 HEADER\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED",
			Err:   ErrContext{"Header", cCHAR, ErrRequiredMissing},
		},
		{ // 6
			Input: "0 HEADER\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cSOUR, ErrRequiredMissing},
		},
		{ // 7
			Input: "0 HEADER\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
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
		{ // 8
			Input: "0 HEADER\n1 SOUR\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cSOUR, ErrContext{"HeaderSource", "line_value", ErrInvalidLength{"ApprovedSystemID", "", 1, 20}}},
		},
		{ // 9
			Input: "0 HEADER\n1 SOUR id\n1 SOUR other\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cSOUR, ErrSingleMultiple},
		},
		{ // 10
			Input:   "0 HEADER\n1 SOUR id\n1 SOUR other\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
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
		{ // 11
			Input: "0 HEADER\n1 SOUR id\n1 DEST destination\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
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
		{ // 12
			Input: "0 HEADER\n1 SOUR id\n1 DEST\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cDEST, ErrInvalidLength{"ReceivingSystemName", "", 1, 20}},
		},
		{ // 13
			Input: "0 HEADER\n1 SOUR id\n1 DEST destination\n1 DEST destination2\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cDEST, ErrSingleMultiple},
		},
		{ // 14
			Input:   "0 HEADER\n1 SOUR id\n1 DEST destination\n1 DEST destination2\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
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
		{ // 15
			Input: "0 HEADER\n1 SOUR id\n1 DATE 2006-05-04\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
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
				TransmissionLDate: TransmissionDateTime{
					TransmissionDate: "2006-05-04",
				},
			},
		},
		{ // 16
			Input: "0 HEADER\n1 SOUR id\n1 DATE\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cDATE, ErrContext{"TransmissionDateTime", "line_value", ErrInvalidLength{"TransmissionDate", "", 10, 11}}},
		},
		{ // 17
			Input: "0 HEADER\n1 SOUR id\n1 DATE 2006-05-04\n1 DATE 2007-06-05\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cDATE, ErrSingleMultiple},
		},
		{ // 18
			Input:   "0 HEADER\n1 SOUR id\n1 DATE 2006-05-04\n1 DATE 2007-06-05\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
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
				TransmissionLDate: TransmissionDateTime{
					TransmissionDate: "2006-05-04",
				},
			},
		},
		{ // 19
			Input: "0 HEADER\n1 SOUR id\n1 DEST destination\n1 SUBM\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cSUBM, ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 20
			Input: "0 HEADER\n1 SOUR id\n1 DEST destination\n1 SUBM submitter\n1 SUBM submitter2\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cSUBM, ErrSingleMultiple},
		},
		{ // 21
			Input:   "0 HEADER\n1 SOUR id\n1 DEST destination\n1 SUBM submitter\n1 SUBM submitter2\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
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
		{ // 22
			Input: "0 HEADER\n1 SUMB xid\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
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
				Submission: "xid",
			},
		},
		{ // 23
			Input: "0 HEADER\n1 SUMB\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cSUMB, ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 24
			Input: "0 HEADER\n1 SUMB xid1\n1 SUMB xid2\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cSUMB, ErrSingleMultiple},
		},
		{ // 25
			Input:   "0 HEADER\n1 SUMB xid1\n1 SUMB xid2\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
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
				Submission: "xid1",
			},
		},
		{ // 26
			Input: "0 HEADER\n1 FILE filename\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
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
				FileName: "filename",
			},
		},
		{ // 27
			Input: "0 HEADER\n1 FILE\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cFILE, ErrInvalidLength{"FileName", "", 1, 90}},
		},
		{ // 28
			Input: "0 HEADER\n1 FILE filename1\n1 FILE filename2\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cFILE, ErrSingleMultiple},
		},
		{ // 29
			Input:   "0 HEADER\n1 FILE filename1\n1 FILE filename2\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
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
				FileName: "filename1",
			},
		},
		{ // 30
			Input: "0 HEADER\n1 COPR copyright\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
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
				Copyright: "copyright",
			},
		},
		{ // 31
			Input: "0 HEADER\n1 COPR\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cCOPR, ErrInvalidLength{"CopyrightGedcomFile", "", 1, 90}},
		},
		{ // 32
			Input: "0 HEADER\n1 COPR copyright1\n1 COPR copyright2\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cCOPR, ErrSingleMultiple},
		},
		{ // 33
			Input:   "0 HEADER\n1 COPR copyright1\n1 COPR copyright2\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
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
				Copyright: "copyright1",
			},
		},
		{ // 34
			Input: "0 HEADER\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cGEDC, ErrContext{"Version", cVERS, ErrRequiredMissing}},
		},
		{ // 35
			Input: "0 HEADER\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 GEDC\n2 VERS 5.6\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cGEDC, ErrSingleMultiple},
		},
		{ // 36
			Input:   "0 HEADER\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 GEDC\n2 VERS 5.6\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
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
		{ // 37
			Input: "0 HEADER\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR utf-8",
			Err:   ErrContext{"Header", cCHAR, ErrContext{"CharacterSetStructure", "line_value", ErrInvalidValue{"CharacterSet", "utf-8"}}},
		},
		{ // 38
			Input: "0 HEADER\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL\n1 CHAR ASCII",
			Err:   ErrContext{"Header", cCHAR, ErrSingleMultiple},
		},
		{ // 39
			Input:   "0 HEADER\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ASCII\n1 CHAR ANSEL",
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
					CharacterSet: "ASCII",
				},
			},
		},
		{ // 40
			Input: "0 HEADER\n1 LANG eng\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
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
				Language: "eng",
			},
		},
		{ // 41
			Input: "0 HEADER\n1 LANG\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cLANG, ErrInvalidLength{"LanguageOfText", "", 1, 15}},
		},
		{ // 42
			Input: "0 HEADER\n1 LANG eng\n1 LANG fre\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cLANG, ErrSingleMultiple},
		},
		{ // 43
			Input:   "0 HEADER\n1 LANG eng\n1 LANG fre\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
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
				Language: "eng",
			},
		},
		{ // 44
			Input: "0 HEADER\n1 PLAC\n2 FORM place\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
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
				Place: HeaderPlace{
					PlaceHierarchy: "place",
				},
			},
		},
		{ // 45
			Input: "0 HEADER\n1 PLAC\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cPLAC, ErrContext{"HeaderPlace", cFORM, ErrRequiredMissing}},
		},
		{ // 46
			Input: "0 HEADER\n1 PLAC\n2 FORM place1\n1 PLAC\n2 FORM place2\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cPLAC, ErrSingleMultiple},
		},
		{ // 47
			Input:   "0 HEADER\n1 PLAC\n2 FORM place1\n1 PLAC\n2 FORM place2\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
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
				Place: HeaderPlace{
					PlaceHierarchy: "place1",
				},
			},
		},
		{ // 48
			Input: "0 HEADER\n1 NOTE note\n2 FORM place\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
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
				ContentDescription: "note",
			},
		},
		{ // 49
			Input: "0 HEADER\n1 NOTE\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cNOTE, ErrInvalidLength{"ContentDescription", "", 1, 248}},
		},
		{ // 50
			Input: "0 HEADER\n1 NOTE note1\n1 NOTE note2\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cNOTE, ErrSingleMultiple},
		},
		{ // 51
			Input:   "0 HEADER\n1 NOTE note1\n1 NOTE note2\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
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
				ContentDescription: "note1",
			},
		},
		{ // 52
			Input: "0 HEADER\n1 UNKNOWN\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", "UNKNOWN", ErrUnknownTag},
		},
		{ // 53
			Input: "0 HEADER\n1 _UNKNOWN\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
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
		{ // 54
			Input:   "0 HEADER\n1 UNKNOWN\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Options: []Option{AllowUnknownTags},
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
	})
}

func TestHeaderSource(t *testing.T) {
	test(t, []testData[HeaderSource]{
		{ // 1
			Input: "0 SOUR\n",
			Err:   ErrContext{"HeaderSource", "line_value", ErrInvalidLength{"ApprovedSystemID", "", 1, 20}},
		},
		{ // 2
			Input: "0 SOUR ID\n",
			Output: HeaderSource{
				SystemID: "ID",
			},
		},
		{ // 3
			Input: "0 SOUR ID\n1 VERS 5.5",
			Output: HeaderSource{
				SystemID:      "ID",
				VersionNumber: "5.5",
			},
		},
		{ // 4
			Input: "0 SOUR ID\n1 VERS\n",
			Err:   ErrContext{"HeaderSource", cVERS, ErrInvalidLength{"VersionNumber", "", 1, 15}},
		},
		{ // 5
			Input: "0 SOUR ID\n1 VERS 5.6\n1 VERS 5.7",
			Err:   ErrContext{"HeaderSource", cVERS, ErrSingleMultiple},
		},
		{ // 6
			Input:   "0 SOUR ID\n1 VERS 5.6\n1 VERS 5.7",
			Options: []Option{AllowMoreThanAllowed},
			Output: HeaderSource{
				SystemID:      "ID",
				VersionNumber: "5.6",
			},
		},
		{ // 7
			Input: "0 SOUR ID\n1 NAME NoP",
			Output: HeaderSource{
				SystemID: "ID",
				Name:     "NoP",
			},
		},
		{ // 8
			Input: "0 SOUR ID\n1 NAME\n",
			Err:   ErrContext{"HeaderSource", cNAME, ErrInvalidLength{"NameOfProduct", "", 1, 90}},
		},
		{ // 9
			Input: "0 SOUR ID\n1 NAME NoQ\n1 NAME NoR",
			Err:   ErrContext{"HeaderSource", cNAME, ErrSingleMultiple},
		},
		{ // 10
			Input:   "0 SOUR ID\n1 NAME NoQ\n1 NAME NoR",
			Options: []Option{AllowMoreThanAllowed},
			Output: HeaderSource{
				SystemID: "ID",
				Name:     "NoQ",
			},
		},
		{ // 11
			Input: "0 SOUR ID\n1 CORP BusinessName",
			Output: HeaderSource{
				SystemID: "ID",
				Business: HeaderBusiness{
					NameOfBusiness: "BusinessName",
					PhoneNumber:    []PhoneNumber{},
				},
			},
		},
		{ // 12
			Input: "0 SOUR ID\n1 CORP\n",
			Err:   ErrContext{"HeaderSource", cCORP, ErrContext{"HeaderBusiness", "line_value", ErrInvalidLength{"NameOfBusiness", "", 1, 90}}},
		},
		{ // 13
			Input: "0 SOUR ID\n1 CORP BusinessName2\n1 CORP BusinessName3",
			Err:   ErrContext{"HeaderSource", cCORP, ErrSingleMultiple},
		},
		{ // 14
			Input:   "0 SOUR ID\n1 CORP BusinessName2\n1 CORP BusinessName3",
			Options: []Option{AllowMoreThanAllowed},
			Output: HeaderSource{
				SystemID: "ID",
				Business: HeaderBusiness{
					NameOfBusiness: "BusinessName2",
					PhoneNumber:    []PhoneNumber{},
				},
			},
		},
		{ // 15
			Input: "0 SOUR ID\n1 DATA DataSourceName",
			Output: HeaderSource{
				SystemID: "ID",
				Data: HeaderDataSource{
					SourceName: "DataSourceName",
				},
			},
		},
		{ // 16
			Input: "0 SOUR ID\n1 DATA\n",
			Err:   ErrContext{"HeaderSource", cDATA, ErrContext{"HeaderDataSource", "line_value", ErrInvalidLength{"NameOfSourceData", "", 1, 90}}},
		},
		{ // 17
			Input: "0 SOUR ID\n1 DATA DataSourceName2\n1 DATA DataSourceName3",
			Err:   ErrContext{"HeaderSource", cDATA, ErrSingleMultiple},
		},
		{ // 18
			Input:   "0 SOUR ID\n1 DATA DataSourceName2\n1 DATA DataSourceName3",
			Options: []Option{AllowMoreThanAllowed},
			Output: HeaderSource{
				SystemID: "ID",
				Data: HeaderDataSource{
					SourceName: "DataSourceName2",
				},
			},
		},
		{ // 19
			Input: "0 SOUR ID\n1 UNKNOWN\n",
			Err:   ErrContext{"HeaderSource", "UNKNOWN", ErrUnknownTag},
		},
		{ // 20
			Input: "0 SOUR ID\n1 _UNKNOWN\n",
			Output: HeaderSource{
				SystemID: "ID",
			},
		},
		{ // 21
			Input:   "0 SOUR ID\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: HeaderSource{
				SystemID: "ID",
			},
		},
	})
}

func TestTransmissionDateTime(t *testing.T) {
	test(t, []testData[TransmissionDateTime]{
		{ // 1
			Input: "0 DATE\n",
			Err:   ErrContext{"TransmissionDateTime", "line_value", ErrInvalidLength{"TransmissionDate", "", 10, 11}},
		},
		{ // 2
			Input: "0 DATE 2006-05-04\n",
			Output: TransmissionDateTime{
				TransmissionDate: "2006-05-04",
			},
		},
		{ // 3
			Input: "0 DATE 2006-05-04\n1 TIME 15:02",
			Output: TransmissionDateTime{
				TransmissionDate: "2006-05-04",
				Time:             "15:02",
			},
		},
		{ // 4
			Input: "0 DATE 2006-05-04\n1 TIME\n",
			Err:   ErrContext{"TransmissionDateTime", cTIME, ErrInvalidLength{"TimeValue", "", 1, 12}},
		},
		{ // 5
			Input: "0 DATE 2006-05-04\n1 TIME 15:03\n1 TIME 15:04",
			Err:   ErrContext{"TransmissionDateTime", cTIME, ErrSingleMultiple},
		},
		{ // 6
			Input:   "0 DATE 2006-05-04\n1 TIME 15:03\n1 TIME 15:04",
			Options: []Option{AllowMoreThanAllowed},
			Output: TransmissionDateTime{
				TransmissionDate: "2006-05-04",
				Time:             "15:03",
			},
		},
		{ // 7
			Input: "0 DATE 2006-05-04\n1 UNKNOWN\n",
			Err:   ErrContext{"TransmissionDateTime", "UNKNOWN", ErrUnknownTag},
		},
		{ // 8
			Input: "0 DATE 2006-05-04\n1 _UNKNOWN\n",
			Output: TransmissionDateTime{
				TransmissionDate: "2006-05-04",
			},
		},
		{ // 9
			Input:   "0 DATE 2006-05-04\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: TransmissionDateTime{
				TransmissionDate: "2006-05-04",
			},
		},
	})
}

func TestHeaderBusiness(t *testing.T) {
	test(t, []testData[HeaderBusiness]{
		{ // 1
			Input: "0 CORP\n",
			Err:   ErrContext{"HeaderBusiness", "line_value", ErrInvalidLength{"NameOfBusiness", "", 1, 90}},
		},
		{ // 2
			Input: "0 CORP business name\n",
			Output: HeaderBusiness{
				NameOfBusiness: "business name",
				PhoneNumber:    []PhoneNumber{},
			},
		},
		{ // 3
			Input: "0 CORP business name\n1 ADDR line 1",
			Output: HeaderBusiness{
				NameOfBusiness: "business name",
				Address: AddressStructure{
					AddressLine: "line 1",
				},
				PhoneNumber: []PhoneNumber{},
			},
		},
		{ // 4
			Input: "0 CORP business name\n1 ADDR\n",
			Err:   ErrContext{"HeaderBusiness", cADDR, ErrContext{"AddressStructure", "line_value", ErrInvalidLength{"AddressLine", "", 1, 60}}},
		},
		{ // 5
			Input: "0 CORP business name\n1 ADDR line_1\n1 ADDR line_2",
			Err:   ErrContext{"HeaderBusiness", cADDR, ErrSingleMultiple},
		},
		{ // 6
			Input:   "0 CORP business name\n1 ADDR line_1\n1 ADDR line_2",
			Options: []Option{AllowMoreThanAllowed},
			Output: HeaderBusiness{
				NameOfBusiness: "business name",
				Address: AddressStructure{
					AddressLine: "line_1",
				},
				PhoneNumber: []PhoneNumber{},
			},
		},
		{ // 7
			Input: "0 CORP business name\n1 PHON number",
			Output: HeaderBusiness{
				NameOfBusiness: "business name",
				PhoneNumber:    []PhoneNumber{"number"},
			},
		},
		{ // 8
			Input: "0 CORP business name\n1 PHON\n",
			Err:   ErrContext{"HeaderBusiness", cPHON, ErrInvalidLength{"PhoneNumber", "", 1, 25}},
		},
		{ // 9
			Input: "0 CORP business name\n1 PHON number 1\n1 PHON number 2",
			Output: HeaderBusiness{
				NameOfBusiness: "business name",
				PhoneNumber:    []PhoneNumber{"number 1", "number 2"},
			},
		},
		{ // 10
			Input: "0 CORP business name\n1 PHON number 1\n1 PHON number 2\n1 PHON number 3",
			Output: HeaderBusiness{
				NameOfBusiness: "business name",
				PhoneNumber:    []PhoneNumber{"number 1", "number 2", "number 3"},
			},
		},
		{ // 11
			Input: "0 CORP business name\n1 PHON number 1\n1 PHON number 2\n1 PHON number 3\n1 PHON number 4",
			Err:   ErrContext{"HeaderBusiness", cPHON, ErrTooMany(3)},
		},
		{ // 12
			Input:   "0 CORP business name\n1 PHON number 1\n1 PHON number 2\n1 PHON number 3\n1 PHON number 4",
			Options: []Option{AllowMoreThanAllowed},
			Output: HeaderBusiness{
				NameOfBusiness: "business name",
				PhoneNumber:    []PhoneNumber{"number 1", "number 2", "number 3"},
			},
		},
		{ // 13
			Input: "0 CORP business name\n1 UNKNOWN\n",
			Err:   ErrContext{"HeaderBusiness", "UNKNOWN", ErrUnknownTag},
		},
		{ // 14
			Input: "0 CORP business name\n1 _UNKNOWN\n",
			Output: HeaderBusiness{
				NameOfBusiness: "business name",
				PhoneNumber:    []PhoneNumber{},
			},
		},
		{ // 15
			Input:   "0 CORP business name\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: HeaderBusiness{
				NameOfBusiness: "business name",
				PhoneNumber:    []PhoneNumber{},
			},
		},
	})
}

func TestHeaderDataSource(t *testing.T) {
	test(t, []testData[HeaderDataSource]{
		{ // 1
			Input: "0 DATA\n",
			Err:   ErrContext{"HeaderDataSource", "line_value", ErrInvalidLength{"NameOfSourceData", "", 1, 90}},
		},
		{ // 2
			Input: "0 DATA data-source\n",
			Output: HeaderDataSource{
				SourceName: "data-source",
			},
		},
		{ // 3
			Input: "0 DATA data-source\n1 DATE 2026-09-30",
			Output: HeaderDataSource{
				SourceName:      "data-source",
				PublicationDate: "2026-09-30",
			},
		},
		{ // 4
			Input: "0 DATA data-source\n1 DATE\n",
			Err:   ErrContext{"HeaderDataSource", cDATE, ErrInvalidLength{"PublicationDate", "", 10, 11}},
		},
		{ // 5
			Input: "0 DATA data-source\n1 DATE 2026-09-30\n1 DATE 2025-08-29",
			Err:   ErrContext{"HeaderDataSource", cDATE, ErrSingleMultiple},
		},
		{ // 6
			Input:   "0 DATA data-source\n1 DATE 2026-09-30\n1 DATE 2025-08-29",
			Options: []Option{AllowMoreThanAllowed},
			Output: HeaderDataSource{
				SourceName:      "data-source",
				PublicationDate: "2026-09-30",
			},
		},
		{ // 7
			Input: "0 DATA data-source\n1 COPR copyright data",
			Output: HeaderDataSource{
				SourceName:          "data-source",
				CopyrightSourceData: "copyright data",
			},
		},
		{ // 8
			Input: "0 DATA data-source\n1 COPR\n",
			Err:   ErrContext{"HeaderDataSource", cCOPR, ErrInvalidLength{"CopyrightSourceData", "", 1, 90}},
		},
		{ // 9
			Input: "0 DATA data-source\n1 COPR copyright-data\n1 COPR data for copyright",
			Err:   ErrContext{"HeaderDataSource", cCOPR, ErrSingleMultiple},
		},
		{ // 10
			Input:   "0 DATA data-source\n1 COPR copyright-data\n1 COPR data for copyright",
			Options: []Option{AllowMoreThanAllowed},
			Output: HeaderDataSource{
				SourceName:          "data-source",
				CopyrightSourceData: "copyright-data",
			},
		},
		{ // 11
			Input: "0 DATA data-source\n1 UNKNOWN\n",
			Err:   ErrContext{"HeaderDataSource", "UNKNOWN", ErrUnknownTag},
		},
		{ // 12
			Input: "0 DATA data-source\n1 _UNKNOWN\n",
			Output: HeaderDataSource{
				SourceName: "data-source",
			},
		},
		{ // 13
			Input:   "0 DATA data-source\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: HeaderDataSource{
				SourceName: "data-source",
			},
		},
	})
}

func TestVersion(t *testing.T) {
	test(t, []testData[Version]{
		{ // 1
			Input: "0 GEDC\n",
			Err:   ErrContext{"Version", cVERS, ErrRequiredMissing},
		},
		{ // 2
			Input: "0 GEDC\n1 VERS 1.1",
			Err:   ErrContext{"Version", cFORM, ErrRequiredMissing},
		},
		{ // 3
			Input: "0 GEDC\n1 VERS 1.1\n1 FORM LINEAGE-LINKED",
			Output: Version{
				VersionNumber: "1.1",
				Form:          "LINEAGE-LINKED",
			},
		},
		{ // 4
			Input: "0 GEDC\n1 VERS\n1 FORM LINEAGE-LINKED",
			Err:   ErrContext{"Version", cVERS, ErrInvalidLength{"VersionNumber", "", 1, 15}},
		},
		{ // 5
			Input: "0 GEDC\n1 VERS 1.1\n1 VERS 2.2\n1 FORM LINEAGE-LINKED",
			Err:   ErrContext{"Version", cVERS, ErrSingleMultiple},
		},
		{ // 6
			Input:   "0 GEDC\n1 VERS 1.1\n1 VERS 2.2\n1 FORM LINEAGE-LINKED",
			Options: []Option{AllowMoreThanAllowed},
			Output: Version{
				VersionNumber: "1.1",
				Form:          "LINEAGE-LINKED",
			},
		},
		{ // 7
			Input: "0 GEDC\n1 VERS 1.1\n1 FORM\n",
			Err:   ErrContext{"Version", cFORM, ErrInvalidLength{"Form", "", 14, 20}},
		},
		{ // 8
			Input: "0 GEDC\n1 VERS 1.1\n1 FORM LINEAGE-LINKED\n1 FORM LINKED LINEAGE",
			Err:   ErrContext{"Version", cFORM, ErrSingleMultiple},
		},
		{ // 9
			Input:   "0 GEDC\n1 VERS 1.1\n1 FORM LINEAGE-LINKED\n1 FORM LINKED LINEAGE",
			Options: []Option{AllowMoreThanAllowed},
			Output: Version{
				VersionNumber: "1.1",
				Form:          "LINEAGE-LINKED",
			},
		},
		{ // 10
			Input: "0 GEDC\n1 VERS 1.1\n1 FORM LINEAGE-LINKED\n1 UNKNOWN\n",
			Err:   ErrContext{"Version", "UNKNOWN", ErrUnknownTag},
		},
		{ // 11
			Input: "0 GEDC\n1 VERS 1.1\n1 FORM LINEAGE-LINKED\n1 _UNKNOWN\n",
			Output: Version{
				VersionNumber: "1.1",
				Form:          "LINEAGE-LINKED",
			},
		},
		{ // 12
			Input:   "0 GEDC\n1 VERS 1.1\n1 FORM LINEAGE-LINKED\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: Version{
				VersionNumber: "1.1",
				Form:          "LINEAGE-LINKED",
			},
		},
	})
}

func TestCharacterSetStructure(t *testing.T) {
	test(t, []testData[CharacterSetStructure]{
		{ // 1
			Input: "0 CHAR\n",
			Err:   ErrContext{"CharacterSetStructure", "line_value", ErrInvalidValue{"CharacterSet", ""}},
		},
		{ // 2
			Input: "0 CHAR ASCII\n",
			Output: CharacterSetStructure{
				CharacterSet: "ASCII",
			},
		},
		{ // 3
			Input: "0 CHAR ASCII\n1 VERS 1.1",
			Output: CharacterSetStructure{
				CharacterSet:  "ASCII",
				VersionNumber: "1.1",
			},
		},
		{ // 4
			Input: "0 CHAR ASCII\n1 VERS\n",
			Err:   ErrContext{"CharacterSetStructure", cVERS, ErrInvalidLength{"VersionNumber", "", 1, 15}},
		},
		{ // 5
			Input: "0 CHAR ASCII\n1 VERS 1.2\n1 VERS 2.3",
			Err:   ErrContext{"CharacterSetStructure", cVERS, ErrSingleMultiple},
		},
		{ // 6
			Input:   "0 CHAR ASCII\n1 VERS 1.2\n1 VERS 2.3",
			Options: []Option{AllowMoreThanAllowed},
			Output: CharacterSetStructure{
				CharacterSet:  "ASCII",
				VersionNumber: "1.2",
			},
		},
		{ // 7
			Input: "0 CHAR ASCII\n1 UNKNOWN\n",
			Err:   ErrContext{"CharacterSetStructure", "UNKNOWN", ErrUnknownTag},
		},
		{ // 8
			Input: "0 CHAR ASCII\n1 _UNKNOWN\n",
			Output: CharacterSetStructure{
				CharacterSet: "ASCII",
			},
		},
		{ // 9
			Input:   "0 CHAR ASCII\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: CharacterSetStructure{
				CharacterSet: "ASCII",
			},
		},
	})
}

func TestHeaderPlace(t *testing.T) {
	test(t, []testData[HeaderPlace]{
		{ // 1
			Input: "0 PLAC\n",
			Err:   ErrContext{"HeaderPlace", cFORM, ErrRequiredMissing},
		},
		{ // 2
			Input: "0 PLAC\n1 FORM A",
			Output: HeaderPlace{
				PlaceHierarchy: "A",
			},
		},
		{ // 3
			Input: "0 PLAC\n1 FORM\n",
			Err:   ErrContext{"HeaderPlace", cFORM, ErrInvalidLength{"PlaceHierarchy", "", 1, 120}},
		},
		{ // 4
			Input: "0 PLAC\n1 FORM A\n1 FORM B",
			Err:   ErrContext{"HeaderPlace", cFORM, ErrSingleMultiple},
		},
		{ // 5
			Input:   "0 PLAC\n1 FORM A\n1 FORM B",
			Options: []Option{AllowMoreThanAllowed},
			Output: HeaderPlace{
				PlaceHierarchy: "A",
			},
		},
		{ // 6
			Input: "0 PLAC\n1 FORM A\n1 UNKNOWN\n",
			Err:   ErrContext{"HeaderPlace", "UNKNOWN", ErrUnknownTag},
		},
		{ // 7
			Input: "0 PLAC\n1 FORM A\n1 _UNKNOWN\n",
			Output: HeaderPlace{
				PlaceHierarchy: "A",
			},
		},
		{ // 8
			Input:   "0 PLAC\n1 FORM A\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: HeaderPlace{
				PlaceHierarchy: "A",
			},
		},
	})
}

func TestFamily(t *testing.T) {
	test(t, []testData[Family]{
		{ // 1
			Input: "0 FAM\n",
			Err:   ErrContext{"Family", "xrefID", ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 2
			Input: "0 @ID@ FAM\n",
			Output: Family{
				ID: "ID",
			},
		},
		{ // 3
			Input: "0 @ID@ FAM\n1 ANUL\n",
			Output: Family{
				ID: "ID",
				Annulment: VerifiedFamilyEventDetail{
					FamilyEventDetail: FamilyEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 4
			Input: "0 @ID@ FAM\n1 ANUL N\n",
			Err:   ErrContext{"Family", cANUL, ErrContext{"VerifiedFamilyEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 5
			Input: "0 @ID@ FAM\n1 ANUL Y\n1 ANUL\n",
			Err:   ErrContext{"Family", cANUL, ErrSingleMultiple},
		},
		{ // 6
			Input:   "0 @ID@ FAM\n1 ANUL Y\n1 ANUL\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Family{
				ID: "ID",
				Annulment: VerifiedFamilyEventDetail{
					Verified: "Y",
					FamilyEventDetail: FamilyEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 7
			Input: "0 @ID@ FAM\n1 CENS\n",
			Output: Family{
				ID: "ID",
				Census: VerifiedFamilyEventDetail{
					FamilyEventDetail: FamilyEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 8
			Input: "0 @ID@ FAM\n1 CENS N\n",
			Err:   ErrContext{"Family", cCENS, ErrContext{"VerifiedFamilyEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 9
			Input: "0 @ID@ FAM\n1 CENS Y\n1 CENS\n",
			Err:   ErrContext{"Family", cCENS, ErrSingleMultiple},
		},
		{ // 10
			Input:   "0 @ID@ FAM\n1 CENS Y\n1 CENS\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Family{
				ID: "ID",
				Census: VerifiedFamilyEventDetail{
					Verified: "Y",
					FamilyEventDetail: FamilyEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 11
			Input: "0 @ID@ FAM\n1 DIV\n",
			Output: Family{
				ID: "ID",
				Divorce: VerifiedFamilyEventDetail{
					FamilyEventDetail: FamilyEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 12
			Input: "0 @ID@ FAM\n1 DIV N\n",
			Err:   ErrContext{"Family", cDIV, ErrContext{"VerifiedFamilyEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 13
			Input: "0 @ID@ FAM\n1 DIV Y\n1 DIV\n",
			Err:   ErrContext{"Family", cDIV, ErrSingleMultiple},
		},
		{ // 14
			Input:   "0 @ID@ FAM\n1 DIV Y\n1 DIV\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Family{
				ID: "ID",
				Divorce: VerifiedFamilyEventDetail{
					Verified: "Y",
					FamilyEventDetail: FamilyEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 15
			Input: "0 @ID@ FAM\n1 DIVF\n",
			Output: Family{
				ID: "ID",
				DivorceFiled: VerifiedFamilyEventDetail{
					FamilyEventDetail: FamilyEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 16
			Input: "0 @ID@ FAM\n1 DIVF N\n",
			Err:   ErrContext{"Family", cDIVF, ErrContext{"VerifiedFamilyEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 17
			Input: "0 @ID@ FAM\n1 DIVF Y\n1 DIVF\n",
			Err:   ErrContext{"Family", cDIVF, ErrSingleMultiple},
		},
		{ // 18
			Input:   "0 @ID@ FAM\n1 DIVF Y\n1 DIVF\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Family{
				ID: "ID",
				DivorceFiled: VerifiedFamilyEventDetail{
					Verified: "Y",
					FamilyEventDetail: FamilyEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 19
			Input: "0 @ID@ FAM\n1 ENGA\n",
			Output: Family{
				ID: "ID",
				Engagement: VerifiedFamilyEventDetail{
					FamilyEventDetail: FamilyEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 20
			Input: "0 @ID@ FAM\n1 ENGA N\n",
			Err:   ErrContext{"Family", cENGA, ErrContext{"VerifiedFamilyEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 21
			Input: "0 @ID@ FAM\n1 ENGA Y\n1 ENGA\n",
			Err:   ErrContext{"Family", cENGA, ErrSingleMultiple},
		},
		{ // 22
			Input:   "0 @ID@ FAM\n1 ENGA Y\n1 ENGA\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Family{
				ID: "ID",
				Engagement: VerifiedFamilyEventDetail{
					Verified: "Y",
					FamilyEventDetail: FamilyEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 23
			Input: "0 @ID@ FAM\n1 MARR\n",
			Output: Family{
				ID: "ID",
				Marriage: VerifiedFamilyEventDetail{
					FamilyEventDetail: FamilyEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 24
			Input: "0 @ID@ FAM\n1 MARR N\n",
			Err:   ErrContext{"Family", cMARR, ErrContext{"VerifiedFamilyEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 25
			Input: "0 @ID@ FAM\n1 MARR Y\n1 MARR\n",
			Err:   ErrContext{"Family", cMARR, ErrSingleMultiple},
		},
		{ // 26
			Input:   "0 @ID@ FAM\n1 MARR Y\n1 MARR\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Family{
				ID: "ID",
				Marriage: VerifiedFamilyEventDetail{
					Verified: "Y",
					FamilyEventDetail: FamilyEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 27
			Input: "0 @ID@ FAM\n1 MARB\n",
			Output: Family{
				ID: "ID",
				MarriageBann: VerifiedFamilyEventDetail{
					FamilyEventDetail: FamilyEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 28
			Input: "0 @ID@ FAM\n1 MARB N\n",
			Err:   ErrContext{"Family", cMARB, ErrContext{"VerifiedFamilyEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 29
			Input: "0 @ID@ FAM\n1 MARB Y\n1 MARB\n",
			Err:   ErrContext{"Family", cMARB, ErrSingleMultiple},
		},
		{ // 30
			Input:   "0 @ID@ FAM\n1 MARB Y\n1 MARB\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Family{
				ID: "ID",
				MarriageBann: VerifiedFamilyEventDetail{
					Verified: "Y",
					FamilyEventDetail: FamilyEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 31
			Input: "0 @ID@ FAM\n1 MARC\n",
			Output: Family{
				ID: "ID",
				MarriageContract: VerifiedFamilyEventDetail{
					FamilyEventDetail: FamilyEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 32
			Input: "0 @ID@ FAM\n1 MARC N\n",
			Err:   ErrContext{"Family", cMARC, ErrContext{"VerifiedFamilyEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 33
			Input: "0 @ID@ FAM\n1 MARC Y\n1 MARC\n",
			Err:   ErrContext{"Family", cMARC, ErrSingleMultiple},
		},
		{ // 34
			Input:   "0 @ID@ FAM\n1 MARC Y\n1 MARC\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Family{
				ID: "ID",
				MarriageContract: VerifiedFamilyEventDetail{
					Verified: "Y",
					FamilyEventDetail: FamilyEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 35
			Input: "0 @ID@ FAM\n1 MARL\n",
			Output: Family{
				ID: "ID",
				MarriageLicense: VerifiedFamilyEventDetail{
					FamilyEventDetail: FamilyEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 36
			Input: "0 @ID@ FAM\n1 MARL N\n",
			Err:   ErrContext{"Family", cMARL, ErrContext{"VerifiedFamilyEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 37
			Input: "0 @ID@ FAM\n1 MARL Y\n1 MARL\n",
			Err:   ErrContext{"Family", cMARL, ErrSingleMultiple},
		},
		{ // 38
			Input:   "0 @ID@ FAM\n1 MARL Y\n1 MARL\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Family{
				ID: "ID",
				MarriageLicense: VerifiedFamilyEventDetail{
					Verified: "Y",
					FamilyEventDetail: FamilyEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 39
			Input: "0 @ID@ FAM\n1 MARS\n",
			Output: Family{
				ID: "ID",
				MarriageSettlement: VerifiedFamilyEventDetail{
					FamilyEventDetail: FamilyEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 40
			Input: "0 @ID@ FAM\n1 MARS N\n",
			Err:   ErrContext{"Family", cMARS, ErrContext{"VerifiedFamilyEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 41
			Input: "0 @ID@ FAM\n1 MARS Y\n1 MARS\n",
			Err:   ErrContext{"Family", cMARS, ErrSingleMultiple},
		},
		{ // 42
			Input:   "0 @ID@ FAM\n1 MARS Y\n1 MARS\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Family{
				ID: "ID",
				MarriageSettlement: VerifiedFamilyEventDetail{
					Verified: "Y",
					FamilyEventDetail: FamilyEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 43
			Input: "0 @ID@ FAM\n1 EVEN\n",
			Output: Family{
				ID: "ID",
				Events: []FamilyEventDetail{
					{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 44
			Input: "0 @ID@ FAM\n1 EVEN\n2 TYPE\n",
			Err:   ErrContext{"Family", cEVEN, ErrContext{"EventDetail", cTYPE, ErrInvalidLength{"EventDescriptor", "", 1, 90}}},
		},
		{ // 45
			Input: "0 @ID@ FAM\n1 EVEN\n2 TYPE A\n1 EVEN\n2 TYPE B\n",
			Output: Family{
				ID: "ID",
				Events: []FamilyEventDetail{
					{
						EventDetail: EventDetail{
							Type:        "A",
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
					{
						EventDetail: EventDetail{
							Type:        "B",
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 46
			Input: "0 @ID@ FAM\n1 HUSB @A@\n",
			Output: Family{
				ID:      "ID",
				Husband: "A",
			},
		},
		{ // 47
			Input: "0 @ID@ FAM\n1 HUSB\n",
			Err:   ErrContext{"Family", cHUSB, ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 48
			Input: "0 @ID@ FAM\n1 HUSB @A@\n1 HUSB @B@",
			Err:   ErrContext{"Family", cHUSB, ErrSingleMultiple},
		},
		{ // 49
			Input:   "0 @ID@ FAM\n1 HUSB @A@\n1 HUSB @B@",
			Options: []Option{AllowMoreThanAllowed},
			Output: Family{
				ID:      "ID",
				Husband: "A",
			},
		},
		{ // 50
			Input: "0 @ID@ FAM\n1 WIFE @A@\n",
			Output: Family{
				ID:   "ID",
				Wife: "A",
			},
		},
		{ // 51
			Input: "0 @ID@ FAM\n1 WIFE\n",
			Err:   ErrContext{"Family", cWIFE, ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 52
			Input: "0 @ID@ FAM\n1 WIFE @A@\n1 WIFE @B@",
			Err:   ErrContext{"Family", cWIFE, ErrSingleMultiple},
		},
		{ // 53
			Input:   "0 @ID@ FAM\n1 WIFE @A@\n1 WIFE @B@",
			Options: []Option{AllowMoreThanAllowed},
			Output: Family{
				ID:   "ID",
				Wife: "A",
			},
		},
		{ // 54
			Input: "0 @ID@ FAM\n1 CHIL @A@",
			Output: Family{
				ID:       "ID",
				Children: []Xref{"A"},
			},
		},
		{ // 55
			Input: "0 @ID@ FAM\n1 CHIL\n",
			Err:   ErrContext{"Family", cCHIL, ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 56
			Input: "0 @ID@ FAM\n1 CHIL @A@\n1 CHIL @B@",
			Output: Family{
				ID:       "ID",
				Children: []Xref{"A", "B"},
			},
		},
		{ // 57
			Input: "0 @ID@ FAM\n1 NCHI 1\n",
			Output: Family{
				ID:          "ID",
				NumChildren: 1,
			},
		},
		{ // 58
			Input: "0 @ID@ FAM\n1 NCHI\n",
			Err:   ErrContext{"Family", cNCHI, ErrInvalidLength{"CountOfChildren", "", 1, 3}},
		},
		{ // 59
			Input: "0 @ID@ FAM\n1 NCHI 1\n1 NCHI 2",
			Err:   ErrContext{"Family", cNCHI, ErrSingleMultiple},
		},
		{ // 60
			Input:   "0 @ID@ FAM\n1 NCHI 1\n1 NCHI 2",
			Options: []Option{AllowMoreThanAllowed},
			Output: Family{
				ID:          "ID",
				NumChildren: 1,
			},
		},
		{ // 61
			Input: "0 @ID@ FAM\n1 SUBM @A@",
			Output: Family{
				ID:         "ID",
				Submitters: []Xref{"A"},
			},
		},
		{ // 62
			Input: "0 @ID@ FAM\n1 SUBM\n",
			Err:   ErrContext{"Family", cSUBM, ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 63
			Input: "0 @ID@ FAM\n1 SUBM @A@\n1 SUBM @B@",
			Output: Family{
				ID:         "ID",
				Submitters: []Xref{"A", "B"},
			},
		},
		{ // 64
			Input: "0 @ID@ FAM\n1 SLGS\n",
			Output: Family{
				ID: "ID",
				LDSSpouseSealing: []LDSSpouseSealing{
					{},
				},
			},
		},
		{ // 65
			Input: "0 @ID@ FAM\n1 SLGS\n2 STAT\n",
			Err:   ErrContext{"Family", cSLGS, ErrContext{"LDSSpouseSealing", cSTAT, ErrInvalidValue{"LDSSpouseSealingDateStatus", ""}}},
		},
		{ // 66
			Input: "0 @ID@ FAM\n1 SLGS\n1 SLGS\n",
			Output: Family{
				ID: "ID",
				LDSSpouseSealing: []LDSSpouseSealing{
					{},
					{},
				},
			},
		},
		{ // 67
			Input: "0 @ID@ FAM\n1 @A@ SOUR\n",
			Output: Family{
				ID: "ID",
				Sources: []SourceCitation{
					{
						Data: &SourceID{
							ID: "A",
						},
					},
				},
			},
		},
		{ // 68
			Input: "0 @ID@ FAM\n1 SOUR\n",
			Err:   ErrContext{"Family", cSOUR, ErrContext{"SourceText", "line_value", ErrInvalidLength{"SourceDescription", "", 1, 248}}},
		},
		{ // 69
			Input: "0 @ID@ FAM\n1 @A@ SOUR\n1 @B@ SOUR\n",
			Output: Family{
				ID: "ID",
				Sources: []SourceCitation{
					{
						Data: &SourceID{
							ID: "A",
						},
					},
					{
						Data: &SourceID{
							ID: "B",
						},
					},
				},
			},
		},
		{ // 70
			Input: "0 @ID@ FAM\n1 @A@ OBJE\n",
			Output: Family{
				ID: "ID",
				Multimedia: []MultimediaLink{
					{
						Data: &MultimediaLinkID{
							ID: "A",
						},
					},
				},
			},
		},
		{ // 71
			Input: "0 @ID@ FAM\n1 OBJE\n",
			Err:   ErrContext{"Family", cOBJE, ErrContext{"MultimediaLinkFile", cFORM, ErrRequiredMissing}},
		},
		{ // 72
			Input: "0 @ID@ FAM\n1 @A@ OBJE\n1 @B@ OBJE\n",
			Output: Family{
				ID: "ID",
				Multimedia: []MultimediaLink{
					{
						Data: &MultimediaLinkID{
							ID: "A",
						},
					},
					{
						Data: &MultimediaLinkID{
							ID: "B",
						},
					},
				},
			},
		},
		{ // 73
			Input: "0 @ID@ FAM\n1 @A@ NOTE\n",
			Output: Family{
				ID: "ID",
				Notes: []NoteStructure{
					{
						Data: &NoteID{
							ID: "A",
						},
					},
				},
			},
		},
		{ // 74
			Input: "0 @ID@ FAM\n1 NOTE\n",
			Err:   ErrContext{"Family", cNOTE, ErrContext{"NoteText", "line_value", ErrInvalidLength{"SubmitterText", "", 1, 248}}},
		},
		{ // 75
			Input: "0 @ID@ FAM\n1 @A@ NOTE\n1 @B@ NOTE\n",
			Output: Family{
				ID: "ID",
				Notes: []NoteStructure{
					{
						Data: &NoteID{
							ID: "A",
						},
					},
					{
						Data: &NoteID{
							ID: "B",
						},
					},
				},
			},
		},
		{ // 76
			Input: "0 @ID@ FAM\n1 UNKNOWN\n",
			Err:   ErrContext{"Family", "UNKNOWN", ErrUnknownTag},
		},
		{ // 77
			Input: "0 @ID@ FAM\n1 _UNKNOWN\n",
			Output: Family{
				ID: "ID",
			},
		},
		{ // 78
			Input:   "0 @ID@ FAM\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: Family{
				ID: "ID",
			},
		},
	})
}

func TestVerifiedFamilyEventDetail(t *testing.T) {
	test(t, []testData[VerifiedFamilyEventDetail]{
		{ // 1
			Input: "0 ANUL\n",
			Output: VerifiedFamilyEventDetail{
				FamilyEventDetail: FamilyEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 2
			Input: "0 ANUL Y\n",
			Output: VerifiedFamilyEventDetail{
				Verified: "Y",
				FamilyEventDetail: FamilyEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 3
			Input: "0 ANUL N\n",
			Err:   ErrContext{"VerifiedFamilyEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}},
		},
	})
}

func TestFamilyEventDetail(t *testing.T) {
	test(t, []testData[FamilyEventDetail]{
		{ // 1
			Input: "0 EVEN\n",
			Output: FamilyEventDetail{
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
		{ // 2
			Input: "0 EVEN\n1 HUSB\n2 AGE 20",
			Output: FamilyEventDetail{
				HusbandAge: AgeStructure{
					Age: "20",
				},
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
		{ // 3
			Input: "0 EVEN\n1 HUSB\n",
			Err:   ErrContext{"FamilyEventDetail", cHUSB, ErrContext{"AgeStructure", cAGE, ErrRequiredMissing}},
		},
		{ // 4
			Input: "0 EVEN\n1 HUSB\n2 AGE 20\n1 HUSB\n2 AGE 21",
			Err:   ErrContext{"FamilyEventDetail", cHUSB, ErrSingleMultiple},
		},
		{ // 5
			Input:   "0 EVEN\n1 HUSB\n2 AGE 20\n1 HUSB\n2 AGE 21",
			Options: []Option{AllowMoreThanAllowed},
			Output: FamilyEventDetail{
				HusbandAge: AgeStructure{
					Age: "20",
				},
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
		{ // 6
			Input: "0 EVEN\n1 WIFE\n2 AGE 20",
			Output: FamilyEventDetail{
				WifeAge: AgeStructure{
					Age: "20",
				},
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
		{ // 7
			Input: "0 EVEN\n1 WIFE\n",
			Err:   ErrContext{"FamilyEventDetail", cWIFE, ErrContext{"AgeStructure", cAGE, ErrRequiredMissing}},
		},
		{ // 8
			Input: "0 EVEN\n1 WIFE\n2 AGE 20\n1 WIFE\n2 AGE 21",
			Err:   ErrContext{"FamilyEventDetail", cWIFE, ErrSingleMultiple},
		},
		{ // 9
			Input:   "0 EVEN\n1 WIFE\n2 AGE 20\n1 WIFE\n2 AGE 21",
			Options: []Option{AllowMoreThanAllowed},
			Output: FamilyEventDetail{
				WifeAge: AgeStructure{
					Age: "20",
				},
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
	})
}

func TestAgeStructure(t *testing.T) {
	test(t, []testData[AgeStructure]{
		{ // 1
			Input: "0 HUSB\n1 AGE 20",
			Output: AgeStructure{
				Age: "20",
			},
		},
		{ // 2
			Input: "0 HUSB\n",
			Err:   ErrContext{"AgeStructure", cAGE, ErrRequiredMissing},
		},
		{ // 3
			Input: "0 HUSB\n1 AGE\n",
			Err:   ErrContext{"AgeStructure", cAGE, ErrInvalidLength{"AgeAtEvent", "", 1, 12}},
		},
		{ // 4
			Input: "0 HUSB\n1 AGE 20\n1 AGE 21",
			Err:   ErrContext{"AgeStructure", cAGE, ErrSingleMultiple},
		},
		{ // 5
			Input:   "0 HUSB\n1 AGE 20\n1 AGE 21",
			Options: []Option{AllowMoreThanAllowed},
			Output: AgeStructure{
				Age: "20",
			},
		},
		{ // 6
			Input: "0 HUSB\n1 AGE 20\n1 UNKNOWN\n",
			Err:   ErrContext{"AgeStructure", "UNKNOWN", ErrUnknownTag},
		},
		{ // 7
			Input: "0 HUSB\n1 AGE 20\n1 _UNKNOWN\n",
			Output: AgeStructure{
				Age: "20",
			},
		},
		{ // 8
			Input:   "0 HUSB\n1 AGE 20\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: AgeStructure{
				Age: "20",
			},
		},
	})
}

func TestIndividual(t *testing.T) {
	test(t, []testData[Individual]{
		{ // 1
			Input: "0 @A@ INDI\n",
			Output: Individual{
				ID: "A",
			},
		},
		{ // 2
			Input: "0 INDI\n",
			Err:   ErrContext{"Individual", "xrefID", ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 3
			Input: "0 @A@ INDI\n1 RESN locked",
			Output: Individual{
				ID:                "A",
				RestrictionNotice: "locked",
			},
		},
		{ // 4
			Input: "0 @A@ INDI\n1 RESN\n",
			Err:   ErrContext{"Individual", cRESN, ErrInvalidValue{"RestrictionNotice", ""}},
		},
		{ // 5
			Input: "0 @A@ INDI\n1 RESN locked\n1 RESN privacy",
			Err:   ErrContext{"Individual", cRESN, ErrSingleMultiple},
		},
		{ // 6
			Input:   "0 @A@ INDI\n1 RESN locked\n1 RESN privacy",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID:                "A",
				RestrictionNotice: "locked",
			},
		},
		{ // 7
			Input: "0 @A@ INDI\n1 NAME my name",
			Output: Individual{
				ID: "A",
				PersonalNameStructure: []PersonalNameStructure{
					{
						NamePersonal: "my name",
					},
				},
			},
		},
		{ // 8
			Input: "0 @A@ INDI\n1 NAME\n",
			Err:   ErrContext{"Individual", cNAME, ErrContext{"PersonalNameStructure", "line_value", ErrInvalidLength{"NamePersonal", "", 1, 120}}},
		},
		{ // 9
			Input: "0 @A@ INDI\n1 NAME my name\n1 NAME my other name",
			Output: Individual{
				ID: "A",
				PersonalNameStructure: []PersonalNameStructure{
					{
						NamePersonal: "my name",
					},
					{
						NamePersonal: "my other name",
					},
				},
			},
		},
		{ // 10
			Input: "0 @A@ INDI\n1 SEX M",
			Output: Individual{
				ID:     "A",
				Gender: "M",
			},
		},
		{ // 11
			Input: "0 @A@ INDI\n1 SEX\n",
			Err:   ErrContext{"Individual", cSEX, ErrInvalidLength{"SexValue", "", 1, 7}},
		},
		{ // 12
			Input: "0 @A@ INDI\n1 SEX M\n1 SEX F",
			Err:   ErrContext{"Individual", cSEX, ErrSingleMultiple},
		},
		{ // 13
			Input:   "0 @A@ INDI\n1 SEX M\n1 SEX F",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID:     "A",
				Gender: "M",
			},
		},
		{ // 14
			Input: "0 @A@ INDI\n1 BIRT\n",
			Output: Individual{
				ID: "A",
				Birth: VerifiedIndividualFamEventDetail{
					VerifiedEventDetail: VerifiedEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 15
			Input: "0 @A@ INDI\n1 BIRT N\n",
			Err:   ErrContext{"Individual", cBIRT, ErrContext{"VerifiedEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 16
			Input: "0 @A@ INDI\n1 BIRT Y\n1 BIRT\n",
			Err:   ErrContext{"Individual", cBIRT, ErrSingleMultiple},
		},
		{ // 17
			Input:   "0 @A@ INDI\n1 BIRT Y\n1 BIRT\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				Birth: VerifiedIndividualFamEventDetail{
					VerifiedEventDetail: VerifiedEventDetail{
						Verified: "Y",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 18
			Input: "0 @A@ INDI\n1 CHR\n",
			Output: Individual{
				ID: "A",
				Christening: VerifiedIndividualFamEventDetail{
					VerifiedEventDetail: VerifiedEventDetail{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 19
			Input: "0 @A@ INDI\n1 CHR N\n",
			Err:   ErrContext{"Individual", cCHR, ErrContext{"VerifiedEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 20
			Input: "0 @A@ INDI\n1 CHR Y\n1 CHR\n",
			Err:   ErrContext{"Individual", cCHR, ErrSingleMultiple},
		},
		{ // 21
			Input:   "0 @A@ INDI\n1 CHR Y\n1 CHR\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				Christening: VerifiedIndividualFamEventDetail{
					VerifiedEventDetail: VerifiedEventDetail{
						Verified: "Y",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 22
			Input: "0 @A@ INDI\n1 DEAT\n",
			Output: Individual{
				ID: "A",
				Death: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 23
			Input: "0 @A@ INDI\n1 DEAT N\n",
			Err:   ErrContext{"Individual", cDEAT, ErrContext{"VerifiedEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 24
			Input: "0 @A@ INDI\n1 DEAT Y\n1 DEAT\n",
			Err:   ErrContext{"Individual", cDEAT, ErrSingleMultiple},
		},
		{ // 25
			Input:   "0 @A@ INDI\n1 DEAT Y\n1 DEAT\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				Death: VerifiedEventDetail{
					Verified: "Y",
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 26
			Input: "0 @A@ INDI\n1 BURI\n",
			Output: Individual{
				ID: "A",
				Buried: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 27
			Input: "0 @A@ INDI\n1 BURI N\n",
			Err:   ErrContext{"Individual", cBURI, ErrContext{"VerifiedEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 28
			Input: "0 @A@ INDI\n1 BURI Y\n1 BURI\n",
			Err:   ErrContext{"Individual", cBURI, ErrSingleMultiple},
		},
		{ // 29
			Input:   "0 @A@ INDI\n1 BURI Y\n1 BURI\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				Buried: VerifiedEventDetail{
					Verified: "Y",
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 30
			Input: "0 @A@ INDI\n1 CREM\n",
			Output: Individual{
				ID: "A",
				Cremation: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 31
			Input: "0 @A@ INDI\n1 CREM N\n",
			Err:   ErrContext{"Individual", cCREM, ErrContext{"VerifiedEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 32
			Input: "0 @A@ INDI\n1 CREM Y\n1 CREM\n",
			Err:   ErrContext{"Individual", cCREM, ErrSingleMultiple},
		},
		{ // 33
			Input:   "0 @A@ INDI\n1 CREM Y\n1 CREM\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				Cremation: VerifiedEventDetail{
					Verified: "Y",
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 34
			Input: "0 @A@ INDI\n1 ADOP\n",
			Output: Individual{
				ID: "A",
				Adoption: AdoptionEvent{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 35
			Input: "0 @A@ INDI\n1 ADOP\n2 FAMC\n",
			Err:   ErrContext{"Individual", cADOP, ErrContext{"AdoptionEvent", cFAMC, ErrContext{"AdoptionReference", "xrefID", ErrInvalidLength{"Xref", "", 1, 22}}}},
		},
		{ // 36
			Input: "0 @A@ INDI\n1 ADOP\n2 @A@ FAMC\n1 ADOP\n2 @B@ FAMC\n",
			Err:   ErrContext{"Individual", cADOP, ErrSingleMultiple},
		},
		{ // 37
			Input:   "0 @A@ INDI\n1 ADOP\n2 @A@ FAMC\n1 ADOP\n2 @B@ FAMC\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				Adoption: AdoptionEvent{
					Family: AdoptionReference{
						ID: "A",
					},
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 38
			Input: "0 @A@ INDI\n1 BAPM\n",
			Output: Individual{
				ID: "A",
				Baptism: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 39
			Input: "0 @A@ INDI\n1 BAPM N\n",
			Err:   ErrContext{"Individual", cBAPM, ErrContext{"VerifiedEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 40
			Input: "0 @A@ INDI\n1 BAPM Y\n1 BAPM\n",
			Err:   ErrContext{"Individual", cBAPM, ErrSingleMultiple},
		},
		{ // 41
			Input:   "0 @A@ INDI\n1 BAPM Y\n1 BAPM\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				Baptism: VerifiedEventDetail{
					Verified: "Y",
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 42
			Input: "0 @A@ INDI\n1 BARM\n",
			Output: Individual{
				ID: "A",
				BarMitzvah: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 43
			Input: "0 @A@ INDI\n1 BARM N\n",
			Err:   ErrContext{"Individual", cBARM, ErrContext{"VerifiedEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 44
			Input: "0 @A@ INDI\n1 BARM Y\n1 BARM\n",
			Err:   ErrContext{"Individual", cBARM, ErrSingleMultiple},
		},
		{ // 45
			Input:   "0 @A@ INDI\n1 BARM Y\n1 BARM\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				BarMitzvah: VerifiedEventDetail{
					Verified: "Y",
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 46
			Input: "0 @A@ INDI\n1 BASM\n",
			Output: Individual{
				ID: "A",
				BasMitzvah: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 47
			Input: "0 @A@ INDI\n1 BASM N\n",
			Err:   ErrContext{"Individual", cBASM, ErrContext{"VerifiedEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 48
			Input: "0 @A@ INDI\n1 BASM Y\n1 BASM\n",
			Err:   ErrContext{"Individual", cBASM, ErrSingleMultiple},
		},
		{ // 49
			Input:   "0 @A@ INDI\n1 BASM Y\n1 BASM\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				BasMitzvah: VerifiedEventDetail{
					Verified: "Y",
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 50
			Input: "0 @A@ INDI\n1 BLES\n",
			Output: Individual{
				ID: "A",
				Blessing: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 51
			Input: "0 @A@ INDI\n1 BLES N\n",
			Err:   ErrContext{"Individual", cBLES, ErrContext{"VerifiedEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 52
			Input: "0 @A@ INDI\n1 BLES Y\n1 BLES\n",
			Err:   ErrContext{"Individual", cBLES, ErrSingleMultiple},
		},
		{ // 53
			Input:   "0 @A@ INDI\n1 BLES Y\n1 BLES\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				Blessing: VerifiedEventDetail{
					Verified: "Y",
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 54
			Input: "0 @A@ INDI\n1 CHRA\n",
			Output: Individual{
				ID: "A",
				AdultChristening: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 55
			Input: "0 @A@ INDI\n1 CHRA N\n",
			Err:   ErrContext{"Individual", cCHRA, ErrContext{"VerifiedEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 56
			Input: "0 @A@ INDI\n1 CHRA Y\n1 CHRA\n",
			Err:   ErrContext{"Individual", cCHRA, ErrSingleMultiple},
		},
		{ // 57
			Input:   "0 @A@ INDI\n1 CHRA Y\n1 CHRA\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				AdultChristening: VerifiedEventDetail{
					Verified: "Y",
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 58
			Input: "0 @A@ INDI\n1 CONF\n",
			Output: Individual{
				ID: "A",
				Confirmation: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 59
			Input: "0 @A@ INDI\n1 CONF N\n",
			Err:   ErrContext{"Individual", cCONF, ErrContext{"VerifiedEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 60
			Input: "0 @A@ INDI\n1 CONF Y\n1 CONF\n",
			Err:   ErrContext{"Individual", cCONF, ErrSingleMultiple},
		},
		{ // 61
			Input:   "0 @A@ INDI\n1 CONF Y\n1 CONF\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				Confirmation: VerifiedEventDetail{
					Verified: "Y",
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 62
			Input: "0 @A@ INDI\n1 FCOM\n",
			Output: Individual{
				ID: "A",
				FirstCommunion: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 63
			Input: "0 @A@ INDI\n1 FCOM N\n",
			Err:   ErrContext{"Individual", cFCOM, ErrContext{"VerifiedEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 64
			Input: "0 @A@ INDI\n1 FCOM Y\n1 FCOM\n",
			Err:   ErrContext{"Individual", cFCOM, ErrSingleMultiple},
		},
		{ // 65
			Input:   "0 @A@ INDI\n1 FCOM Y\n1 FCOM\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				FirstCommunion: VerifiedEventDetail{
					Verified: "Y",
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 66
			Input: "0 @A@ INDI\n1 ORDN\n",
			Output: Individual{
				ID: "A",
				Ordination: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 67
			Input: "0 @A@ INDI\n1 ORDN N\n",
			Err:   ErrContext{"Individual", cORDN, ErrContext{"VerifiedEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 68
			Input: "0 @A@ INDI\n1 ORDN Y\n1 ORDN\n",
			Err:   ErrContext{"Individual", cORDN, ErrSingleMultiple},
		},
		{ // 69
			Input:   "0 @A@ INDI\n1 ORDN Y\n1 ORDN\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				Ordination: VerifiedEventDetail{
					Verified: "Y",
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 70
			Input: "0 @A@ INDI\n1 NATU\n",
			Output: Individual{
				ID: "A",
				Naturalization: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 71
			Input: "0 @A@ INDI\n1 NATU N\n",
			Err:   ErrContext{"Individual", cNATU, ErrContext{"VerifiedEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 72
			Input: "0 @A@ INDI\n1 NATU Y\n1 NATU\n",
			Err:   ErrContext{"Individual", cNATU, ErrSingleMultiple},
		},
		{ // 73
			Input:   "0 @A@ INDI\n1 NATU Y\n1 NATU\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				Naturalization: VerifiedEventDetail{
					Verified: "Y",
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 74
			Input: "0 @A@ INDI\n1 EMIG\n",
			Output: Individual{
				ID: "A",
				Emigrated: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 75
			Input: "0 @A@ INDI\n1 EMIG N\n",
			Err:   ErrContext{"Individual", cEMIG, ErrContext{"VerifiedEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 76
			Input: "0 @A@ INDI\n1 EMIG Y\n1 EMIG\n",
			Err:   ErrContext{"Individual", cEMIG, ErrSingleMultiple},
		},
		{ // 77
			Input:   "0 @A@ INDI\n1 EMIG Y\n1 EMIG\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				Emigrated: VerifiedEventDetail{
					Verified: "Y",
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 78
			Input: "0 @A@ INDI\n1 IMMI\n",
			Output: Individual{
				ID: "A",
				Immigrated: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 79
			Input: "0 @A@ INDI\n1 IMMI N\n",
			Err:   ErrContext{"Individual", cIMMI, ErrContext{"VerifiedEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 80
			Input: "0 @A@ INDI\n1 IMMI Y\n1 IMMI\n",
			Err:   ErrContext{"Individual", cIMMI, ErrSingleMultiple},
		},
		{ // 81
			Input:   "0 @A@ INDI\n1 IMMI Y\n1 IMMI\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				Immigrated: VerifiedEventDetail{
					Verified: "Y",
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 82
			Input: "0 @A@ INDI\n1 CENS\n",
			Output: Individual{
				ID: "A",
				Census: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 83
			Input: "0 @A@ INDI\n1 CENS N\n",
			Err:   ErrContext{"Individual", cCENS, ErrContext{"VerifiedEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 84
			Input: "0 @A@ INDI\n1 CENS Y\n1 CENS\n",
			Err:   ErrContext{"Individual", cCENS, ErrSingleMultiple},
		},
		{ // 85
			Input:   "0 @A@ INDI\n1 CENS Y\n1 CENS\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				Census: VerifiedEventDetail{
					Verified: "Y",
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 86
			Input: "0 @A@ INDI\n1 PROB\n",
			Output: Individual{
				ID: "A",
				Probate: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 87
			Input: "0 @A@ INDI\n1 PROB N\n",
			Err:   ErrContext{"Individual", cPROB, ErrContext{"VerifiedEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 88
			Input: "0 @A@ INDI\n1 PROB Y\n1 PROB\n",
			Err:   ErrContext{"Individual", cPROB, ErrSingleMultiple},
		},
		{ // 89
			Input:   "0 @A@ INDI\n1 PROB Y\n1 PROB\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				Probate: VerifiedEventDetail{
					Verified: "Y",
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 90
			Input: "0 @A@ INDI\n1 WILL\n",
			Output: Individual{
				ID: "A",
				Will: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 91
			Input: "0 @A@ INDI\n1 WILL N\n",
			Err:   ErrContext{"Individual", cWILL, ErrContext{"VerifiedEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 92
			Input: "0 @A@ INDI\n1 WILL Y\n1 WILL\n",
			Err:   ErrContext{"Individual", cWILL, ErrSingleMultiple},
		},
		{ // 93
			Input:   "0 @A@ INDI\n1 WILL Y\n1 WILL\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				Will: VerifiedEventDetail{
					Verified: "Y",
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 94
			Input: "0 @A@ INDI\n1 GRAD\n",
			Output: Individual{
				ID: "A",
				Graduated: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 95
			Input: "0 @A@ INDI\n1 GRAD N\n",
			Err:   ErrContext{"Individual", cGRAD, ErrContext{"VerifiedEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 96
			Input: "0 @A@ INDI\n1 GRAD Y\n1 GRAD\n",
			Err:   ErrContext{"Individual", cGRAD, ErrSingleMultiple},
		},
		{ // 97
			Input:   "0 @A@ INDI\n1 GRAD Y\n1 GRAD\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				Graduated: VerifiedEventDetail{
					Verified: "Y",
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 98
			Input: "0 @A@ INDI\n1 RETI\n",
			Output: Individual{
				ID: "A",
				Retired: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 99
			Input: "0 @A@ INDI\n1 RETI N\n",
			Err:   ErrContext{"Individual", cRETI, ErrContext{"VerifiedEventDetail", "line_value", ErrInvalidValue{"Verified", "N"}}},
		},
		{ // 100
			Input: "0 @A@ INDI\n1 RETI Y\n1 RETI\n",
			Err:   ErrContext{"Individual", cRETI, ErrSingleMultiple},
		},
		{ // 101
			Input:   "0 @A@ INDI\n1 RETI Y\n1 RETI\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "A",
				Retired: VerifiedEventDetail{
					Verified: "Y",
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 102
			Input: "0 @A@ INDI\n1 EVEN\n",
			Output: Individual{
				ID: "A",
				Events: []IndividualEventDetail{
					{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 103
			Input: "0 @A@ INDI\n1 EVEN\n2 TYPE\n",
			Err:   ErrContext{"Individual", cEVEN, ErrContext{"EventDetail", cTYPE, ErrInvalidLength{"EventDescriptor", "", 1, 90}}},
		},
		{ // 104
			Input: "0 @A@ INDI\n1 EVEN\n2 TYPE A\n1 EVEN\n2 TYPE B",
			Output: Individual{
				ID: "A",
				Events: []IndividualEventDetail{
					{
						EventDetail: EventDetail{
							Type:        "A",
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
					{
						EventDetail: EventDetail{
							Type:        "B",
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 105
			Input: "0 @A@ INDI\n1 CAST A\n",
			Output: Individual{
				ID: "A",
				Caste: []CasteEvent{
					{
						CasteName: "A",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 106
			Input: "0 @A@ INDI\n1 CAST\n",
			Err:   ErrContext{"Individual", cCAST, ErrContext{"CasteEvent", "line_value", ErrInvalidLength{"CasteName", "", 1, 90}}},
		},
		{ // 107
			Input: "0 @A@ INDI\n1 CAST A\n1 CAST B",
			Output: Individual{
				ID: "A",
				Caste: []CasteEvent{
					{
						CasteName: "A",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
					{
						CasteName: "B",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 108
			Input: "0 @A@ INDI\n1 DSCR A\n",
			Output: Individual{
				ID: "A",
				Description: []DescriptionEvent{
					{
						PhysicalDescription: "A",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 109
			Input: "0 @A@ INDI\n1 DSCR\n",
			Err:   ErrContext{"Individual", cDSCR, ErrContext{"DescriptionEvent", "line_value", ErrInvalidLength{"PhysicalDescription", "", 1, 248}}},
		},
		{ // 110
			Input: "0 @A@ INDI\n1 DSCR A\n1 DSCR B",
			Output: Individual{
				ID: "A",
				Description: []DescriptionEvent{
					{
						PhysicalDescription: "A",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
					{
						PhysicalDescription: "B",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 111
			Input: "0 @A@ INDI\n1 EDUC A\n",
			Output: Individual{
				ID: "A",
				ScholasticAchievement: []ScholasticEvent{
					{
						ScholasticAchievement: "A",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 112
			Input: "0 @A@ INDI\n1 EDUC\n",
			Err:   ErrContext{"Individual", cEDUC, ErrContext{"ScholasticEvent", "line_value", ErrInvalidLength{"ScholasticAchievement", "", 1, 248}}},
		},
		{ // 113
			Input: "0 @A@ INDI\n1 EDUC A\n1 EDUC B",
			Output: Individual{
				ID: "A",
				ScholasticAchievement: []ScholasticEvent{
					{
						ScholasticAchievement: "A",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
					{
						ScholasticAchievement: "B",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 114
			Input: "0 @A@ INDI\n1 IDNO A\n",
			Output: Individual{
				ID: "A",
				NationalID: []NationalIDEvent{
					{
						NationalIDNumber: "A",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 115
			Input: "0 @A@ INDI\n1 IDNO\n",
			Err:   ErrContext{"Individual", cIDNO, ErrContext{"NationalIDEvent", "line_value", ErrInvalidLength{"NationalIDNumber", "", 1, 30}}},
		},
		{ // 116
			Input: "0 @A@ INDI\n1 IDNO A\n1 IDNO B",
			Output: Individual{
				ID: "A",
				NationalID: []NationalIDEvent{
					{
						NationalIDNumber: "A",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
					{
						NationalIDNumber: "B",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 117
			Input: "0 @A@ INDI\n1 NATI A\n",
			Output: Individual{
				ID: "A",
				NationalTribalOrigin: []NationalOriginEvent{
					{
						NationalOrTribalOrigin: "A",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 118
			Input: "0 @A@ INDI\n1 NATI\n",
			Err:   ErrContext{"Individual", cNATI, ErrContext{"NationalOriginEvent", "line_value", ErrInvalidLength{"NationalOrTribalOrigin", "", 1, 120}}},
		},
		{ // 119
			Input: "0 @A@ INDI\n1 NATI A\n1 NATI B",
			Output: Individual{
				ID: "A",
				NationalTribalOrigin: []NationalOriginEvent{
					{
						NationalOrTribalOrigin: "A",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
					{
						NationalOrTribalOrigin: "B",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 120
			Input: "0 @A@ INDI\n1 NCHI 1\n",
			Output: Individual{
				ID: "A",
				CountOfChildren: []ChildrenEvent{
					{
						CountOfChildren: 1,
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 121
			Input: "0 @A@ INDI\n1 NCHI\n",
			Err:   ErrContext{"Individual", cNCHI, ErrContext{"ChildrenEvent", "line_value", ErrInvalidLength{"CountOfChildren", "", 1, 3}}},
		},
		{ // 122
			Input: "0 @A@ INDI\n1 NCHI 1\n1 NCHI 2",
			Output: Individual{
				ID: "A",
				CountOfChildren: []ChildrenEvent{
					{
						CountOfChildren: 1,
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
					{
						CountOfChildren: 2,
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 123
			Input: "0 @A@ INDI\n1 NMR 1\n",
			Output: Individual{
				ID: "A",
				CountOfMarriages: []MarriagesEvent{
					{
						CountOfMarriages: 1,
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 124
			Input: "0 @A@ INDI\n1 NMR\n",
			Err:   ErrContext{"Individual", cNMR, ErrContext{"MarriagesEvent", "line_value", ErrInvalidLength{"CountOfMarriages", "", 1, 3}}},
		},
		{ // 125
			Input: "0 @A@ INDI\n1 NMR 1\n1 NMR 2",
			Output: Individual{
				ID: "A",
				CountOfMarriages: []MarriagesEvent{
					{
						CountOfMarriages: 1,
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
					{
						CountOfMarriages: 2,
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 126
			Input: "0 @A@ INDI\n1 OCCU A\n",
			Output: Individual{
				ID: "A",
				Occupation: []OccupationEvent{
					{
						Occupation: "A",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 127
			Input: "0 @A@ INDI\n1 OCCU\n",
			Err:   ErrContext{"Individual", cOCCU, ErrContext{"OccupationEvent", "line_value", ErrInvalidLength{"Occupation", "", 1, 90}}},
		},
		{ // 128
			Input: "0 @A@ INDI\n1 OCCU A\n1 OCCU B",
			Output: Individual{
				ID: "A",
				Occupation: []OccupationEvent{
					{
						Occupation: "A",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
					{
						Occupation: "B",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 129
			Input: "0 @A@ INDI\n1 PROP A\n",
			Output: Individual{
				ID: "A",
				Possessions: []PossessionEvent{
					{
						Possessions: "A",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 130
			Input: "0 @A@ INDI\n1 PROP\n",
			Err:   ErrContext{"Individual", cPROP, ErrContext{"PossessionEvent", "line_value", ErrInvalidLength{"Possessions", "", 1, 248}}},
		},
		{ // 131
			Input: "0 @A@ INDI\n1 PROP A\n1 PROP B",
			Output: Individual{
				ID: "A",
				Possessions: []PossessionEvent{
					{
						Possessions: "A",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
					{
						Possessions: "B",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 132
			Input: "0 @A@ INDI\n1 RELI A\n",
			Output: Individual{
				ID: "A",
				ReligiousAffiliation: []ReligiousEvent{
					{
						ReligiousAffiliation: "A",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 133
			Input: "0 @A@ INDI\n1 RELI\n",
			Err:   ErrContext{"Individual", cRELI, ErrContext{"ReligiousEvent", "line_value", ErrInvalidLength{"ReligiousAffiliation", "", 1, 90}}},
		},
		{ // 134
			Input: "0 @A@ INDI\n1 RELI A\n1 RELI B",
			Output: Individual{
				ID: "A",
				ReligiousAffiliation: []ReligiousEvent{
					{
						ReligiousAffiliation: "A",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
					{
						ReligiousAffiliation: "B",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 135
			Input: "0 @A@ INDI\n1 RESI A\n",
			Output: Individual{
				ID: "A",
				Residences: []ResidenceEvent{
					{
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 136
			Input: "0 @A@ INDI\n1 RESI\n2 TYPE\n",
			Err:   ErrContext{"Individual", cRESI, ErrContext{"EventDetail", cTYPE, ErrInvalidLength{"EventDescriptor", "", 1, 90}}},
		},
		{ // 137
			Input: "0 @A@ INDI\n1 RESI\n2 TYPE A\n1 RESI\n2 TYPE B",
			Output: Individual{
				ID: "A",
				Residences: []ResidenceEvent{
					{
						EventDetail: EventDetail{
							Type:        "A",
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
					{
						EventDetail: EventDetail{
							Type:        "B",
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 138
			Input: "0 @A@ INDI\n1 SSN AAAAAAAAA\n",
			Output: Individual{
				ID: "A",
				SocialSecurity: []SSNEvent{
					{
						SocialSecurityNumber: "AAAAAAAAA",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 139
			Input: "0 @A@ INDI\n1 SSN\n",
			Err:   ErrContext{"Individual", cSSN, ErrContext{"SSNEvent", "line_value", ErrInvalidLength{"SocialSecurityNumber", "", 9, 11}}},
		},
		{ // 140
			Input: "0 @A@ INDI\n1 SSN AAAAAAAAA\n1 SSN BBBBBBBBB",
			Output: Individual{
				ID: "A",
				SocialSecurity: []SSNEvent{
					{
						SocialSecurityNumber: "AAAAAAAAA",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
					{
						SocialSecurityNumber: "BBBBBBBBB",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 141
			Input: "0 @A@ INDI\n1 TITL A\n",
			Output: Individual{
				ID: "A",
				NobilityTypeTitle: []NobilityEvent{
					{
						NobilityTypeTitle: "A",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 142
			Input: "0 @A@ INDI\n1 TITL\n",
			Err:   ErrContext{"Individual", cTITL, ErrContext{"NobilityEvent", "line_value", ErrInvalidLength{"NobilityTypeTitle", "", 1, 120}}},
		},
		{ // 143
			Input: "0 @A@ INDI\n1 TITL A\n1 TITL B",
			Output: Individual{
				ID: "A",
				NobilityTypeTitle: []NobilityEvent{
					{
						NobilityTypeTitle: "A",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
					{
						NobilityTypeTitle: "B",
						EventDetail: EventDetail{
							PhoneNumber: make([]PhoneNumber, 0, 3),
						},
					},
				},
			},
		},
		{ // 144
			Input: "0 @A@ INDI\n1 FAMC @B@\n",
			Output: Individual{
				ID: "A",
				ChildOf: []ChildToFamilyLink{
					{
						ID: "B",
					},
				},
			},
		},
		{ // 145
			Input: "0 @A@ INDI\n1 FAMC\n",
			Err:   ErrContext{"Individual", cFAMC, ErrContext{"ChildToFamilyLink", "line_value", ErrInvalidLength{"Xref", "", 1, 22}}},
		},
		{ // 146
			Input: "0 @A@ INDI\n1 FAMC @B@\n1 FAMC @C@",
			Output: Individual{
				ID: "A",
				ChildOf: []ChildToFamilyLink{
					{
						ID: "B",
					},
					{
						ID: "C",
					},
				},
			},
		},
		{ // 147
			Input: "0 @A@ INDI\n1 FAMS @B@\n",
			Output: Individual{
				ID: "A",
				SpouseOf: []SpouseToFamilyLink{
					{
						ID: "B",
					},
				},
			},
		},
		{ // 148
			Input: "0 @A@ INDI\n1 FAMS\n",
			Err:   ErrContext{"Individual", cFAMS, ErrContext{"SpouseToFamilyLink", "line_value", ErrInvalidLength{"Xref", "", 1, 22}}},
		},
		{ // 149
			Input: "0 @A@ INDI\n1 FAMS @B@\n1 FAMS @C@",
			Output: Individual{
				ID: "A",
				SpouseOf: []SpouseToFamilyLink{
					{
						ID: "B",
					},
					{
						ID: "C",
					},
				},
			},
		},
		{ // 150
			Input: "0 @A@ INDI\n1 SUBM @B@\n",
			Output: Individual{
				ID: "A",
				Submitters: []Xref{
					"B",
				},
			},
		},
		{ // 151
			Input: "0 @A@ INDI\n1 SUBM\n",
			Err:   ErrContext{"Individual", cSUBM, ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 152
			Input: "0 @A@ INDI\n1 SUBM @C@\n1 SUBM @D@",
			Output: Individual{
				ID: "A",
				Submitters: []Xref{
					"C",
					"D",
				},
			},
		},
		{ // 153
			Input: "0 @A@ INDI\n1 @B@ ASSO\n2 TYPE FAM\n2 RELA C",
			Output: Individual{
				ID: "A",
				Associations: []AssociationStructure{
					{
						ID:         "B",
						RecordType: "FAM",
						Relation:   "C",
					},
				},
			},
		},
		{ // 154
			Input: "0 @A@ INDI\n1 ASSO\n",
			Err:   ErrContext{"Individual", cASSO, ErrContext{"AssociationStructure", "xrefID", ErrInvalidLength{"Xref", "", 1, 22}}},
		},
		{ // 155
			Input: "0 @A@ INDI\n1 @B@ ASSO\n2 TYPE FAM\n2 RELA C\n1 @C@ ASSO\n2 TYPE FAM\n2 RELA D",
			Output: Individual{
				ID: "A",
				Associations: []AssociationStructure{
					{
						ID:         "B",
						RecordType: "FAM",
						Relation:   "C",
					},
					{
						ID:         "C",
						RecordType: "FAM",
						Relation:   "D",
					},
				},
			},
		},
		{ // 156
			Input: "0 @A@ INDI\n1 ALIA @B@\n",
			Output: Individual{
				ID: "A",
				Aliases: []Xref{
					"B",
				},
			},
		},
		{ // 157
			Input: "0 @A@ INDI\n1 ALIA\n",
			Err:   ErrContext{"Individual", cALIA, ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 158
			Input: "0 @A@ INDI\n1 ALIA @B@\n1 ALIA @C@",
			Output: Individual{
				ID: "A",
				Aliases: []Xref{
					"B",
					"C",
				},
			},
		},
		{ // 159
			Input: "0 @A@ INDI\n1 ANCI @B@\n",
			Output: Individual{
				ID: "A",
				AncestorInterest: []Xref{
					"B",
				},
			},
		},
		{ // 160
			Input: "0 @A@ INDI\n1 ANCI\n",
			Err:   ErrContext{"Individual", cANCI, ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 161
			Input: "0 @A@ INDI\n1 ANCI @B@\n1 ANCI @C@",
			Output: Individual{
				ID: "A",
				AncestorInterest: []Xref{
					"B",
					"C",
				},
			},
		},
		{ // 162
			Input: "0 @A@ INDI\n1 DESI @B@\n",
			Output: Individual{
				ID: "A",
				DescendentInterest: []Xref{
					"B",
				},
			},
		},
		{ // 163
			Input: "0 @A@ INDI\n1 DESI\n",
			Err:   ErrContext{"Individual", cDESI, ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 164
			Input: "0 @A@ INDI\n1 DESI @B@\n1 DESI @C@",
			Output: Individual{
				ID: "A",
				DescendentInterest: []Xref{
					"B",
					"C",
				},
			},
		},
		{ // 165
			Input: "0 @ID@ INDI\n1 @A@ SOUR\n",
			Output: Individual{
				ID: "ID",
				Sources: []SourceCitation{
					{
						Data: &SourceID{
							ID: "A",
						},
					},
				},
			},
		},
		{ // 166
			Input: "0 @ID@ INDI\n1 SOUR\n",
			Err:   ErrContext{"Individual", cSOUR, ErrContext{"SourceText", "line_value", ErrInvalidLength{"SourceDescription", "", 1, 248}}},
		},
		{ // 167
			Input: "0 @ID@ INDI\n1 @A@ SOUR\n1 @B@ SOUR\n",
			Output: Individual{
				ID: "ID",
				Sources: []SourceCitation{
					{
						Data: &SourceID{
							ID: "A",
						},
					},
					{
						Data: &SourceID{
							ID: "B",
						},
					},
				},
			},
		},
		{ // 168
			Input: "0 @ID@ INDI\n1 @A@ OBJE\n",
			Output: Individual{
				ID: "ID",
				Multimedia: []MultimediaLink{
					{
						Data: &MultimediaLinkID{
							ID: "A",
						},
					},
				},
			},
		},
		{ // 169
			Input: "0 @ID@ INDI\n1 OBJE\n",
			Err:   ErrContext{"Individual", cOBJE, ErrContext{"MultimediaLinkFile", cFORM, ErrRequiredMissing}},
		},
		{ // 170
			Input: "0 @ID@ INDI\n1 @A@ OBJE\n1 @B@ OBJE\n",
			Output: Individual{
				ID: "ID",
				Multimedia: []MultimediaLink{
					{
						Data: &MultimediaLinkID{
							ID: "A",
						},
					},
					{
						Data: &MultimediaLinkID{
							ID: "B",
						},
					},
				},
			},
		},
		{ // 171
			Input: "0 @ID@ INDI\n1 @A@ NOTE\n",
			Output: Individual{
				ID: "ID",
				Notes: []NoteStructure{
					{
						Data: &NoteID{
							ID: "A",
						},
					},
				},
			},
		},
		{ // 172
			Input: "0 @ID@ INDI\n1 NOTE\n",
			Err:   ErrContext{"Individual", cNOTE, ErrContext{"NoteText", "line_value", ErrInvalidLength{"SubmitterText", "", 1, 248}}},
		},
		{ // 173
			Input: "0 @ID@ INDI\n1 @A@ NOTE\n1 @B@ NOTE\n",
			Output: Individual{
				ID: "ID",
				Notes: []NoteStructure{
					{
						Data: &NoteID{
							ID: "A",
						},
					},
					{
						Data: &NoteID{
							ID: "B",
						},
					},
				},
			},
		},
		{ // 174
			Input: "0 @ID@ INDI\n1 RFN A\n",
			Output: Individual{
				ID:              "ID",
				PermanentRecord: "A",
			},
		},
		{ // 175
			Input: "0 @ID@ INDI\n1 RFN\n",
			Err:   ErrContext{"Individual", cRFN, ErrInvalidLength{"PermanentRecordFileNumber", "", 1, 90}},
		},
		{ // 176
			Input: "0 @ID@ INDI\n1 RFN A\n1 RFN B",
			Err:   ErrContext{"Individual", cRFN, ErrSingleMultiple},
		},
		{ // 177
			Input:   "0 @ID@ INDI\n1 RFN A\n1 RFN B",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID:              "ID",
				PermanentRecord: "A",
			},
		},
		{ // 178
			Input: "0 @ID@ INDI\n1 AFN A\n",
			Output: Individual{
				ID:                  "ID",
				AncestralFileNumber: "A",
			},
		},
		{ // 179
			Input: "0 @ID@ INDI\n1 AFN\n",
			Err:   ErrContext{"Individual", cAFN, ErrInvalidLength{"AncestralFileNumber", "", 1, 12}},
		},
		{ // 180
			Input: "0 @ID@ INDI\n1 AFN A\n1 AFN B",
			Err:   ErrContext{"Individual", cAFN, ErrSingleMultiple},
		},
		{ // 181
			Input:   "0 @ID@ INDI\n1 AFN A\n1 AFN B",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID:                  "ID",
				AncestralFileNumber: "A",
			},
		},
		{ // 182
			Input: "0 @ID@ INDI\n1 REFN A\n",
			Output: Individual{
				ID: "ID",
				UserReferences: []UserReferenceStructure{
					{
						UserReferenceNumber: "A",
					},
				},
			},
		},
		{ // 183
			Input: "0 @ID@ INDI\n1 REFN\n",
			Err:   ErrContext{"Individual", cREFN, ErrContext{"UserReferenceStructure", "line_value", ErrInvalidLength{"UserReferenceNumber", "", 1, 20}}},
		},
		{ // 184
			Input: "0 @ID@ INDI\n1 REFN A\n1 REFN B\n",
			Output: Individual{
				ID: "ID",
				UserReferences: []UserReferenceStructure{
					{
						UserReferenceNumber: "A",
					},
					{
						UserReferenceNumber: "B",
					},
				},
			},
		},
		{ // 185
			Input: "0 @ID@ INDI\n1 RIN A\n",
			Output: Individual{
				ID:                "ID",
				AutomatedRecordID: "A",
			},
		},
		{ // 186
			Input: "0 @ID@ INDI\n1 RIN\n",
			Err:   ErrContext{"Individual", cRIN, ErrInvalidLength{"AutomatedRecordID", "", 1, 12}},
		},
		{ // 187
			Input: "0 @ID@ INDI\n1 RIN A\n1 RIN B",
			Err:   ErrContext{"Individual", cRIN, ErrSingleMultiple},
		},
		{ // 188
			Input:   "0 @ID@ INDI\n1 RIN A\n1 RIN B",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID:                "ID",
				AutomatedRecordID: "A",
			},
		},
		{ // 189
			Input: "0 @ID@ INDI\n1 CHAN\n2 DATE 2006-05-04",
			Output: Individual{
				ID: "ID",
				ChangeDate: ChangeDateStructure{
					Date: ChangeDateTime{
						ChangeDate: "2006-05-04",
					},
				},
			},
		},
		{ // 190
			Input: "0 @ID@ INDI\n1 CHAN\n",
			Err:   ErrContext{"Individual", cCHAN, ErrContext{"ChangeDateStructure", cDATE, ErrRequiredMissing}},
		},
		{ // 191
			Input: "0 @ID@ INDI\n1 CHAN\n2 DATE 2006-05-04\n1 CHAN\n2 DATE 2000-01-01",
			Err:   ErrContext{"Individual", cCHAN, ErrSingleMultiple},
		},
		{ // 192
			Input:   "0 @ID@ INDI\n1 CHAN\n2 DATE 2006-05-04\n1 CHAN\n2 DATE 2000-01-01",
			Options: []Option{AllowMoreThanAllowed},
			Output: Individual{
				ID: "ID",
				ChangeDate: ChangeDateStructure{
					Date: ChangeDateTime{
						ChangeDate: "2006-05-04",
					},
				},
			},
		},
		{ // 193
			Input: "0 @ID@ INDI\n1 UNKNOWN\n",
			Err:   ErrContext{"Individual", "UNKNOWN", ErrUnknownTag},
		},
		{ // 194
			Input: "0 @ID@ INDI\n1 _UNKNOWN\n",
			Output: Individual{
				ID: "ID",
			},
		},
		{ // 195
			Input:   "0 @ID@ INDI\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: Individual{
				ID: "ID",
			},
		},
	})
}

func TestVerifiedIndividualFamEventDetail(t *testing.T) {
	test(t, []testData[VerifiedIndividualFamEventDetail]{
		{ // 1
			Input: "0 BIRT\n",
			Output: VerifiedIndividualFamEventDetail{
				VerifiedEventDetail: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 2
			Input: "0 BIRT\n1 FAMC @A@\n",
			Output: VerifiedIndividualFamEventDetail{
				Famc: "A",
				VerifiedEventDetail: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 3
			Input: "0 BIRT\n1 FAMC\n",
			Err:   ErrContext{"VerifiedIndividualFamEventDetail", cFAMC, ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 4
			Input: "0 BIRT\n1 FAMC @A@\n1 FAMC @B@",
			Err:   ErrContext{"VerifiedIndividualFamEventDetail", cFAMC, ErrSingleMultiple},
		},
		{ // 5
			Input:   "0 BIRT\n1 FAMC @A@\n1 FAMC @B@",
			Options: []Option{AllowMoreThanAllowed},
			Output: VerifiedIndividualFamEventDetail{
				Famc: "A",
				VerifiedEventDetail: VerifiedEventDetail{
					EventDetail: EventDetail{
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
		{ // 6
			Input: "0 BIRT\n1 TYPE A",
			Output: VerifiedIndividualFamEventDetail{
				VerifiedEventDetail: VerifiedEventDetail{
					EventDetail: EventDetail{
						Type:        "A",
						PhoneNumber: make([]PhoneNumber, 0, 3),
					},
				},
			},
		},
	})
}

func TestVerifiedEventDetail(t *testing.T) {
	test(t, []testData[VerifiedEventDetail]{
		{ // 1
			Input: "0 DEAT\n",
			Output: VerifiedEventDetail{
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
		{ // 2
			Input: "0 DEAT Y\n",
			Output: VerifiedEventDetail{
				Verified: "Y",
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
	})
}

func TestAdoptionEvent(t *testing.T) {
	test(t, []testData[AdoptionEvent]{
		{ // 1
			Input: "0 ADOP\n",
			Output: AdoptionEvent{
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
		{ // 2
			Input: "0 ADOP\n1 @A@ FAMC\n",
			Output: AdoptionEvent{
				Family: AdoptionReference{
					ID: "A",
				},
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
		{ // 3
			Input: "0 ADOP\n1 FAMC\n",
			Err:   ErrContext{"AdoptionEvent", cFAMC, ErrContext{"AdoptionReference", "xrefID", ErrInvalidLength{"Xref", "", 1, 22}}},
		},
		{ // 4
			Input: "0 ADOP\n1 @A@ FAMC\n1 @B@ FAMC\n",
			Err:   ErrContext{"AdoptionEvent", cFAMC, ErrSingleMultiple},
		},
		{ // 5
			Input:   "0 ADOP\n1 @A@ FAMC\n1 @B@ FAMC\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: AdoptionEvent{
				Family: AdoptionReference{
					ID: "A",
				},
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
		{ // 6
			Input: "0 ADOP\n1 TYPE A",
			Output: AdoptionEvent{
				EventDetail: EventDetail{
					Type:        "A",
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
	})
}

func TestAdoptionReference(t *testing.T) {
	test(t, []testData[AdoptionReference]{
		{ // 1
			Input: "0 @A@ FAMC\n",
			Output: AdoptionReference{
				ID: "A",
			},
		},
		{ // 2
			Input: "0 @A@ FAMC\n1 ADOP HUSB\n",
			Output: AdoptionReference{
				ID:        "A",
				AdoptedBy: cHUSB,
			},
		},
		{ // 3
			Input: "0 @A@ FAMC\n1 ADOP\n",
			Err:   ErrContext{"AdoptionReference", cADOP, ErrInvalidValue{"AdoptedBy", ""}},
		},
		{ // 4
			Input: "0 @A@ INDI\n1 ADOP WIFE\n1 ADOP BOTH",
			Err:   ErrContext{"AdoptionReference", cADOP, ErrSingleMultiple},
		},
		{ // 5
			Input:   "0 @A@ INDI\n1 ADOP WIFE\n1 ADOP BOTH",
			Options: []Option{AllowMoreThanAllowed},
			Output: AdoptionReference{
				ID:        "A",
				AdoptedBy: cWIFE,
			},
		},
		{ // 6
			Input: "0 @A@ FAMC\n1 UNKNOWN\n",
			Err:   ErrContext{"AdoptionReference", "UNKNOWN", ErrUnknownTag},
		},
		{ // 7
			Input: "0 @A@ FAMC\n1 _UNKNOWN\n",
			Output: AdoptionReference{
				ID: "A",
			},
		},
		{ // 8
			Input:   "0 @A@ FAMC\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: AdoptionReference{
				ID: "A",
			},
		},
	})
}

func TestCasteEvent(t *testing.T) {
	test(t, []testData[CasteEvent]{
		{ // 1
			Input: "0 CAST A\n",
			Output: CasteEvent{
				CasteName: "A",
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
		{ // 2
			Input: "0 CAST\n",
			Err:   ErrContext{"CasteEvent", "line_value", ErrInvalidLength{"CasteName", "", 1, 90}},
		},
		{ // 3
			Input: "0 CAST A\n1 TYPE B",
			Output: CasteEvent{
				CasteName: "A",
				EventDetail: EventDetail{
					Type:        "B",
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
	})
}

func TestDescriptionEvent(t *testing.T) {
	test(t, []testData[DescriptionEvent]{
		{ // 1
			Input: "0 DSCR A\n",
			Output: DescriptionEvent{
				PhysicalDescription: "A",
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
		{ // 2
			Input: "0 DSCR\n",
			Err:   ErrContext{"DescriptionEvent", "line_value", ErrInvalidLength{"PhysicalDescription", "", 1, 248}},
		},
		{ // 3
			Input: "0 DSCR A\n1 TYPE B",
			Output: DescriptionEvent{
				PhysicalDescription: "A",
				EventDetail: EventDetail{
					Type:        "B",
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
	})
}

func TestScholasticEvent(t *testing.T) {
	test(t, []testData[ScholasticEvent]{
		{ // 1
			Input: "0 EDUC A\n",
			Output: ScholasticEvent{
				ScholasticAchievement: "A",
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
		{ // 2
			Input: "0 EDUC\n",
			Err:   ErrContext{"ScholasticEvent", "line_value", ErrInvalidLength{"ScholasticAchievement", "", 1, 248}},
		},
		{ // 3
			Input: "0 EDUC A\n1 TYPE B",
			Output: ScholasticEvent{
				ScholasticAchievement: "A",
				EventDetail: EventDetail{
					Type:        "B",
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
	})
}

func TestNationalIDEvent(t *testing.T) {
	test(t, []testData[NationalIDEvent]{
		{ // 1
			Input: "0 IDNO A\n",
			Output: NationalIDEvent{
				NationalIDNumber: "A",
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
		{ // 2
			Input: "0 IDNO\n",
			Err:   ErrContext{"NationalIDEvent", "line_value", ErrInvalidLength{"NationalIDNumber", "", 1, 30}},
		},
		{ // 3
			Input: "0 IDNO A\n1 TYPE B",
			Output: NationalIDEvent{
				NationalIDNumber: "A",
				EventDetail: EventDetail{
					Type:        "B",
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
	})
}

func TestNationalOriginEvent(t *testing.T) {
	test(t, []testData[NationalOriginEvent]{
		{ // 1
			Input: "0 NATI A\n",
			Output: NationalOriginEvent{
				NationalOrTribalOrigin: "A",
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
		{ // 2
			Input: "0 NATI\n",
			Err:   ErrContext{"NationalOriginEvent", "line_value", ErrInvalidLength{"NationalOrTribalOrigin", "", 1, 120}},
		},
		{ // 3
			Input: "0 NATI A\n1 TYPE B",
			Output: NationalOriginEvent{
				NationalOrTribalOrigin: "A",
				EventDetail: EventDetail{
					Type:        "B",
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
	})
}

func TestChildrenEvent(t *testing.T) {
	test(t, []testData[ChildrenEvent]{
		{ // 1
			Input: "0 NCHI 1\n",
			Output: ChildrenEvent{
				CountOfChildren: 1,
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
		{ // 2
			Input: "0 NCHI\n",
			Err:   ErrContext{"ChildrenEvent", "line_value", ErrInvalidLength{"CountOfChildren", "", 1, 3}},
		},
		{ // 3
			Input: "0 NCHI 1\n1 TYPE B",
			Output: ChildrenEvent{
				CountOfChildren: 1,
				EventDetail: EventDetail{
					Type:        "B",
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
	})
}

func TestMarriagesEvent(t *testing.T) {
	test(t, []testData[MarriagesEvent]{
		{ // 1
			Input: "0 NMR 1\n",
			Output: MarriagesEvent{
				CountOfMarriages: 1,
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
		{ // 2
			Input: "0 NMR\n",
			Err:   ErrContext{"MarriagesEvent", "line_value", ErrInvalidLength{"CountOfMarriages", "", 1, 3}},
		},
		{ // 3
			Input: "0 NMR 1\n1 TYPE B",
			Output: MarriagesEvent{
				CountOfMarriages: 1,
				EventDetail: EventDetail{
					Type:        "B",
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
	})
}

func TestOccupationEvent(t *testing.T) {
	test(t, []testData[OccupationEvent]{
		{ // 1
			Input: "0 NMR A\n",
			Output: OccupationEvent{
				Occupation: "A",
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
		{ // 2
			Input: "0 NMR\n",
			Err:   ErrContext{"OccupationEvent", "line_value", ErrInvalidLength{"Occupation", "", 1, 90}},
		},
		{ // 3
			Input: "0 NMR A\n1 TYPE B",
			Output: OccupationEvent{
				Occupation: "A",
				EventDetail: EventDetail{
					Type:        "B",
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
	})
}

func TestPossessionEvent(t *testing.T) {
	test(t, []testData[PossessionEvent]{
		{ // 1
			Input: "0 PROP A\n",
			Output: PossessionEvent{
				Possessions: "A",
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
		{ // 2
			Input: "0 PROP\n",
			Err:   ErrContext{"PossessionEvent", "line_value", ErrInvalidLength{"Possessions", "", 1, 248}},
		},
		{ // 3
			Input: "0 PROP A\n1 TYPE B",
			Output: PossessionEvent{
				Possessions: "A",
				EventDetail: EventDetail{
					Type:        "B",
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
	})
}

func TestReligiousEvent(t *testing.T) {
	test(t, []testData[ReligiousEvent]{
		{ // 1
			Input: "0 RELI A\n",
			Output: ReligiousEvent{
				ReligiousAffiliation: "A",
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
		{ // 2
			Input: "0 RELI\n",
			Err:   ErrContext{"ReligiousEvent", "line_value", ErrInvalidLength{"ReligiousAffiliation", "", 1, 90}},
		},
		{ // 3
			Input: "0 RELI A\n1 TYPE B",
			Output: ReligiousEvent{
				ReligiousAffiliation: "A",
				EventDetail: EventDetail{
					Type:        "B",
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
	})
}

func TestSSNEvent(t *testing.T) {
	test(t, []testData[SSNEvent]{
		{ // 1
			Input: "0 SSN AAAAAAAAA\n",
			Output: SSNEvent{
				SocialSecurityNumber: "AAAAAAAAA",
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
		{ // 2
			Input: "0 SSN\n",
			Err:   ErrContext{"SSNEvent", "line_value", ErrInvalidLength{"SocialSecurityNumber", "", 9, 11}},
		},
		{ // 3
			Input: "0 SSN AAAAAAAAA\n1 TYPE B",
			Output: SSNEvent{
				SocialSecurityNumber: "AAAAAAAAA",
				EventDetail: EventDetail{
					Type:        "B",
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
	})
}

func TestNobilityEvent(t *testing.T) {
	test(t, []testData[NobilityEvent]{
		{ // 1
			Input: "0 TITL A\n",
			Output: NobilityEvent{
				NobilityTypeTitle: "A",
				EventDetail: EventDetail{
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
		{ // 2
			Input: "0 RELI\n",
			Err:   ErrContext{"NobilityEvent", "line_value", ErrInvalidLength{"NobilityTypeTitle", "", 1, 120}},
		},
		{ // 3
			Input: "0 TITL A\n1 TYPE B",
			Output: NobilityEvent{
				NobilityTypeTitle: "A",
				EventDetail: EventDetail{
					Type:        "B",
					PhoneNumber: make([]PhoneNumber, 0, 3),
				},
			},
		},
	})
}

func TestUserReferenceStructure(t *testing.T) {
	test(t, []testData[UserReferenceStructure]{
		{ // 1
			Input: "0 REFN A\n",
			Output: UserReferenceStructure{
				UserReferenceNumber: "A",
			},
		},
		{ // 2
			Input: "0 REFN A\n1 TYPE B\n",
			Output: UserReferenceStructure{
				UserReferenceNumber: "A",
				Type:                "B",
			},
		},
		{ // 3
			Input: "0 REFN A\n1 TYPE\n",
			Err:   ErrContext{"UserReferenceStructure", cTYPE, ErrInvalidLength{"UserReferenceType", "", 1, 40}},
		},
		{ // 4
			Input: "0 REFN A\n1 TYPE B\n1 TYPE C",
			Err:   ErrContext{"UserReferenceStructure", cTYPE, ErrSingleMultiple},
		},
		{ // 5
			Input:   "0 REFN A\n1 TYPE B\n1 TYPE C",
			Options: []Option{AllowMoreThanAllowed},
			Output: UserReferenceStructure{
				UserReferenceNumber: "A",
				Type:                "B",
			},
		},
		{ // 6
			Input: "0 REFN A\n1 UNKNOWN\n",
			Err:   ErrContext{"UserReferenceStructure", "UNKNOWN", ErrUnknownTag},
		},
		{ // 7
			Input: "0 REFN A\n1 _UNKNOWN\n",
			Output: UserReferenceStructure{
				UserReferenceNumber: "A",
			},
		},
		{ // 8
			Input:   "0 REFN A\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: UserReferenceStructure{
				UserReferenceNumber: "A",
			},
		},
	})
}

func TestMultimediaRecord(t *testing.T) {
	test(t, []testData[MultimediaRecord]{
		{ // 1
			Input: "0 OBJE\n",
			Err:   ErrContext{"MultimediaRecord", "xrefID", ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 2
			Input: "0 @A@ OBJE\n",
			Err:   ErrContext{"MultimediaRecord", cBLOB, ErrRequiredMissing},
		},
		{ // 3
			Input: "0 @A@ OBJE\n1 BLOB\n",
			Err:   ErrContext{"MultimediaRecord", cFORM, ErrRequiredMissing},
		},
		{ // 4
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP",
			Output: MultimediaRecord{
				ID:     "A",
				Format: cbmp,
			},
		},
		{ // 5
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM\n",
			Err:   ErrContext{"MultimediaRecord", cFORM, ErrInvalidValue{"MultimediaFormat", ""}},
		},
		{ // 6
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 FORM GIF",
			Err:   ErrContext{"MultimediaRecord", cFORM, ErrSingleMultiple},
		},
		{ // 7
			Input:   "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 FORM GIF",
			Options: []Option{AllowMoreThanAllowed},
			Output: MultimediaRecord{
				ID:     "A",
				Format: cbmp,
			},
		},
		{ // 8
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 TITL B",
			Output: MultimediaRecord{
				ID:     "A",
				Format: cbmp,
				Title:  "B",
			},
		},
		{ // 9
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 TITL\n",
			Err:   ErrContext{"MultimediaRecord", cTITL, ErrInvalidLength{"DescriptiveTitle", "", 1, 248}},
		},
		{ // 10
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 TITL B\n1 TITL C",
			Err:   ErrContext{"MultimediaRecord", cTITL, ErrSingleMultiple},
		},
		{ // 11
			Input:   "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 TITL B\n1 TITL C",
			Options: []Option{AllowMoreThanAllowed},
			Output: MultimediaRecord{
				ID:     "A",
				Format: cbmp,
				Title:  "B",
			},
		},
		{ // 12
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 @B@ NOTE\n",
			Output: MultimediaRecord{
				ID:     "A",
				Format: cbmp,
				Notes: []NoteStructure{
					{
						Data: &NoteID{
							ID: "B",
						},
					},
				},
			},
		},
		{ // 13
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 NOTE\n",
			Err:   ErrContext{"MultimediaRecord", cNOTE, ErrContext{"NoteText", "line_value", ErrInvalidLength{"SubmitterText", "", 1, 248}}},
		},
		{ // 14
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 @C@ NOTE\n1 @D@ NOTE\n",
			Output: MultimediaRecord{
				ID:     "A",
				Format: cbmp,
				Notes: []NoteStructure{
					{
						Data: &NoteID{
							ID: "C",
						},
					},
					{
						Data: &NoteID{
							ID: "D",
						},
					},
				},
			},
		},
		{ // 15
			Input: "0 @A@ OBJE\n1 BLOB " + strings.Repeat("A", 88) + "\n1 FORM BMP\n",
			Err:   ErrContext{"MultimediaRecord", cBLOB, ErrInvalidLength{"EncodedMultimediaLine", strings.Repeat("A", 88), 0, 87}},
		},
		{ // 16
			Input: "0 @A@ OBJE\n1 BLOB A\n1 FORM BMP\n1 BLOB B\n",
			Err:   ErrContext{"MultimediaRecord", cBLOB, ErrSingleMultiple},
		},
		{ // 17
			Input:   "0 @A@ OBJE\n1 BLOB B\n1 FORM BMP\n1 BLOB C\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: MultimediaRecord{
				ID:     "A",
				Blob:   "B",
				Format: cbmp,
			},
		},
		{ // 18
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 OBJE @B@",
			Output: MultimediaRecord{
				ID:              "A",
				Format:          cbmp,
				ContinuedObject: "B",
			},
		},
		{ // 19
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 OBJE\n",
			Err:   ErrContext{"MultimediaRecord", cOBJE, ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 20
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 OBJE @B@\n1 OBJE @C@",
			Err:   ErrContext{"MultimediaRecord", cOBJE, ErrSingleMultiple},
		},
		{ // 21
			Input:   "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 OBJE @B@\n1 OBJE @C@",
			Options: []Option{AllowMoreThanAllowed},
			Output: MultimediaRecord{
				ID:              "A",
				Format:          cbmp,
				ContinuedObject: "B",
			},
		},
		{ // 22
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 REFN @B@",
			Output: MultimediaRecord{
				ID:     "A",
				Format: cbmp,
				UserReferences: []UserReferenceStructure{
					{
						UserReferenceNumber: "B",
					},
				},
			},
		},
		{ // 23
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 REFN\n",
			Err:   ErrContext{"MultimediaRecord", cREFN, ErrContext{"UserReferenceStructure", "line_value", ErrInvalidLength{"UserReferenceNumber", "", 1, 20}}},
		},
		{ // 24
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 REFN @B@\n1 REFN @C@",
			Output: MultimediaRecord{
				ID:     "A",
				Format: cbmp,
				UserReferences: []UserReferenceStructure{
					{
						UserReferenceNumber: "B",
					},
					{
						UserReferenceNumber: "C",
					},
				},
			},
		},
		{ // 25
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 RIN B",
			Output: MultimediaRecord{
				ID:                "A",
				Format:            cbmp,
				AutomatedRecordID: "B",
			},
		},
		{ // 26
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 RIN\n",
			Err:   ErrContext{"MultimediaRecord", cRIN, ErrInvalidLength{"AutomatedRecordID", "", 1, 12}},
		},
		{ // 27
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 RIN B\n1 RIN C",
			Err:   ErrContext{"MultimediaRecord", cRIN, ErrSingleMultiple},
		},
		{ // 28
			Input:   "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 RIN B\n1 RIN C",
			Options: []Option{AllowMoreThanAllowed},
			Output: MultimediaRecord{
				ID:                "A",
				Format:            cbmp,
				AutomatedRecordID: "B",
			},
		},
		{ // 29
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 CHAN\n2 DATE 2006-05-04",
			Output: MultimediaRecord{
				ID:     "A",
				Format: cbmp,
				ChangeDate: ChangeDateStructure{
					Date: ChangeDateTime{
						ChangeDate: "2006-05-04",
					},
				},
			},
		},
		{ // 30
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 CHAN\n",
			Err:   ErrContext{"MultimediaRecord", cCHAN, ErrContext{"ChangeDateStructure", cDATE, ErrRequiredMissing}},
		},
		{ // 31
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 CHAN\n2 DATE 2006-05-04\n1 CHAN\n2 DATE 2000-01-01",
			Err:   ErrContext{"MultimediaRecord", cCHAN, ErrSingleMultiple},
		},
		{ // 32
			Input:   "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 CHAN\n2 DATE 2006-05-04\n1 CHAN\n2 DATE 2000-01-01",
			Options: []Option{AllowMoreThanAllowed},
			Output: MultimediaRecord{
				ID:     "A",
				Format: cbmp,
				ChangeDate: ChangeDateStructure{
					Date: ChangeDateTime{
						ChangeDate: "2006-05-04",
					},
				},
			},
		},
		{ // 33
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 UNKNOWN\n",
			Err:   ErrContext{"MultimediaRecord", "UNKNOWN", ErrUnknownTag},
		},
		{ // 34
			Input: "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 _UNKNOWN\n",
			Output: MultimediaRecord{
				ID:     "A",
				Format: cbmp,
			},
		},
		{ // 35
			Input:   "0 @A@ OBJE\n1 BLOB\n1 FORM BMP\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: MultimediaRecord{
				ID:     "A",
				Format: cbmp,
			},
		},
	})
}

func TestNoteRecord(t *testing.T) {
	test(t, []testData[NoteRecord]{
		{ // 1
			Input: "0 NOTE\n",
			Err:   ErrContext{"NoteRecord", "xrefID", ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 2
			Input: "0 @A@ NOTE\n",
			Err:   ErrContext{"NoteRecord", "line_value", ErrInvalidLength{"SubmitterText", "", 1, 248}},
		},
		{ // 3
			Input: "0 @A@ NOTE B\n",
			Output: NoteRecord{
				ID:            "A",
				SubmitterText: "B",
			},
		},
		{ // 4
			Input: "0 @A@ NOTE B\n1 @C@ SOUR\n",
			Output: NoteRecord{
				ID:            "A",
				SubmitterText: "B",
				Sources: []SourceCitation{
					{
						Data: &SourceID{
							ID: "C",
						},
					},
				},
			},
		},
		{ // 5
			Input: "0 @A@ NOTE B\n1 SOUR\n",
			Err:   ErrContext{"NoteRecord", cSOUR, ErrContext{"SourceText", "line_value", ErrInvalidLength{"SourceDescription", "", 1, 248}}},
		},
		{ // 6
			Input: "0 @A@ NOTE B\n1 @C@ SOUR\n1 @D@ SOUR\n",
			Output: NoteRecord{
				ID:            "A",
				SubmitterText: "B",
				Sources: []SourceCitation{
					{
						Data: &SourceID{
							ID: "C",
						},
					},
					{
						Data: &SourceID{
							ID: "D",
						},
					},
				},
			},
		},
		{ // 7
			Input: "0 @A@ NOTE B\n1 REFN @C@\n",
			Output: NoteRecord{
				ID:            "A",
				SubmitterText: "B",
				UserReferences: []UserReferenceStructure{
					{
						UserReferenceNumber: "C",
					},
				},
			},
		},
		{ // 8
			Input: "0 @A@ NOTE B\n1 REFN\n",
			Err:   ErrContext{"NoteRecord", cREFN, ErrContext{"UserReferenceStructure", "line_value", ErrInvalidLength{"UserReferenceNumber", "", 1, 20}}},
		},
		{ // 9
			Input: "0 @A@ NOTE B\n1 REFN @C@\n1 REFN @D@",
			Output: NoteRecord{
				ID:            "A",
				SubmitterText: "B",
				UserReferences: []UserReferenceStructure{
					{
						UserReferenceNumber: "C",
					},
					{
						UserReferenceNumber: "D",
					},
				},
			},
		},
		{ // 10
			Input: "0 @A@ NOTE B\n1 RIN B",
			Output: NoteRecord{
				ID:                "A",
				SubmitterText:     "B",
				AutomatedRecordID: "B",
			},
		},
		{ // 11
			Input: "0 @A@ NOTE B\n1 RIN\n",
			Err:   ErrContext{"NoteRecord", cRIN, ErrInvalidLength{"AutomatedRecordID", "", 1, 12}},
		},
		{ // 12
			Input: "0 @A@ NOTE B\n1 RIN B\n1 RIN C",
			Err:   ErrContext{"NoteRecord", cRIN, ErrSingleMultiple},
		},
		{ // 13
			Input:   "0 @A@ NOTE B\n1 RIN B\n1 RIN C",
			Options: []Option{AllowMoreThanAllowed},
			Output: NoteRecord{
				ID:                "A",
				SubmitterText:     "B",
				AutomatedRecordID: "B",
			},
		},
		{ // 14
			Input: "0 @A@ NOTE B\n1 CHAN\n2 DATE 2006-05-04",
			Output: NoteRecord{
				ID:            "A",
				SubmitterText: "B",
				ChangeDate: ChangeDateStructure{
					Date: ChangeDateTime{
						ChangeDate: "2006-05-04",
					},
				},
			},
		},
		{ // 15
			Input: "0 @A@ NOTE B\n1 CHAN\n",
			Err:   ErrContext{"NoteRecord", cCHAN, ErrContext{"ChangeDateStructure", cDATE, ErrRequiredMissing}},
		},
		{ // 16
			Input: "0 @A@ NOTE B\n1 CHAN\n2 DATE 2006-05-04\n1 CHAN\n2 DATE 2000-01-01",
			Err:   ErrContext{"NoteRecord", cCHAN, ErrSingleMultiple},
		},
		{ // 17
			Input:   "0 @A@ NOTE B\n1 CHAN\n2 DATE 2006-05-04\n1 CHAN\n2 DATE 2000-01-01",
			Options: []Option{AllowMoreThanAllowed},
			Output: NoteRecord{
				ID:            "A",
				SubmitterText: "B",
				ChangeDate: ChangeDateStructure{
					Date: ChangeDateTime{
						ChangeDate: "2006-05-04",
					},
				},
			},
		},
		{ // 18
			Input: "0 @A@ NOTE B\n1 UNKNOWN\n",
			Err:   ErrContext{"NoteRecord", "UNKNOWN", ErrUnknownTag},
		},
		{ // 19
			Input: "0 @A@ NOTE B\n1 _UNKNOWN\n",
			Output: NoteRecord{
				ID:            "A",
				SubmitterText: "B",
			},
		},
		{ // 20
			Input:   "0 @A@ NOTE B\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: NoteRecord{
				ID:            "A",
				SubmitterText: "B",
			},
		},
	})
}

func TestRepositoryRecord(t *testing.T) {
	test(t, []testData[RepositoryRecord]{
		{ // 1
			Input: "0 REPO\n",
			Err:   ErrContext{"RepositoryRecord", "xrefID", ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 2
			Input: "0 @A@ REPO\n",
			Output: RepositoryRecord{
				ID:          "A",
				PhoneNumber: make([]PhoneNumber, 0, 3),
			},
		},
		{ // 3
			Input: "0 @A@ REPO\n1 NAME B",
			Output: RepositoryRecord{
				ID:               "A",
				PhoneNumber:      make([]PhoneNumber, 0, 3),
				NameOfRepository: "B",
			},
		},
		{ // 4
			Input: "0 @A@ REPO\n1 NAME\n",
			Err:   ErrContext{"RepositoryRecord", cNAME, ErrInvalidLength{"NameOfRepository", "", 1, 90}},
		},
		{ // 5
			Input: "0 @A@ REPO\n1 NAME B\n1 NAME C",
			Err:   ErrContext{"RepositoryRecord", cNAME, ErrSingleMultiple},
		},
		{ // 6
			Input:   "0 @A@ REPO\n1 NAME B\n1 NAME C",
			Options: []Option{AllowMoreThanAllowed},
			Output: RepositoryRecord{
				ID:               "A",
				PhoneNumber:      make([]PhoneNumber, 0, 3),
				NameOfRepository: "B",
			},
		},
		{ // 7
			Input: "0 @A@ REPO\n1 ADDR line 1",
			Output: RepositoryRecord{
				ID:          "A",
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Address: AddressStructure{
					AddressLine: "line 1",
				},
			},
		},
		{ // 8
			Input: "0 @A@ REPO\n1 ADDR\n",
			Err:   ErrContext{"RepositoryRecord", cADDR, ErrContext{"AddressStructure", "line_value", ErrInvalidLength{"AddressLine", "", 1, 60}}},
		},
		{ // 9
			Input: "0 @A@ REPO\n1 ADDR line_1\n1 ADDR line_2",
			Err:   ErrContext{"RepositoryRecord", cADDR, ErrSingleMultiple},
		},
		{ // 10
			Input:   "0 @A@ REPO\n1 ADDR line_1\n1 ADDR line_2",
			Options: []Option{AllowMoreThanAllowed},
			Output: RepositoryRecord{
				ID:          "A",
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Address: AddressStructure{
					AddressLine: "line_1",
				},
			},
		},
		{ // 11
			Input: "0 @A@ REPO\n1 PHON number",
			Output: RepositoryRecord{
				ID:          "A",
				PhoneNumber: []PhoneNumber{"number"},
			},
		},
		{ // 12
			Input: "0 @A@ REPO\n1 PHON\n",
			Err:   ErrContext{"RepositoryRecord", cPHON, ErrInvalidLength{"PhoneNumber", "", 1, 25}},
		},
		{ // 13
			Input: "0 @A@ REPO\n1 PHON number 1\n1 PHON number 2",
			Output: RepositoryRecord{
				ID:          "A",
				PhoneNumber: []PhoneNumber{"number 1", "number 2"},
			},
		},
		{ // 14
			Input: "0 @A@ REPO\n1 PHON number 1\n1 PHON number 2\n1 PHON number 3",
			Output: RepositoryRecord{
				ID:          "A",
				PhoneNumber: []PhoneNumber{"number 1", "number 2", "number 3"},
			},
		},
		{ // 15
			Input: "0 @A@ REPO\n1 PHON number 1\n1 PHON number 2\n1 PHON number 3\n1 PHON number 4",
			Err:   ErrContext{"RepositoryRecord", cPHON, ErrTooMany(3)},
		},
		{ // 16
			Input:   "0 @A@ REPO\n1 PHON number 1\n1 PHON number 2\n1 PHON number 3\n1 PHON number 4",
			Options: []Option{AllowMoreThanAllowed},
			Output: RepositoryRecord{
				ID:          "A",
				PhoneNumber: []PhoneNumber{"number 1", "number 2", "number 3"},
			},
		},
		{ // 17
			Input: "0 @A@ REPO\n1 @A@ NOTE\n",
			Output: RepositoryRecord{
				ID:          "A",
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Notes: []NoteStructure{
					{
						Data: &NoteID{
							ID: "A",
						},
					},
				},
			},
		},
		{ // 18
			Input: "0 @A@ REPO\n1 NOTE\n",
			Err:   ErrContext{"RepositoryRecord", cNOTE, ErrContext{"NoteText", "line_value", ErrInvalidLength{"SubmitterText", "", 1, 248}}},
		},
		{ // 19
			Input: "0 @A@ REPO\n1 @A@ NOTE\n1 @B@ NOTE\n",
			Output: RepositoryRecord{
				ID:          "A",
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Notes: []NoteStructure{
					{
						Data: &NoteID{
							ID: "A",
						},
					},
					{
						Data: &NoteID{
							ID: "B",
						},
					},
				},
			},
		},
		{ // 20
			Input: "0 @A@ REPO\n1 REFN @C@\n",
			Output: RepositoryRecord{
				ID:          "A",
				PhoneNumber: make([]PhoneNumber, 0, 3),
				UserReferences: []UserReferenceStructure{
					{
						UserReferenceNumber: "C",
					},
				},
			},
		},
		{ // 21
			Input: "0 @A@ REPO\n1 REFN\n",
			Err:   ErrContext{"RepositoryRecord", cREFN, ErrContext{"UserReferenceStructure", "line_value", ErrInvalidLength{"UserReferenceNumber", "", 1, 20}}},
		},
		{ // 22
			Input: "0 @A@ REPO\n1 REFN @C@\n1 REFN @D@",
			Output: RepositoryRecord{
				ID:          "A",
				PhoneNumber: make([]PhoneNumber, 0, 3),
				UserReferences: []UserReferenceStructure{
					{
						UserReferenceNumber: "C",
					},
					{
						UserReferenceNumber: "D",
					},
				},
			},
		},
		{ // 23
			Input: "0 @A@ REPO\n1 RIN B",
			Output: RepositoryRecord{
				ID:                "A",
				PhoneNumber:       make([]PhoneNumber, 0, 3),
				AutomatedRecordID: "B",
			},
		},
		{ // 24
			Input: "0 @A@ REPO\n1 RIN\n",
			Err:   ErrContext{"RepositoryRecord", cRIN, ErrInvalidLength{"AutomatedRecordID", "", 1, 12}},
		},
		{ // 25
			Input: "0 @A@ REPO\n1 RIN B\n1 RIN C",
			Err:   ErrContext{"RepositoryRecord", cRIN, ErrSingleMultiple},
		},
		{ // 26
			Input:   "0 @A@ REPO\n1 RIN B\n1 RIN C",
			Options: []Option{AllowMoreThanAllowed},
			Output: RepositoryRecord{
				ID:                "A",
				PhoneNumber:       make([]PhoneNumber, 0, 3),
				AutomatedRecordID: "B",
			},
		},
		{ // 27
			Input: "0 @A@ REPO\n1 CHAN\n2 DATE 2006-05-04",
			Output: RepositoryRecord{
				ID:          "A",
				PhoneNumber: make([]PhoneNumber, 0, 3),
				ChangeDate: ChangeDateStructure{
					Date: ChangeDateTime{
						ChangeDate: "2006-05-04",
					},
				},
			},
		},
		{ // 28
			Input: "0 @A@ REPO\n1 CHAN\n",
			Err:   ErrContext{"RepositoryRecord", cCHAN, ErrContext{"ChangeDateStructure", cDATE, ErrRequiredMissing}},
		},
		{ // 29
			Input: "0 @A@ REPO\n1 CHAN\n2 DATE 2006-05-04\n1 CHAN\n2 DATE 2000-01-01",
			Err:   ErrContext{"RepositoryRecord", cCHAN, ErrSingleMultiple},
		},
		{ // 30
			Input:   "0 @A@ REPO\n1 CHAN\n2 DATE 2006-05-04\n1 CHAN\n2 DATE 2000-01-01",
			Options: []Option{AllowMoreThanAllowed},
			Output: RepositoryRecord{
				ID:          "A",
				PhoneNumber: make([]PhoneNumber, 0, 3),
				ChangeDate: ChangeDateStructure{
					Date: ChangeDateTime{
						ChangeDate: "2006-05-04",
					},
				},
			},
		},
		{ // 31
			Input: "0 @A@ REPO\n1 UNKNOWN\n",
			Err:   ErrContext{"RepositoryRecord", "UNKNOWN", ErrUnknownTag},
		},
		{ // 32
			Input: "0 @A@ REPO\n1 _UNKNOWN\n",
			Output: RepositoryRecord{
				ID:          "A",
				PhoneNumber: make([]PhoneNumber, 0, 3),
			},
		},
		{ // 33
			Input:   "0 @A@ REPO\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: RepositoryRecord{
				ID:          "A",
				PhoneNumber: make([]PhoneNumber, 0, 3),
			},
		},
	})
}

func TestSourceRecord(t *testing.T) {
	test(t, []testData[SourceRecord]{
		{ // 1
			Input: "0 SOUR\n",
			Err:   ErrContext{"SourceRecord", "xrefID", ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 2
			Input: "0 @A@ SOUR\n",
			Err:   ErrContext{"SourceRecord", cREPO, ErrRequiredMissing},
		},
		{ // 3
			Input: "0 @A@ SOUR\n1 @B@ REPO\n",
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
			},
		},
		{ // 4
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 DATA\n2 AGNC C",
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				Data: SourceRecordDataStructure{
					ResponsibleAgency: "C",
				},
			},
		},
		{ // 5
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 DATA\n2 AGNC\n",
			Err:   ErrContext{"SourceRecord", cDATA, ErrContext{"SourceRecordDataStructure", cAGNC, ErrInvalidLength{"ResponsibleAgency", "", 1, 120}}},
		},
		{ // 6
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 DATA\n2 AGNC C\n1 DATA\n2 AGNC D",
			Err:   ErrContext{"SourceRecord", cDATA, ErrSingleMultiple},
		},
		{ // 7
			Input:   "0 @A@ SOUR\n1 @B@ REPO\n1 DATA\n2 AGNC C\n1 DATA\n2 AGNC D",
			Options: []Option{AllowMoreThanAllowed},
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				Data: SourceRecordDataStructure{
					ResponsibleAgency: "C",
				},
			},
		},
		{ // 8
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 AUTH C",
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				Originator: "C",
			},
		},
		{ // 9
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 AUTH\n",
			Err:   ErrContext{"SourceRecord", cAUTH, ErrInvalidLength{"SourceOriginator", "", 1, 248}},
		},
		{ // 10
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 AUTH C\n1 AUTH D",
			Err:   ErrContext{"SourceRecord", cAUTH, ErrSingleMultiple},
		},
		{ // 11
			Input:   "0 @A@ SOUR\n1 @B@ REPO\n1 AUTH C\n1 AUTH D",
			Options: []Option{AllowMoreThanAllowed},
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				Originator: "C",
			},
		},
		{ // 12
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 TITL C",
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				Title: "C",
			},
		},
		{ // 13
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 TITL\n",
			Err:   ErrContext{"SourceRecord", cTITL, ErrInvalidLength{"SourceDescriptiveTitle", "", 1, 248}},
		},
		{ // 14
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 TITL C\n1 TITL D",
			Err:   ErrContext{"SourceRecord", cTITL, ErrSingleMultiple},
		},
		{ // 15
			Input:   "0 @A@ SOUR\n1 @B@ REPO\n1 TITL C\n1 TITL D",
			Options: []Option{AllowMoreThanAllowed},
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				Title: "C",
			},
		},
		{ // 16
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 ABBR C",
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				FiledBy: "C",
			},
		},
		{ // 17
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 ABBR\n",
			Err:   ErrContext{"SourceRecord", cABBR, ErrInvalidLength{"SourceFiledByEntry", "", 1, 60}},
		},
		{ // 18
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 ABBR C\n1 ABBR D",
			Err:   ErrContext{"SourceRecord", cABBR, ErrSingleMultiple},
		},
		{ // 19
			Input:   "0 @A@ SOUR\n1 @B@ REPO\n1 ABBR C\n1 ABBR D",
			Options: []Option{AllowMoreThanAllowed},
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				FiledBy: "C",
			},
		},
		{ // 20
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 PUBL C",
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				PublicationFacts: "C",
			},
		},
		{ // 21
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 PUBL\n",
			Err:   ErrContext{"SourceRecord", cPUBL, ErrInvalidLength{"SourcePublicationFacts", "", 1, 248}},
		},
		{ // 22
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 PUBL C\n1 PUBL D",
			Err:   ErrContext{"SourceRecord", cPUBL, ErrSingleMultiple},
		},
		{ // 23
			Input:   "0 @A@ SOUR\n1 @B@ REPO\n1 PUBL C\n1 PUBL D",
			Options: []Option{AllowMoreThanAllowed},
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				PublicationFacts: "C",
			},
		},
		{ // 24
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 TEXT C",
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				TextFromSource: "C",
			},
		},
		{ // 25
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 TEXT\n",
			Err:   ErrContext{"SourceRecord", cTEXT, ErrInvalidLength{"TextFromSource", "", 1, 248}},
		},
		{ // 26
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 TEXT C\n1 TEXT D",
			Err:   ErrContext{"SourceRecord", cTEXT, ErrSingleMultiple},
		},
		{ // 27
			Input:   "0 @A@ SOUR\n1 @B@ REPO\n1 TEXT C\n1 TEXT D",
			Options: []Option{AllowMoreThanAllowed},
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				TextFromSource: "C",
			},
		},
		{ // 28
			Input: "0 @A@ SOUR\n1 @B@ REPO\n",
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
			},
		},
		{ // 29
			Input: "0 @A@ SOUR\n1 REPO\n",
			Err:   ErrContext{"SourceRecord", cREPO, ErrContext{"SourceRepositoryCitation", "xrefID", ErrInvalidLength{"Xref", "", 1, 22}}},
		},
		{ // 30
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 @C@ REPO\n",
			Err:   ErrContext{"SourceRecord", cREPO, ErrSingleMultiple},
		},
		{ // 31
			Input:   "0 @A@ SOUR\n1 @B@ REPO\n1 @C@ REPO\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
			},
		},
		{ // 32
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 OBJE @C@\n",
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				ContinuedObject: "C",
			},
		},
		{ // 33
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 OBJE\n",
			Err:   ErrContext{"SourceRecord", cOBJE, ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 34
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 OBJE @C@\n1 OBJE @D@\n",
			Err:   ErrContext{"SourceRecord", cOBJE, ErrSingleMultiple},
		},
		{ // 35
			Input:   "0 @A@ SOUR\n1 @B@ REPO\n1 OBJE @C@\n1 OBJE @D@\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				ContinuedObject: "C",
			},
		},
		{ // 36
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 @A@ NOTE\n",
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				Notes: []NoteStructure{
					{
						Data: &NoteID{
							ID: "A",
						},
					},
				},
			},
		},
		{ // 37
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 NOTE\n",
			Err:   ErrContext{"SourceRecord", cNOTE, ErrContext{"NoteText", "line_value", ErrInvalidLength{"SubmitterText", "", 1, 248}}},
		},
		{ // 38
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 @A@ NOTE\n1 @B@ NOTE\n",
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				Notes: []NoteStructure{
					{
						Data: &NoteID{
							ID: "A",
						},
					},
					{
						Data: &NoteID{
							ID: "B",
						},
					},
				},
			},
		},
		{ // 39
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 REFN @C@\n",
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				UserReferences: []UserReferenceStructure{
					{
						UserReferenceNumber: "C",
					},
				},
			},
		},
		{ // 40
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 REFN\n",
			Err:   ErrContext{"SourceRecord", cREFN, ErrContext{"UserReferenceStructure", "line_value", ErrInvalidLength{"UserReferenceNumber", "", 1, 20}}},
		},
		{ // 41
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 REFN @C@\n1 REFN @D@",
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				UserReferences: []UserReferenceStructure{
					{
						UserReferenceNumber: "C",
					},
					{
						UserReferenceNumber: "D",
					},
				},
			},
		},
		{ // 42
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 RIN B",
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				AutomatedRecordID: "B",
			},
		},
		{ // 43
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 RIN\n",
			Err:   ErrContext{"SourceRecord", cRIN, ErrInvalidLength{"AutomatedRecordID", "", 1, 12}},
		},
		{ // 44
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 RIN B\n1 RIN C",
			Err:   ErrContext{"SourceRecord", cRIN, ErrSingleMultiple},
		},
		{ // 45
			Input:   "0 @A@ SOUR\n1 @B@ REPO\n1 RIN B\n1 RIN C",
			Options: []Option{AllowMoreThanAllowed},
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				AutomatedRecordID: "B",
			},
		},
		{ // 46
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 CHAN\n2 DATE 2006-05-04",
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				ChangeDate: ChangeDateStructure{
					Date: ChangeDateTime{
						ChangeDate: "2006-05-04",
					},
				},
			},
		},
		{ // 47
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 CHAN\n",
			Err:   ErrContext{"SourceRecord", cCHAN, ErrContext{"ChangeDateStructure", cDATE, ErrRequiredMissing}},
		},
		{ // 48
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 CHAN\n2 DATE 2006-05-04\n1 CHAN\n2 DATE 2000-01-01",
			Err:   ErrContext{"SourceRecord", cCHAN, ErrSingleMultiple},
		},
		{ // 49
			Input:   "0 @A@ SOUR\n1 @B@ REPO\n1 CHAN\n2 DATE 2006-05-04\n1 CHAN\n2 DATE 2000-01-01",
			Options: []Option{AllowMoreThanAllowed},
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
				ChangeDate: ChangeDateStructure{
					Date: ChangeDateTime{
						ChangeDate: "2006-05-04",
					},
				},
			},
		},
		{ // 50
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 UNKNOWN\n",
			Err:   ErrContext{"SourceRecord", "UNKNOWN", ErrUnknownTag},
		},
		{ // 51
			Input: "0 @A@ SOUR\n1 @B@ REPO\n1 _UNKNOWN\n",
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
			},
		},
		{ // 52
			Input:   "0 @A@ SOUR\n1 @B@ REPO\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: SourceRecord{
				ID: "A",
				SourceRepositoryCitation: SourceRepositoryCitation{
					ID: "B",
				},
			},
		},
	})
}

func TestSourceRecordDataStructure(t *testing.T) {
	test(t, []testData[SourceRecordDataStructure]{
		{ // 1
			Input:  "0 DATA\n",
			Output: SourceRecordDataStructure{},
		},
		{ // 2
			Input: "0 DATA\n1 EVEN\n2 PLAC A",
			Output: SourceRecordDataStructure{
				EventsRecorded: []EventsRecordedStructure{
					{
						SourceJurisdictionPlace: "A",
					},
				},
			},
		},
		{ // 3
			Input: "0 DATA\n1 EVEN\n2 PLAC\n",
			Err:   ErrContext{"SourceRecordDataStructure", cEVEN, ErrContext{"EventsRecordedStructure", cPLAC, ErrInvalidLength{"SourceJurisdictionPlace", "", 1, 120}}},
		},
		{ // 4
			Input: "0 DATA\n1 EVEN\n2 PLAC A\n1 EVEN\n2 PLAC B",
			Output: SourceRecordDataStructure{
				EventsRecorded: []EventsRecordedStructure{
					{
						SourceJurisdictionPlace: "A",
					},
					{
						SourceJurisdictionPlace: "B",
					},
				},
			},
		},
		{ // 5
			Input: "0 DATA\n1 EVEN\n2 PLAC A\n1 AGNC B",
			Output: SourceRecordDataStructure{
				EventsRecorded: []EventsRecordedStructure{
					{
						SourceJurisdictionPlace: "A",
					},
				},
				ResponsibleAgency: "B",
			},
		},
		{ // 6
			Input: "0 DATA\n1 EVEN\n2 PLAC A\n1 AGNC\n",
			Err:   ErrContext{"SourceRecordDataStructure", cAGNC, ErrInvalidLength{"ResponsibleAgency", "", 1, 120}},
		},
		{ // 7
			Input: "0 DATA\n1 EVEN\n2 PLAC A\n1 AGNC B\n1 AGNC C",
			Err:   ErrContext{"SourceRecordDataStructure", cAGNC, ErrSingleMultiple},
		},
		{ // 8
			Input:   "0 DATA\n1 EVEN\n2 PLAC A\n1 AGNC B\n1 AGNC C",
			Options: []Option{AllowMoreThanAllowed},
			Output: SourceRecordDataStructure{
				EventsRecorded: []EventsRecordedStructure{
					{
						SourceJurisdictionPlace: "A",
					},
				},
				ResponsibleAgency: "B",
			},
		},
		{ // 9
			Input: "0 DATA\n1 EVEN\n2 PLAC A\n1 @A@ NOTE\n",
			Output: SourceRecordDataStructure{
				EventsRecorded: []EventsRecordedStructure{
					{
						SourceJurisdictionPlace: "A",
					},
				},
				Notes: []NoteStructure{
					{
						Data: &NoteID{
							ID: "A",
						},
					},
				},
			},
		},
		{ // 10
			Input: "0 DATA\n1 EVEN\n2 PLAC A\n1 NOTE\n",
			Err:   ErrContext{"SourceRecordDataStructure", cNOTE, ErrContext{"NoteText", "line_value", ErrInvalidLength{"SubmitterText", "", 1, 248}}},
		},
		{ // 11
			Input: "0 DATA\n1 EVEN\n2 PLAC A\n1 @A@ NOTE\n1 @B@ NOTE\n",
			Output: SourceRecordDataStructure{
				EventsRecorded: []EventsRecordedStructure{
					{
						SourceJurisdictionPlace: "A",
					},
				},
				Notes: []NoteStructure{
					{
						Data: &NoteID{
							ID: "A",
						},
					},
					{
						Data: &NoteID{
							ID: "B",
						},
					},
				},
			},
		},
		{ // 12
			Input: "0 DATA\n1 UNKNOWN\n",
			Err:   ErrContext{"SourceRecordDataStructure", "UNKNOWN", ErrUnknownTag},
		},
		{ // 13
			Input: "0 DATA\n1 _UNKNOWN\n",
		},
		{ // 14
			Input:   "0 DATA\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
		},
	})
}

func TestEventsRecordedStructure(t *testing.T) {
	test(t, []testData[EventsRecordedStructure]{
		{ // 1
			Input:  "0 EVEN\n",
			Output: EventsRecordedStructure{},
		},
		{ // 2
			Input: "0 EVEN\n1 DATE 2006-01-02",
			Output: EventsRecordedStructure{
				DatePeriod: "2006-01-02",
			},
		},
		{ // 3
			Input: "0 EVEN\n1 DATE\n",
			Err:   ErrContext{"EventsRecordedStructure", cDATE, ErrInvalidLength{"DatePeriod", "", 7, 35}},
		},
		{ // 4
			Input: "0 EVEN\n1 DATE 2006-01-02\n1 DATE 2012-12-25",
			Err:   ErrContext{"EventsRecordedStructure", cDATE, ErrSingleMultiple},
		},
		{ // 5
			Input:   "0 EVEN\n1 DATE 2006-01-02\n1 DATE 2012-12-25",
			Options: []Option{AllowMoreThanAllowed},
			Output: EventsRecordedStructure{
				DatePeriod: "2006-01-02",
			},
		},
		{ // 6
			Input: "0 EVEN\n1 PLAC A",
			Output: EventsRecordedStructure{
				SourceJurisdictionPlace: "A",
			},
		},
		{ // 7
			Input: "0 EVEN\n1 PLAC\n",
			Err:   ErrContext{"EventsRecordedStructure", cPLAC, ErrInvalidLength{"SourceJurisdictionPlace", "", 1, 120}},
		},
		{ // 8
			Input: "0 EVEN\n1 PLAC A\n1 PLAC B",
			Err:   ErrContext{"EventsRecordedStructure", cPLAC, ErrSingleMultiple},
		},
		{ // 9
			Input:   "0 EVEN\n1 PLAC A\n1 PLAC B",
			Options: []Option{AllowMoreThanAllowed},
			Output: EventsRecordedStructure{
				SourceJurisdictionPlace: "A",
			},
		},
		{ // 10
			Input: "0 EVEN\n1 UNKNOWN\n",
			Err:   ErrContext{"EventsRecordedStructure", "UNKNOWN", ErrUnknownTag},
		},
		{ // 11
			Input: "0 EVEN\n1 _UNKNOWN\n",
		},
		{ // 12
			Input:   "0 EVEN\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
		},
	})
}

func TestSubmissionRecord(t *testing.T) {
	test(t, []testData[SubmissionRecord]{
		{ // 1
			Input: "0 SUBM\n",
			Err:   ErrContext{"SubmissionRecord", "xrefID", ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 2
			Input: "0 @A@ SUBM\n",
			Output: SubmissionRecord{
				ID: "A",
			},
		},
		{ // 3
			Input: "0 @A@ SUBM\n1 FAMF B",
			Output: SubmissionRecord{
				ID:               "A",
				NameOfFamilyFile: "B",
			},
		},
		{ // 4
			Input: "0 @A@ SUBM\n1 FAMF\n",
			Err:   ErrContext{"SubmissionRecord", cFAMF, ErrInvalidLength{"NameOfFamilyFile", "", 1, 20}},
		},
		{ // 5
			Input: "0 @A@ SUBM\n1 FAMF B\n1 FAMF C",
			Err:   ErrContext{"SubmissionRecord", cFAMF, ErrSingleMultiple},
		},
		{ // 6
			Input:   "0 @A@ SUBM\n1 FAMF B\n1 FAMF C",
			Options: []Option{AllowMoreThanAllowed},
			Output: SubmissionRecord{
				ID:               "A",
				NameOfFamilyFile: "B",
			},
		},
		{ // 7
			Input: "0 @A@ SUBM\n1 TEMP BBBB",
			Output: SubmissionRecord{
				ID:         "A",
				TempleCode: "BBBB",
			},
		},
		{ // 8
			Input: "0 @A@ SUBM\n1 TEMP\n",
			Err:   ErrContext{"SubmissionRecord", cTEMP, ErrInvalidLength{"TempleCode", "", 4, 5}},
		},
		{ // 9
			Input: "0 @A@ SUBM\n1 TEMP BBBB\n1 TEMP CCCC",
			Err:   ErrContext{"SubmissionRecord", cTEMP, ErrSingleMultiple},
		},
		{ // 10
			Input:   "0 @A@ SUBM\n1 TEMP BBBB\n1 TEMP CCCC",
			Options: []Option{AllowMoreThanAllowed},
			Output: SubmissionRecord{
				ID:         "A",
				TempleCode: "BBBB",
			},
		},
		{ // 11
			Input: "0 @A@ SUBM\n1 ANCE 1",
			Output: SubmissionRecord{
				ID:                     "A",
				GenerationsOfAncestors: 1,
			},
		},
		{ // 12
			Input: "0 @A@ SUBM\n1 ANCE\n",
			Err:   ErrContext{"SubmissionRecord", cANCE, ErrInvalidLength{"GenerationsOfAncestors", "", 1, 4}},
		},
		{ // 13
			Input: "0 @A@ SUBM\n1 ANCE 1\n1 ANCE 2",
			Err:   ErrContext{"SubmissionRecord", cANCE, ErrSingleMultiple},
		},
		{ // 14
			Input:   "0 @A@ SUBM\n1 ANCE 1\n1 ANCE 2",
			Options: []Option{AllowMoreThanAllowed},
			Output: SubmissionRecord{
				ID:                     "A",
				GenerationsOfAncestors: 1,
			},
		},
		{ // 15
			Input: "0 @A@ SUBM\n1 DESC 1",
			Output: SubmissionRecord{
				ID:                       "A",
				GenerationsOfDescendants: 1,
			},
		},
		{ // 16
			Input: "0 @A@ SUBM\n1 DESC\n",
			Err:   ErrContext{"SubmissionRecord", cDESC, ErrInvalidLength{"GenerationsOfDescendants", "", 1, 4}},
		},
		{ // 17
			Input: "0 @A@ SUBM\n1 DESC 1\n1 DESC 2",
			Err:   ErrContext{"SubmissionRecord", cDESC, ErrSingleMultiple},
		},
		{ // 18
			Input:   "0 @A@ SUBM\n1 DESC 1\n1 DESC 2",
			Options: []Option{AllowMoreThanAllowed},
			Output: SubmissionRecord{
				ID:                       "A",
				GenerationsOfDescendants: 1,
			},
		},
		{ // 19
			Input: "0 @A@ SUBM\n1 ORDI yes",
			Output: SubmissionRecord{
				ID:                   "A",
				OrdinanceProcessFlag: "yes",
			},
		},
		{ // 20
			Input: "0 @A@ SUBM\n1 ORDI\n",
			Err:   ErrContext{"SubmissionRecord", cORDI, ErrInvalidValue{"OrdinanceProcessFlag", ""}},
		},
		{ // 21
			Input: "0 @A@ SUBM\n1 ORDI yes\n1 ORDI no",
			Err:   ErrContext{"SubmissionRecord", cORDI, ErrSingleMultiple},
		},
		{ // 22
			Input:   "0 @A@ SUBM\n1 ORDI yes\n1 ORDI no",
			Options: []Option{AllowMoreThanAllowed},
			Output: SubmissionRecord{
				ID:                   "A",
				OrdinanceProcessFlag: "yes",
			},
		},
		{ // 23
			Input: "0 @A@ SUBM\n1 RIN B",
			Output: SubmissionRecord{
				ID:                "A",
				AutomatedRecordID: "B",
			},
		},
		{ // 24
			Input: "0 @A@ SUBM\n1 RIN\n",
			Err:   ErrContext{"SubmissionRecord", cRIN, ErrInvalidLength{"AutomatedRecordID", "", 1, 12}},
		},
		{ // 25
			Input: "0 @A@ SUBM\n1 RIN B\n1 RIN C",
			Err:   ErrContext{"SubmissionRecord", cRIN, ErrSingleMultiple},
		},
		{ // 26
			Input:   "0 @A@ SUBM\n1 RIN B\n1 RIN C",
			Options: []Option{AllowMoreThanAllowed},
			Output: SubmissionRecord{
				ID:                "A",
				AutomatedRecordID: "B",
			},
		},
		{ // 27
			Input: "0 @A@ SUBM\n1 UNKNOWN\n",
			Err:   ErrContext{"SubmissionRecord", "UNKNOWN", ErrUnknownTag},
		},
		{ // 28
			Input: "0 @A@ SUBM\n1 _UNKNOWN\n",
			Output: SubmissionRecord{
				ID: "A",
			},
		},
		{ // 29
			Input:   "0 @A@ SUBM\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: SubmissionRecord{
				ID: "A",
			},
		},
	})
}

func TestSubmitterRecord(t *testing.T) {
	test(t, []testData[SubmitterRecord]{
		{ // 1
			Input: "0 SUBN\n",
			Err:   ErrContext{"SubmitterRecord", "xrefID", ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 2
			Input: "0 @A@ SUBN\n",
			Err:   ErrContext{"SubmitterRecord", cNAME, ErrRequiredMissing},
		},
		{ // 3
			Input: "0 @A@ SUBN\n1 NAME B",
			Output: SubmitterRecord{
				ID:                 "A",
				SubmitterName:      "B",
				PhoneNumber:        make([]PhoneNumber, 0, 3),
				LanguagePreference: make([]LanguagePreference, 0, 3),
			},
		},
		{ // 4
			Input: "0 @A@ SUBN\n1 NAME\n",
			Err:   ErrContext{"SubmitterRecord", cNAME, ErrInvalidLength{"SubmitterName", "", 1, 60}},
		},
		{ // 5
			Input: "0 @A@ SUBN\n1 NAME B\n1 NAME C",
			Err:   ErrContext{"SubmitterRecord", cNAME, ErrSingleMultiple},
		},
		{ // 6
			Input:   "0 @A@ SUBN\n1 NAME B\n1 NAME C",
			Options: []Option{AllowMoreThanAllowed},
			Output: SubmitterRecord{
				ID:                 "A",
				SubmitterName:      "B",
				PhoneNumber:        make([]PhoneNumber, 0, 3),
				LanguagePreference: make([]LanguagePreference, 0, 3),
			},
		},
		{ // 7
			Input: "0 @A@ SUBN\n1 NAME B\n1 ADDR C",
			Output: SubmitterRecord{
				ID:            "A",
				SubmitterName: "B",
				Address: AddressStructure{
					AddressLine: "C",
				},
				PhoneNumber:        make([]PhoneNumber, 0, 3),
				LanguagePreference: make([]LanguagePreference, 0, 3),
			},
		},
		{ // 8
			Input: "0 @A@ SUBN\n1 NAME B\n1 ADDR\n",
			Err:   ErrContext{"SubmitterRecord", cADDR, ErrContext{"AddressStructure", "line_value", ErrInvalidLength{"AddressLine", "", 1, 60}}},
		},
		{ // 9
			Input: "0 @A@ SUBN\n1 NAME B\n1 ADDR C\n1 ADDR D",
			Err:   ErrContext{"SubmitterRecord", cADDR, ErrSingleMultiple},
		},
		{ // 10
			Input:   "0 @A@ SUBN\n1 NAME B\n1 ADDR C\n1 ADDR D",
			Options: []Option{AllowMoreThanAllowed},
			Output: SubmitterRecord{
				ID:            "A",
				SubmitterName: "B",
				Address: AddressStructure{
					AddressLine: "C",
				},
				PhoneNumber:        make([]PhoneNumber, 0, 3),
				LanguagePreference: make([]LanguagePreference, 0, 3),
			},
		},
		{ // 11
			Input: "0 @A@ SUBN\n1 NAME B\n1 PHON number",
			Output: SubmitterRecord{
				ID:                 "A",
				SubmitterName:      "B",
				PhoneNumber:        []PhoneNumber{"number"},
				LanguagePreference: make([]LanguagePreference, 0, 3),
			},
		},
		{ // 12
			Input: "0 @A@ SUBN\n1 NAME B\n1 PHON\n",
			Err:   ErrContext{"SubmitterRecord", cPHON, ErrInvalidLength{"PhoneNumber", "", 1, 25}},
		},
		{ // 13
			Input: "0 @A@ SUBN\n1 NAME B\n1 PHON number 1\n1 PHON number 2",
			Output: SubmitterRecord{
				ID:                 "A",
				SubmitterName:      "B",
				PhoneNumber:        []PhoneNumber{"number 1", "number 2"},
				LanguagePreference: make([]LanguagePreference, 0, 3),
			},
		},
		{ // 14
			Input: "0 @A@ SUBN\n1 NAME B\n1 PHON number 1\n1 PHON number 2\n1 PHON number 3",
			Output: SubmitterRecord{
				ID:                 "A",
				SubmitterName:      "B",
				PhoneNumber:        []PhoneNumber{"number 1", "number 2", "number 3"},
				LanguagePreference: make([]LanguagePreference, 0, 3),
			},
		},
		{ // 15
			Input: "0 @A@ SUBN\n1 NAME B\n1 PHON number 1\n1 PHON number 2\n1 PHON number 3\n1 PHON number 4",
			Err:   ErrContext{"SubmitterRecord", cPHON, ErrTooMany(3)},
		},
		{ // 16
			Input:   "0 @A@ SUBN\n1 NAME B\n1 PHON number 1\n1 PHON number 2\n1 PHON number 3\n1 PHON number 4",
			Options: []Option{AllowMoreThanAllowed},
			Output: SubmitterRecord{
				ID:                 "A",
				SubmitterName:      "B",
				PhoneNumber:        []PhoneNumber{"number 1", "number 2", "number 3"},
				LanguagePreference: make([]LanguagePreference, 0, 3),
			},
		},
		{ // 17
			Input: "0 @A@ SUBN\n1 NAME B\n1 @A@ OBJE\n",
			Output: SubmitterRecord{
				ID:                 "A",
				SubmitterName:      "B",
				PhoneNumber:        make([]PhoneNumber, 0, 3),
				LanguagePreference: make([]LanguagePreference, 0, 3),
				Multimedia: []MultimediaLink{
					{
						Data: &MultimediaLinkID{
							ID: "A",
						},
					},
				},
			},
		},
		{ // 18
			Input: "0 @A@ SUBN\n1 NAME B\n1 OBJE\n",
			Err:   ErrContext{"SubmitterRecord", cOBJE, ErrContext{"MultimediaLinkFile", cFORM, ErrRequiredMissing}},
		},
		{ // 19
			Input: "0 @A@ SUBN\n1 NAME B\n1 @A@ OBJE\n1 @B@ OBJE\n",
			Output: SubmitterRecord{
				ID:                 "A",
				SubmitterName:      "B",
				PhoneNumber:        make([]PhoneNumber, 0, 3),
				LanguagePreference: make([]LanguagePreference, 0, 3),
				Multimedia: []MultimediaLink{
					{
						Data: &MultimediaLinkID{
							ID: "A",
						},
					},
					{
						Data: &MultimediaLinkID{
							ID: "B",
						},
					},
				},
			},
		},
		{ // 20
			Input: "0 @A@ SUBN\n1 NAME B\n1 LANG C",
			Output: SubmitterRecord{
				ID:            "A",
				SubmitterName: "B",
				PhoneNumber:   make([]PhoneNumber, 0, 3),
				LanguagePreference: []LanguagePreference{
					"C",
				},
			},
		},
		{ // 21
			Input: "0 @A@ SUBN\n1 NAME B\n1 LANG\n",
			Err:   ErrContext{"SubmitterRecord", cLANG, ErrInvalidLength{"LanguagePreference", "", 1, 90}},
		},
		{ // 22
			Input: "0 @A@ SUBN\n1 NAME B\n1 LANG C\n1 LANG D",
			Output: SubmitterRecord{
				ID:            "A",
				SubmitterName: "B",
				PhoneNumber:   make([]PhoneNumber, 0, 3),
				LanguagePreference: []LanguagePreference{
					"C",
					"D",
				},
			},
		},
		{ // 23
			Input: "0 @A@ SUBN\n1 NAME B\n1 LANG C\n1 LANG D\n1 LANG E",
			Output: SubmitterRecord{
				ID:            "A",
				SubmitterName: "B",
				PhoneNumber:   make([]PhoneNumber, 0, 3),
				LanguagePreference: []LanguagePreference{
					"C",
					"D",
					"E",
				},
			},
		},
		{ // 24
			Input: "0 @A@ SUBN\n1 NAME B\n1 LANG C\n1 LANG D\n1 LANG E\n1 LANG F",
			Err:   ErrContext{"SubmitterRecord", cLANG, ErrTooMany(3)},
		},
		{ // 25
			Input:   "0 @A@ SUBN\n1 NAME B\n1 LANG C\n1 LANG D\n1 LANG E\n1 LANG F",
			Options: []Option{AllowMoreThanAllowed},
			Output: SubmitterRecord{
				ID:            "A",
				SubmitterName: "B",
				PhoneNumber:   make([]PhoneNumber, 0, 3),
				LanguagePreference: []LanguagePreference{
					"C",
					"D",
					"E",
				},
			},
		},
		{ // 26
			Input: "0 @A@ SUBN\n1 NAME B\n1 RFN C",
			Output: SubmitterRecord{
				ID:                     "A",
				SubmitterName:          "B",
				PhoneNumber:            make([]PhoneNumber, 0, 3),
				LanguagePreference:     make([]LanguagePreference, 0, 3),
				SubmitterRegisteredRFN: "C",
			},
		},
		{ // 27
			Input: "0 @A@ SUBN\n1 NAME B\n1 RFN\n",
			Err:   ErrContext{"SubmitterRecord", cRFN, ErrInvalidLength{"SubmitterRegisteredRFN", "", 1, 30}},
		},
		{ // 28
			Input: "0 @A@ SUBN\n1 NAME B\n1 RFN C\n1 RFN D",
			Err:   ErrContext{"SubmitterRecord", cRFN, ErrSingleMultiple},
		},
		{ // 29
			Input:   "0 @A@ SUBN\n1 NAME B\n1 RFN C\n1 RFN D",
			Options: []Option{AllowMoreThanAllowed},
			Output: SubmitterRecord{
				ID:                     "A",
				SubmitterName:          "B",
				PhoneNumber:            make([]PhoneNumber, 0, 3),
				LanguagePreference:     make([]LanguagePreference, 0, 3),
				SubmitterRegisteredRFN: "C",
			},
		},
		{ // 30
			Input: "0 @A@ SUBN\n1 NAME B\n1 RIN C",
			Output: SubmitterRecord{
				ID:                 "A",
				SubmitterName:      "B",
				PhoneNumber:        make([]PhoneNumber, 0, 3),
				LanguagePreference: make([]LanguagePreference, 0, 3),
				AutomatedRecordID:  "C",
			},
		},
		{ // 31
			Input: "0 @A@ SUBN\n1 NAME B\n1 RIN\n",
			Err:   ErrContext{"SubmitterRecord", cRIN, ErrInvalidLength{"AutomatedRecordID", "", 1, 12}},
		},
		{ // 32
			Input: "0 @A@ SUBM\n1 RIN B\n1 RIN C",
			Err:   ErrContext{"SubmitterRecord", cRIN, ErrSingleMultiple},
		},
		{ // 33
			Input:   "0 @A@ SUBN\n1 NAME B\n1 RIN C\n1 RIN D",
			Options: []Option{AllowMoreThanAllowed},
			Output: SubmitterRecord{
				ID:                 "A",
				SubmitterName:      "B",
				PhoneNumber:        make([]PhoneNumber, 0, 3),
				LanguagePreference: make([]LanguagePreference, 0, 3),
				AutomatedRecordID:  "C",
			},
		},
		{ // 34
			Input: "0 @A@ SUBN\n1 NAME B\n1 CHAN\n2 DATE 2006-05-04",
			Output: SubmitterRecord{
				ID:                 "A",
				SubmitterName:      "B",
				PhoneNumber:        make([]PhoneNumber, 0, 3),
				LanguagePreference: make([]LanguagePreference, 0, 3),
				ChangeDate: ChangeDateStructure{
					Date: ChangeDateTime{
						ChangeDate: "2006-05-04",
					},
				},
			},
		},
		{ // 35
			Input: "0 @A@ SUBN\n1 NAME B\n1 CHAN\n",
			Err:   ErrContext{"SubmitterRecord", cCHAN, ErrContext{"ChangeDateStructure", cDATE, ErrRequiredMissing}},
		},
		{ // 36
			Input: "0 @A@ SUBN\n1 NAME B\n1 CHAN\n2 DATE 2006-05-04\n1 CHAN\n2 DATE 2000-01-01",
			Err:   ErrContext{"SubmitterRecord", cCHAN, ErrSingleMultiple},
		},
		{ // 37
			Input:   "0 @A@ SUBN\n1 NAME B\n1 CHAN\n2 DATE 2006-05-04\n1 CHAN\n2 DATE 2000-01-01",
			Options: []Option{AllowMoreThanAllowed},
			Output: SubmitterRecord{
				ID:                 "A",
				SubmitterName:      "B",
				PhoneNumber:        make([]PhoneNumber, 0, 3),
				LanguagePreference: make([]LanguagePreference, 0, 3),
				ChangeDate: ChangeDateStructure{
					Date: ChangeDateTime{
						ChangeDate: "2006-05-04",
					},
				},
			},
		},
		{ // 38
			Input: "0 @A@ SUBN\n1 NAME B\n1 UNKNOWN\n",
			Err:   ErrContext{"SubmitterRecord", "UNKNOWN", ErrUnknownTag},
		},
		{ // 39
			Input: "0 @A@ SUBN\n1 NAME B\n1 _UNKNOWN\n",
			Output: SubmitterRecord{
				ID:                 "A",
				SubmitterName:      "B",
				PhoneNumber:        make([]PhoneNumber, 0, 3),
				LanguagePreference: make([]LanguagePreference, 0, 3),
			},
		},
		{ // 40
			Input:   "0 @A@ SUBN\n1 NAME B\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: SubmitterRecord{
				ID:                 "A",
				SubmitterName:      "B",
				PhoneNumber:        make([]PhoneNumber, 0, 3),
				LanguagePreference: make([]LanguagePreference, 0, 3),
			},
		},
	})
}

func TestAddressStructure(t *testing.T) {
	test(t, []testData[AddressStructure]{
		{ // 1
			Input: "0 ADDR\n",
			Err:   ErrContext{"AddressStructure", "line_value", ErrInvalidLength{"AddressLine", "", 1, 60}},
		},
		{ // 2
			Input: "0 ADDR A",
			Output: AddressStructure{
				AddressLine: "A",
			},
		},
		{ // 3
			Input: "0 ADDR A\n1 ADR1 B",
			Output: AddressStructure{
				AddressLine:  "A",
				AddressLine1: "B",
			},
		},
		{ // 4
			Input: "0 ADDR A\n1 ADR1\n",
			Err:   ErrContext{"AddressStructure", cADR1, ErrInvalidLength{"AddressLine1", "", 1, 60}},
		},
		{ // 5
			Input: "0 ADDR A\n1 ADR1 B\n1 ADR1 C",
			Err:   ErrContext{"AddressStructure", cADR1, ErrSingleMultiple},
		},
		{ // 6
			Input:   "0 ADDR A\n1 ADR1 B\n1 ADR1 C",
			Options: []Option{AllowMoreThanAllowed},
			Output: AddressStructure{
				AddressLine:  "A",
				AddressLine1: "B",
			},
		},
		{ // 7
			Input: "0 ADDR A\n1 ADR2 B",
			Output: AddressStructure{
				AddressLine:  "A",
				AddressLine2: "B",
			},
		},
		{ // 8
			Input: "0 ADDR A\n1 ADR2\n",
			Err:   ErrContext{"AddressStructure", cADR2, ErrInvalidLength{"AddressLine2", "", 1, 60}},
		},
		{ // 9
			Input: "0 ADDR A\n1 ADR2 B\n1 ADR2 C",
			Err:   ErrContext{"AddressStructure", cADR2, ErrSingleMultiple},
		},
		{ // 10
			Input:   "0 ADDR A\n1 ADR2 B\n1 ADR2 C",
			Options: []Option{AllowMoreThanAllowed},
			Output: AddressStructure{
				AddressLine:  "A",
				AddressLine2: "B",
			},
		},
		{ // 11
			Input: "0 ADDR A\n1 CITY B",
			Output: AddressStructure{
				AddressLine: "A",
				City:        "B",
			},
		},
		{ // 12
			Input: "0 ADDR A\n1 CITY\n",
			Err:   ErrContext{"AddressStructure", cCITY, ErrInvalidLength{"AddressCity", "", 1, 60}},
		},
		{ // 13
			Input: "0 ADDR A\n1 CITY B\n1 CITY C",
			Err:   ErrContext{"AddressStructure", cCITY, ErrSingleMultiple},
		},
		{ // 14
			Input:   "0 ADDR A\n1 CITY B\n1 CITY C",
			Options: []Option{AllowMoreThanAllowed},
			Output: AddressStructure{
				AddressLine: "A",
				City:        "B",
			},
		},
		{ // 15
			Input: "0 ADDR A\n1 STAE B",
			Output: AddressStructure{
				AddressLine: "A",
				State:       "B",
			},
		},
		{ // 16
			Input: "0 ADDR A\n1 STAE\n",
			Err:   ErrContext{"AddressStructure", cSTAE, ErrInvalidLength{"AddressState", "", 1, 60}},
		},
		{ // 17
			Input: "0 ADDR A\n1 STAE B\n1 STAE C",
			Err:   ErrContext{"AddressStructure", cSTAE, ErrSingleMultiple},
		},
		{ // 18
			Input:   "0 ADDR A\n1 STAE B\n1 STAE C",
			Options: []Option{AllowMoreThanAllowed},
			Output: AddressStructure{
				AddressLine: "A",
				State:       "B",
			},
		},
		{ // 19
			Input: "0 ADDR A\n1 POST B",
			Output: AddressStructure{
				AddressLine: "A",
				PostalCode:  "B",
			},
		},
		{ // 20
			Input: "0 ADDR A\n1 POST\n",
			Err:   ErrContext{"AddressStructure", cPOST, ErrInvalidLength{"AddressPostalCode", "", 1, 10}},
		},
		{ // 21
			Input: "0 ADDR A\n1 POST B\n1 POST C",
			Err:   ErrContext{"AddressStructure", cPOST, ErrSingleMultiple},
		},
		{ // 22
			Input:   "0 ADDR A\n1 POST B\n1 POST C",
			Options: []Option{AllowMoreThanAllowed},
			Output: AddressStructure{
				AddressLine: "A",
				PostalCode:  "B",
			},
		},
		{ // 23
			Input: "0 ADDR A\n1 CTRY B",
			Output: AddressStructure{
				AddressLine: "A",
				Country:     "B",
			},
		},
		{ // 24
			Input: "0 ADDR A\n1 CTRY\n",
			Err:   ErrContext{"AddressStructure", cCTRY, ErrInvalidLength{"AddressCountry", "", 1, 60}},
		},
		{ // 25
			Input: "0 ADDR A\n1 CTRY B\n1 CTRY C",
			Err:   ErrContext{"AddressStructure", cCTRY, ErrSingleMultiple},
		},
		{ // 26
			Input:   "0 ADDR A\n1 CTRY B\n1 CTRY C",
			Options: []Option{AllowMoreThanAllowed},
			Output: AddressStructure{
				AddressLine: "A",
				Country:     "B",
			},
		},
		{ // 27
			Input: "0 ADDR A\n1 UNKNOWN\n",
			Err:   ErrContext{"AddressStructure", "UNKNOWN", ErrUnknownTag},
		},
		{ // 28
			Input: "0 ADDR A\n1 _UNKNOWN\n",
			Output: AddressStructure{
				AddressLine: "A",
			},
		},
		{ // 29
			Input:   "0 ADDR A\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: AddressStructure{
				AddressLine: "A",
			},
		},
	})
}

func TestAssociationStructure(t *testing.T) {
	test(t, []testData[AssociationStructure]{
		{ // 1
			Input: "0 ASSO\n",
			Err:   ErrContext{"AssociationStructure", "xrefID", ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 2
			Input: "0 @A@ ASSO\n",
			Err:   ErrContext{"AssociationStructure", cRELA, ErrRequiredMissing},
		},
		{ // 3
			Input: "0 @A@ ASSO\n1 RELA B\n",
			Err:   ErrContext{"AssociationStructure", cTYPE, ErrRequiredMissing},
		},
		{ // 4
			Input: "0 @A@ ASSO\n1 RELA B\n1 TYPE FAM",
			Output: AssociationStructure{
				ID:         "A",
				Relation:   "B",
				RecordType: cFAM,
			},
		},
		{ // 5
			Input: "0 @A@ ASSO\n1 RELA B\n1 TYPE\n",
			Err:   ErrContext{"AssociationStructure", cTYPE, ErrInvalidValue{"RecordType", ""}},
		},
		{ // 6
			Input: "0 @A@ ASSO\n1 RELA B\n1 TYPE FAM\n1 TYPE INDI\n",
			Err:   ErrContext{"AssociationStructure", cTYPE, ErrSingleMultiple},
		},
		{ // 7
			Input:   "0 @A@ ASSO\n1 RELA B\n1 TYPE FAM\n1 TYPE INDI\n",
			Options: []Option{AllowMoreThanAllowed},
			Output: AssociationStructure{
				ID:         "A",
				Relation:   "B",
				RecordType: cFAM,
			},
		},
		{ // 8
			Input: "0 @A@ ASSO\n1 RELA\n1 TYPE FAM\n",
			Err:   ErrContext{"AssociationStructure", cRELA, ErrInvalidLength{"RelationIsDescriptor", "", 1, 25}},
		},
		{ // 9
			Input: "0 @A@ ASSO\n1 RELA B\n1 TYPE FAM\n1 RELA C",
			Err:   ErrContext{"AssociationStructure", cRELA, ErrSingleMultiple},
		},
		{ // 10
			Input:   "0 @A@ ASSO\n1 RELA B\n1 TYPE FAM\n1 RELA C",
			Options: []Option{AllowMoreThanAllowed},
			Output: AssociationStructure{
				ID:         "A",
				Relation:   "B",
				RecordType: cFAM,
			},
		},
		{ // 11
			Input: "0 @A@ ASSO\n1 RELA B\n1 TYPE FAM\n1 @A@ NOTE\n",
			Output: AssociationStructure{
				ID:         "A",
				Relation:   "B",
				RecordType: cFAM,
				Notes: []NoteStructure{
					{
						Data: &NoteID{
							ID: "A",
						},
					},
				},
			},
		},
		{ // 12
			Input: "0 @A@ ASSO\n1 RELA B\n1 TYPE FAM\n1 NOTE\n",
			Err:   ErrContext{"AssociationStructure", cNOTE, ErrContext{"NoteText", "line_value", ErrInvalidLength{"SubmitterText", "", 1, 248}}},
		},
		{ // 13
			Input: "0 @A@ ASSO\n1 RELA B\n1 TYPE FAM\n1 @A@ NOTE\n1 @B@ NOTE\n",
			Output: AssociationStructure{
				ID:         "A",
				Relation:   "B",
				RecordType: cFAM,
				Notes: []NoteStructure{
					{
						Data: &NoteID{
							ID: "A",
						},
					},
					{
						Data: &NoteID{
							ID: "B",
						},
					},
				},
			},
		},
		{ // 14
			Input: "0 @A@ ASSO\n1 RELA B\n1 TYPE FAM\n1 @C@ SOUR\n",
			Output: AssociationStructure{
				ID:         "A",
				Relation:   "B",
				RecordType: cFAM,
				Sources: []SourceCitation{
					{
						Data: &SourceID{
							ID: "C",
						},
					},
				},
			},
		},
		{ // 15
			Input: "0 @A@ ASSO\n1 RELA B\n1 TYPE FAM\n1 SOUR\n",
			Err:   ErrContext{"AssociationStructure", cSOUR, ErrContext{"SourceText", "line_value", ErrInvalidLength{"SourceDescription", "", 1, 248}}},
		},
		{ // 16
			Input: "0 @A@ ASSO\n1 RELA B\n1 TYPE FAM\n1 @C@ SOUR\n1 @D@ SOUR\n",
			Output: AssociationStructure{
				ID:         "A",
				Relation:   "B",
				RecordType: cFAM,
				Sources: []SourceCitation{
					{
						Data: &SourceID{
							ID: "C",
						},
					},
					{
						Data: &SourceID{
							ID: "D",
						},
					},
				},
			},
		},
		{ // 17
			Input: "0 @A@ ASSO\n1 RELA B\n1 TYPE FAM\n1 UNKNOWN\n",
			Err:   ErrContext{"AssociationStructure", "UNKNOWN", ErrUnknownTag},
		},
		{ // 18
			Input: "0 @A@ ASSO\n1 RELA B\n1 TYPE FAM\n1 _UNKNOWN\n",
			Output: AssociationStructure{
				ID:         "A",
				Relation:   "B",
				RecordType: cFAM,
			},
		},
		{ // 19
			Input:   "0 @A@ ASSO\n1 RELA B\n1 TYPE FAM\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: AssociationStructure{
				ID:         "A",
				Relation:   "B",
				RecordType: cFAM,
			},
		},
	})
}

func TestChangeDateStructure(t *testing.T) {
	test(t, []testData[ChangeDateStructure]{
		{ // 1
			Input: "0 CHAN\n",
			Err:   ErrContext{"ChangeDateStructure", cDATE, ErrRequiredMissing},
		},
		{ // 2
			Input: "0 CHAN\n1 DATE 2006-05-01",
			Output: ChangeDateStructure{
				Date: ChangeDateTime{
					ChangeDate: "2006-05-01",
				},
			},
		},
		{ // 3
			Input: "0 CHAN\n1 DATE\n",
			Err:   ErrContext{"ChangeDateStructure", cDATE, ErrContext{"ChangeDateTime", "line_value", ErrInvalidLength{"ChangeDate", "", 10, 11}}},
		},
		{ // 4
			Input: "0 CHAN\n1 DATE 2006-05-01\n1 DATE 2007-06-02",
			Err:   ErrContext{"ChangeDateStructure", cDATE, ErrSingleMultiple},
		},
		{ // 5
			Input:   "0 CHAN\n1 DATE 2006-05-01\n1 DATE 2007-06-02",
			Options: []Option{AllowMoreThanAllowed},
			Output: ChangeDateStructure{
				Date: ChangeDateTime{
					ChangeDate: "2006-05-01",
				},
			},
		},
		{ // 6
			Input: "0 CHAN\n1 DATE 2006-05-01\n1 @A@ NOTE\n",
			Output: ChangeDateStructure{
				Date: ChangeDateTime{
					ChangeDate: "2006-05-01",
				},
				Notes: []NoteStructure{
					{
						Data: &NoteID{
							ID: "A",
						},
					},
				},
			},
		},
		{ // 7
			Input: "0 CHAN\n1 DATE 2006-05-01\n1 NOTE\n",
			Err:   ErrContext{"ChangeDateStructure", cNOTE, ErrContext{"NoteText", "line_value", ErrInvalidLength{"SubmitterText", "", 1, 248}}},
		},
		{ // 8
			Input: "0 CHAN\n1 DATE 2006-05-01\n1 @A@ NOTE\n1 @B@ NOTE\n",
			Output: ChangeDateStructure{
				Date: ChangeDateTime{
					ChangeDate: "2006-05-01",
				},
				Notes: []NoteStructure{
					{
						Data: &NoteID{
							ID: "A",
						},
					},
					{
						Data: &NoteID{
							ID: "B",
						},
					},
				},
			},
		},
		{ // 9
			Input: "0 CHAN\n1 DATE 2006-05-01\n1 UNKNOWN\n",
			Err:   ErrContext{"ChangeDateStructure", "UNKNOWN", ErrUnknownTag},
		},
		{ // 10
			Input: "0 CHAN\n1 DATE 2006-05-01\n1 _UNKNOWN\n",
			Output: ChangeDateStructure{
				Date: ChangeDateTime{
					ChangeDate: "2006-05-01",
				},
			},
		},
		{ // 11
			Input:   "0 CHAN\n1 DATE 2006-05-01\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: ChangeDateStructure{
				Date: ChangeDateTime{
					ChangeDate: "2006-05-01",
				},
			},
		},
	})
}

func TestChangeDateTime(t *testing.T) {
	test(t, []testData[ChangeDateTime]{
		{ // 1
			Input: "0 DATE\n",
			Err:   ErrContext{"ChangeDateTime", "line_value", ErrInvalidLength{"ChangeDate", "", 10, 11}},
		},
		{ // 2
			Input: "0 DATE 2006-05-01",
			Output: ChangeDateTime{
				ChangeDate: "2006-05-01",
			},
		},
		{ // 3
			Input: "0 DATE 2006-05-01\n1 TIME 12:00",
			Output: ChangeDateTime{
				ChangeDate: "2006-05-01",
				Time:       "12:00",
			},
		},
		{ // 4
			Input: "0 DATE 2006-05-01\n1 TIME\n",
			Err:   ErrContext{"ChangeDateTime", cTIME, ErrInvalidLength{"TimeValue", "", 1, 12}},
		},
		{ // 5
			Input: "0 DATE 2006-05-01\n1 TIME 12:00\n1 TIME 13:00",
			Err:   ErrContext{"ChangeDateTime", cTIME, ErrSingleMultiple},
		},
		{ // 6
			Input:   "0 DATE 2006-05-01\n1 TIME 12:00\n1 TIME 13:00",
			Options: []Option{AllowMoreThanAllowed},
			Output: ChangeDateTime{
				ChangeDate: "2006-05-01",
				Time:       "12:00",
			},
		},
		{ // 7
			Input: "0 DATE 2006-05-01\n1 UNKNOWN\n",
			Err:   ErrContext{"ChangeDateTime", "UNKNOWN", ErrUnknownTag},
			Output: ChangeDateTime{
				ChangeDate: "2006-05-01",
			},
		},
		{ // 8
			Input: "0 DATE 2006-05-01\n1 _UNKNOWN\n",
			Output: ChangeDateTime{
				ChangeDate: "2006-05-01",
			},
		},
		{ // 9
			Input:   "0 DATE 2006-05-01\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: ChangeDateTime{
				ChangeDate: "2006-05-01",
			},
		},
	})
}

func TestChildToFamilyLink(t *testing.T) {
	test(t, []testData[ChildToFamilyLink]{
		{ // 1
			Input: "0 FAMC\n",
			Err:   ErrContext{"ChildToFamilyLink", "line_value", ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 2
			Input: "0 DATE @A@",
			Output: ChildToFamilyLink{
				ID: "A",
			},
		},
		{ // 3
			Input: "0 DATE @A@\n1 PEDI BIRTH",
			Output: ChildToFamilyLink{
				ID: "A",
				PedigreeLinkageType: []PedigreeLinkageType{
					cbirth,
				},
			},
		},
		{ // 4
			Input: "0 DATE @A@\n1 PEDI\n",
			Err:   ErrContext{"ChildToFamilyLink", cPEDI, ErrInvalidValue{"PedigreeLinkageType", ""}},
		},
		{ // 5
			Input: "0 DATE @A@\n1 PEDI BIRTH\n1 PEDI FOSTER\n",
			Output: ChildToFamilyLink{
				ID: "A",
				PedigreeLinkageType: []PedigreeLinkageType{
					cbirth,
					cfoster,
				},
			},
		},
		{ // 6
			Input: "0 DATE @A@\n1 @A@ NOTE\n",
			Output: ChildToFamilyLink{
				ID: "A",
				Notes: []NoteStructure{
					{
						Data: &NoteID{
							ID: "A",
						},
					},
				},
			},
		},
		{ // 7
			Input: "0 DATE @A@\n1 NOTE\n",
			Err:   ErrContext{"ChildToFamilyLink", cNOTE, ErrContext{"NoteText", "line_value", ErrInvalidLength{"SubmitterText", "", 1, 248}}},
		},
		{ // 8
			Input: "0 DATE @A@\n1 @A@ NOTE\n1 @B@ NOTE\n",
			Output: ChildToFamilyLink{
				ID: "A",
				Notes: []NoteStructure{
					{
						Data: &NoteID{
							ID: "A",
						},
					},
					{
						Data: &NoteID{
							ID: "B",
						},
					},
				},
			},
		},
		{ // 9
			Input: "0 DATE @A@\n1 UNKNOWN\n",
			Err:   ErrContext{"ChildToFamilyLink", "UNKNOWN", ErrUnknownTag},
		},
		{ // 10
			Input: "0 DATE @A@\n1 _UNKNOWN\n",
			Output: ChildToFamilyLink{
				ID: "A",
			},
		},
		{ // 11
			Input:   "0 DATE @A@\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: ChildToFamilyLink{
				ID: "A",
			},
		},
	})
}

func TestEventDetail(t *testing.T) {
	test(t, []testData[EventDetail]{
		{ // 1
			Input: "0 EVEN\n",
			Output: EventDetail{
				PhoneNumber: make([]PhoneNumber, 0, 3),
			},
		},
		{ // 2
			Input: "0 EVEN\n1 TYPE A",
			Output: EventDetail{
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Type:        "A",
			},
		},
		{ // 3
			Input: "0 EVEN\n1 TYPE\n",
			Err:   ErrContext{"EventDetail", cTYPE, ErrInvalidLength{"EventDescriptor", "", 1, 90}},
		},
		{ // 4
			Input: "0 EVEN\n1 TYPE A\n1 TYPE B",
			Err:   ErrContext{"EventDetail", cTYPE, ErrSingleMultiple},
		},
		{ // 5
			Input:   "0 EVEN\n1 TYPE A\n1 TYPE B",
			Options: []Option{AllowMoreThanAllowed},
			Output: EventDetail{
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Type:        "A",
			},
		},
		{ // 6
			Input: "0 EVEN\n1 TYPE A\n1 DATE B",
			Output: EventDetail{
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Type:        "A",
				Date:        "B",
			},
		},
		{ // 7
			Input: "0 EVEN\n1 TYPE A\n1 DATE\n",
			Err:   ErrContext{"EventDetail", cDATE, ErrInvalidLength{"DateValue", "", 1, 35}},
		},
		{ // 8
			Input: "0 EVEN\n1 TYPE A\n1 DATE B \n1 DATE C",
			Err:   ErrContext{"EventDetail", cDATE, ErrSingleMultiple},
		},
		{ // 9
			Input:   "0 EVEN\n1 TYPE A\n1 DATE B \n1 DATE C",
			Options: []Option{AllowMoreThanAllowed},
			Output: EventDetail{
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Type:        "A",
				Date:        "B",
			},
		},
		{ // 10
			Input: "0 EVEN\n1 TYPE A\n1 PLAC B",
			Output: EventDetail{
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Type:        "A",
				Place: PlaceStructure{
					Place: "B",
				},
			},
		},
		{ // 11
			Input: "0 EVEN\n1 TYPE A\n1 PLAC\n",
			Err:   ErrContext{"EventDetail", cPLAC, ErrContext{"PlaceStructure", "line_value", ErrInvalidLength{"PlaceValue", "", 1, 120}}},
		},
		{ // 12
			Input: "0 EVEN\n1 TYPE A\n1 PLAC B\n1 PLAC C",
			Err:   ErrContext{"EventDetail", cPLAC, ErrSingleMultiple},
		},
		{ // 13
			Input:   "0 EVEN\n1 TYPE A\n1 PLAC B\n1 PLAC C",
			Options: []Option{AllowMoreThanAllowed},
			Output: EventDetail{
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Type:        "A",
				Place: PlaceStructure{
					Place: "B",
				},
			},
		},
		{ // 14
			Input: "0 EVEN\n1 TYPE A\n1 ADDR B",
			Output: EventDetail{
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Type:        "A",
				Address: AddressStructure{
					AddressLine: "B",
				},
			},
		},
		{ // 15
			Input: "0 EVEN\n1 TYPE A\n1 ADDR\n",
			Err:   ErrContext{"EventDetail", cADDR, ErrContext{"AddressStructure", "line_value", ErrInvalidLength{"AddressLine", "", 1, 60}}},
		},
		{ // 16
			Input: "0 EVEN\n1 TYPE A\n1 ADDR B\n1 ADDR C",
			Err:   ErrContext{"EventDetail", cADDR, ErrSingleMultiple},
		},
		{ // 17
			Input:   "0 EVEN\n1 TYPE A\n1 ADDR B\n1 ADDR C",
			Options: []Option{AllowMoreThanAllowed},
			Output: EventDetail{
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Type:        "A",
				Address: AddressStructure{
					AddressLine: "B",
				},
			},
		},
		{ // 18
			Input: "0 EVEN\n1 TYPE A\n1 PHON number",
			Output: EventDetail{
				Type:        "A",
				PhoneNumber: []PhoneNumber{"number"},
			},
		},
		{ // 19
			Input: "0 EVEN\n1 TYPE A\n1 PHON\n",
			Err:   ErrContext{"EventDetail", cPHON, ErrInvalidLength{"PhoneNumber", "", 1, 25}},
		},
		{ // 20
			Input: "0 EVEN\n1 TYPE A\n1 PHON number 1\n1 PHON number 2",
			Output: EventDetail{
				Type:        "A",
				PhoneNumber: []PhoneNumber{"number 1", "number 2"},
			},
		},
		{ // 21
			Input: "0 EVEN\n1 TYPE A\n1 PHON number 1\n1 PHON number 2\n1 PHON number 3",
			Output: EventDetail{
				Type:        "A",
				PhoneNumber: []PhoneNumber{"number 1", "number 2", "number 3"},
			},
		},
		{ // 22
			Input: "0 EVEN\n1 TYPE A\n1 PHON number 1\n1 PHON number 2\n1 PHON number 3\n1 PHON number 4",
			Err:   ErrContext{"EventDetail", cPHON, ErrTooMany(3)},
		},
		{ // 23
			Input:   "0 EVEN\n1 TYPE A\n1 PHON number 1\n1 PHON number 2\n1 PHON number 3\n1 PHON number 4",
			Options: []Option{AllowMoreThanAllowed},
			Output: EventDetail{
				Type:        "A",
				PhoneNumber: []PhoneNumber{"number 1", "number 2", "number 3"},
			},
		},
		{ // 24
			Input: "0 EVEN\n1 TYPE A\n1 AGE B",
			Output: EventDetail{
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Type:        "A",
				Age:         "B",
			},
		},
		{ // 25
			Input: "0 EVEN\n1 TYPE A\n1 AGE\n",
			Err:   ErrContext{"EventDetail", cAGE, ErrInvalidLength{"AgeAtEvent", "", 1, 12}},
		},
		{ // 26
			Input: "0 EVEN\n1 TYPE A\n1 AGE B\n1 AGE C",
			Err:   ErrContext{"EventDetail", cAGE, ErrSingleMultiple},
		},
		{ // 27
			Input:   "0 EVEN\n1 TYPE A\n1 AGE B\n1 AGE C",
			Options: []Option{AllowMoreThanAllowed},
			Output: EventDetail{
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Type:        "A",
				Age:         "B",
			},
		},
		{ // 28
			Input: "0 EVEN\n1 TYPE A\n1 AGNC B",
			Output: EventDetail{
				PhoneNumber:       make([]PhoneNumber, 0, 3),
				Type:              "A",
				ResponsibleAgency: "B",
			},
		},
		{ // 29
			Input: "0 EVEN\n1 TYPE A\n1 AGNC\n",
			Err:   ErrContext{"EventDetail", cAGNC, ErrInvalidLength{"ResponsibleAgency", "", 1, 120}},
		},
		{ // 30
			Input: "0 EVEN\n1 TYPE A\n1 AGNC B\n1 AGNC C",
			Err:   ErrContext{"EventDetail", cAGNC, ErrSingleMultiple},
		},
		{ // 31
			Input:   "0 EVEN\n1 TYPE A\n1 AGNC B\n1 AGNC C",
			Options: []Option{AllowMoreThanAllowed},
			Output: EventDetail{
				PhoneNumber:       make([]PhoneNumber, 0, 3),
				Type:              "A",
				ResponsibleAgency: "B",
			},
		},
		{ // 32
			Input: "0 EVEN\n1 TYPE A\n1 CAUS B",
			Output: EventDetail{
				PhoneNumber:  make([]PhoneNumber, 0, 3),
				Type:         "A",
				CauseOfEvent: "B",
			},
		},
		{ // 33
			Input: "0 EVEN\n1 TYPE A\n1 CAUS\n",
			Err:   ErrContext{"EventDetail", cCAUS, ErrInvalidLength{"CauseOfEvent", "", 1, 90}},
		},
		{ // 34
			Input: "0 EVEN\n1 TYPE A\n1 CAUS B\n1 CAUS C",
			Err:   ErrContext{"EventDetail", cCAUS, ErrSingleMultiple},
		},
		{ // 35
			Input:   "0 EVEN\n1 TYPE A\n1 CAUS B\n1 CAUS C",
			Options: []Option{AllowMoreThanAllowed},
			Output: EventDetail{
				PhoneNumber:  make([]PhoneNumber, 0, 3),
				Type:         "A",
				CauseOfEvent: "B",
			},
		},
		{ // 36
			Input: "0 EVEN\n1 TYPE A\n1 @B@ SOUR\n",
			Output: EventDetail{
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Type:        "A",
				Sources: []SourceCitation{
					{
						Data: &SourceID{
							ID: "B",
						},
					},
				},
			},
		},
		{ // 37
			Input: "0 EVEN\n1 TYPE A\n1 SOUR\n",
			Err:   ErrContext{"EventDetail", cSOUR, ErrContext{"SourceText", "line_value", ErrInvalidLength{"SourceDescription", "", 1, 248}}},
		},
		{ // 38
			Input: "0 EVEN\n1 TYPE A\n1 @B@ SOUR\n1 @C@ SOUR\n",
			Output: EventDetail{
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Type:        "A",
				Sources: []SourceCitation{
					{
						Data: &SourceID{
							ID: "B",
						},
					},
					{
						Data: &SourceID{
							ID: "C",
						},
					},
				},
			},
		},
		{ // 39
			Input: "0 EVEN\n1 TYPE A\n1 @A@ OBJE\n",
			Output: EventDetail{
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Type:        "A",
				Multimedia: []MultimediaLink{
					{
						Data: &MultimediaLinkID{
							ID: "A",
						},
					},
				},
			},
		},
		{ // 40
			Input: "0 EVEN\n1 TYPE A\n1 OBJE\n",
			Err:   ErrContext{"EventDetail", cOBJE, ErrContext{"MultimediaLinkFile", cFORM, ErrRequiredMissing}},
		},
		{ // 41
			Input: "0 EVEN\n1 TYPE A\n1 @A@ OBJE\n1 @B@ OBJE\n",
			Output: EventDetail{
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Type:        "A",
				Multimedia: []MultimediaLink{
					{
						Data: &MultimediaLinkID{
							ID: "A",
						},
					},
					{
						Data: &MultimediaLinkID{
							ID: "B",
						},
					},
				},
			},
		},
		{ // 42
			Input: "0 EVEN\n1 TYPE A\n1 @B@ NOTE\n",
			Output: EventDetail{
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Type:        "A",
				Notes: []NoteStructure{
					{
						Data: &NoteID{
							ID: "B",
						},
					},
				},
			},
		},
		{ // 43
			Input: "0 EVEN\n1 TYPE A\n1 NOTE\n",
			Err:   ErrContext{"EventDetail", cNOTE, ErrContext{"NoteText", "line_value", ErrInvalidLength{"SubmitterText", "", 1, 248}}},
		},
		{ // 44
			Input: "0 EVEN\n1 TYPE A\n1 @B@ NOTE\n1 @C@ NOTE\n",
			Output: EventDetail{
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Type:        "A",
				Notes: []NoteStructure{
					{
						Data: &NoteID{
							ID: "B",
						},
					},
					{
						Data: &NoteID{
							ID: "C",
						},
					},
				},
			},
		},
		{ // 45
			Input: "0 EVEN\n1 TYPE A\n1 UNKNOWN\n",
			Err:   ErrContext{"EventDetail", "UNKNOWN", ErrUnknownTag},
		},
		{ // 46
			Input: "0 EVEN\n1 TYPE A\n1 _UNKNOWN\n",
			Output: EventDetail{
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Type:        "A",
			},
		},
		{ // 47
			Input:   "0 EVEN\n1 TYPE A\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: EventDetail{
				PhoneNumber: make([]PhoneNumber, 0, 3),
				Type:        "A",
			},
		},
	})
}

func TestLDSSpouseSealing(t *testing.T) {
	test(t, []testData[LDSSpouseSealing]{
		{ // 1
			Input:  "0 SLGS\n",
			Output: LDSSpouseSealing{},
		},
		{ // 2
			Input: "0 SLGS\n1 STAT DNS",
			Output: LDSSpouseSealing{
				LDSSpouseSealingDateStatus: cDNS,
			},
		},
		{ // 3
			Input: "0 SLGS\n1 STAT\n",
			Err:   ErrContext{"LDSSpouseSealing", cSTAT, ErrInvalidValue{"LDSSpouseSealingDateStatus", ""}},
		},
		{ // 4
			Input: "0 SLGS\n1 STAT DNS\n1 STAT CLEARED",
			Err:   ErrContext{"LDSSpouseSealing", cSTAT, ErrSingleMultiple},
		},
		{ // 5
			Input:   "0 SLGS\n1 STAT DNS\n1 STAT CLEARED",
			Options: []Option{AllowMoreThanAllowed},
			Output: LDSSpouseSealing{
				LDSSpouseSealingDateStatus: cDNS,
			},
		},
		{ // 6
			Input: "0 SLGS\n1 DATE AAAA",
			Output: LDSSpouseSealing{
				Date: "AAAA",
			},
		},
		{ // 7
			Input: "0 SLGS\n1 DATE\n",
			Err:   ErrContext{"LDSSpouseSealing", cDATE, ErrInvalidLength{"DateLDSOrd", "", 4, 35}},
		},
		{ // 8
			Input: "0 SLGS\n1 DATE AAAA\n1 DATE BBBB",
			Err:   ErrContext{"LDSSpouseSealing", cDATE, ErrSingleMultiple},
		},
		{ // 9
			Input:   "0 SLGS\n1 DATE AAAA\n1 DATE BBBB",
			Options: []Option{AllowMoreThanAllowed},
			Output: LDSSpouseSealing{
				Date: "AAAA",
			},
		},
		{ // 10
			Input: "0 SLGS\n1 TEMP AAAA",
			Output: LDSSpouseSealing{
				TempleCode: "AAAA",
			},
		},
		{ // 11
			Input: "0 SLGS\n1 TEMP\n",
			Err:   ErrContext{"LDSSpouseSealing", cTEMP, ErrInvalidLength{"TempleCode", "", 4, 5}},
		},
		{ // 12
			Input: "0 SLGS\n1 TEMP AAAA\n1 TEMP BBBB",
			Err:   ErrContext{"LDSSpouseSealing", cTEMP, ErrSingleMultiple},
		},
		{ // 13
			Input:   "0 SLGS\n1 TEMP AAAA\n1 TEMP BBBB",
			Options: []Option{AllowMoreThanAllowed},
			Output: LDSSpouseSealing{
				TempleCode: "AAAA",
			},
		},
	})
}
