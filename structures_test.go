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
			Input: "0 HEADER\n",
			Err:   ErrContext{"Header", cSUBM, ErrRequiredMissing},
		},
		{ // 2
			Input:   "0 HEADER\n",
			Options: []Option{AllowMissingRequired},
		},
		{ // 3
			Input: "0 HEADER\n1 SOUR id",
			Err:   ErrContext{"Header", cSUBM, ErrRequiredMissing},
		},
		{ // 4
			Input: "0 HEADER\n1 SOUR id\n1 SUBM submitter",
			Err:   ErrContext{"Header", cGEDC, ErrRequiredMissing},
		},
		{ // 5
			Input: "0 HEADER\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED",
			Err:   ErrContext{"Header", cCHAR, ErrRequiredMissing},
		},
		{ // 6
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
		{ // 7
			Input: "0 HEADER\n1 SOUR\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cSOUR, ErrContext{"HeaderSource", "line_value", ErrInvalidLength{"ApprovedSystemID", "", 1, 20}}},
		},
		{ // 8
			Input: "0 HEADER\n1 SOUR id\n1 SOUR other\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cSOUR, ErrSingleMultiple},
		},
		{ // 9
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
		{ // 10
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
		{ // 11
			Input: "0 HEADER\n1 SOUR id\n1 DEST\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cDEST, ErrInvalidLength{"ReceivingSystemName", "", 1, 20}},
		},
		{ // 12
			Input: "0 HEADER\n1 SOUR id\n1 DEST destination\n1 DEST destination2\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cDEST, ErrSingleMultiple},
		},
		{ // 13
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
		{ // 14
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
		{ // 15
			Input: "0 HEADER\n1 SOUR id\n1 DATE\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cDATE, ErrContext{"TransmissionDateTime", "line_value", ErrInvalidLength{"TransmissionDate", "", 10, 11}}},
		},
		{ // 16
			Input: "0 HEADER\n1 SOUR id\n1 DATE 2006-05-04\n1 DATE 2007-06-05\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cDATE, ErrSingleMultiple},
		},
		{ // 17
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
		{ // 18
			Input: "0 HEADER\n1 SOUR id\n1 DEST destination\n1 SUBM\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cSUBM, ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 19
			Input: "0 HEADER\n1 SOUR id\n1 DEST destination\n1 SUBM submitter\n1 SUBM submitter2\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cSUBM, ErrSingleMultiple},
		},
		{ // 20
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
		{ // 21
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
		{ // 22
			Input: "0 HEADER\n1 SUMB\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cSUMB, ErrInvalidLength{"Xref", "", 1, 22}},
		},
		{ // 23
			Input: "0 HEADER\n1 SUMB xid1\n1 SUMB xid2\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cSUMB, ErrSingleMultiple},
		},
		{ // 24
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
		{ // 25
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
		{ // 26
			Input: "0 HEADER\n1 FILE\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cFILE, ErrInvalidLength{"FileName", "", 1, 90}},
		},
		{ // 27
			Input: "0 HEADER\n1 FILE filename1\n1 FILE filename2\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cFILE, ErrSingleMultiple},
		},
		{ // 28
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
		{ // 29
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
		{ // 30
			Input: "0 HEADER\n1 COPR\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cCOPR, ErrInvalidLength{"CopyrightGedcomFile", "", 1, 90}},
		},
		{ // 31
			Input: "0 HEADER\n1 COPR copyright1\n1 COPR copyright2\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cCOPR, ErrSingleMultiple},
		},
		{ // 32
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
		{ // 33
			Input: "0 HEADER\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cGEDC, ErrContext{"Version", cVERS, ErrRequiredMissing}},
		},
		{ // 34
			Input: "0 HEADER\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 GEDC\n2 VERS 5.6\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cGEDC, ErrSingleMultiple},
		},
		{ // 35
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
		{ // 36
			Input: "0 HEADER\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR utf-8",
			Err:   ErrContext{"Header", cCHAR, ErrContext{"CharacterSetStructure", "line_value", ErrInvalidValue{"CharacterSet", "utf-8"}}},
		},
		{ // 37
			Input: "0 HEADER\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL\n1 CHAR ASCII",
			Err:   ErrContext{"Header", cCHAR, ErrSingleMultiple},
		},
		{ // 38
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
		{ // 39
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
		{ // 40
			Input: "0 HEADER\n1 LANG\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cLANG, ErrInvalidLength{"LanguageOfText", "", 1, 15}},
		},
		{ // 41
			Input: "0 HEADER\n1 LANG eng\n1 LANG fre\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cLANG, ErrSingleMultiple},
		},
		{ // 42
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
		{ // 43
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
		{ // 44
			Input: "0 HEADER\n1 PLAC\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cPLAC, ErrContext{"HeaderPlace", cFORM, ErrRequiredMissing}},
		},
		{ // 45
			Input: "0 HEADER\n1 PLAC\n2 FORM place1\n1 PLAC\n2 FORM place2\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cPLAC, ErrSingleMultiple},
		},
		{ // 46
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
		{ // 47
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
		{ // 48
			Input: "0 HEADER\n1 NOTE\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cNOTE, ErrInvalidLength{"ContentDescription", "", 1, 248}},
		},
		{ // 49
			Input: "0 HEADER\n1 NOTE note1\n1 NOTE note2\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", cNOTE, ErrSingleMultiple},
		},
		{ // 50
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
		{ // 51
			Input: "0 HEADER\n1 UNKNOWN\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED\n1 CHAR ANSEL",
			Err:   ErrContext{"Header", "UNKNOWN", ErrUnknownTag},
		},
		{ // 52
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
		{ // 53
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
				t.Errorf("test %d: expecting %#v, got %#v", n+1, test.Output, s)
			}
		}
	}
}

func TestHeaderSource(t *testing.T) {
	for n, test := range [...]struct {
		Input   string
		Options []Option
		Output  HeaderSource
		Err     error
	}{
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
	} {
		var s HeaderSource

		r := NewReader(strings.NewReader(test.Input), test.Options...)

		if lines, err := readLines(r); err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
		} else {
			l := parseLines(lines)

			if err = s.parse(&l, r.options); !errors.Is(err, test.Err) {
				t.Errorf("test %d: expecting error %v, got %v", n+1, test.Err, err)
			} else if test.Err == nil && !reflect.DeepEqual(test.Output, s) {
				t.Errorf("test %d: expecting %#v, got %#v", n+1, test.Output, s)
			}
		}
	}
}

func TestTransmissionDateTime(t *testing.T) {
	for n, test := range [...]struct {
		Input   string
		Options []Option
		Output  TransmissionDateTime
		Err     error
	}{
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
	} {
		var s TransmissionDateTime

		r := NewReader(strings.NewReader(test.Input), test.Options...)

		if lines, err := readLines(r); err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
		} else {
			l := parseLines(lines)

			if err = s.parse(&l, r.options); !errors.Is(err, test.Err) {
				t.Errorf("test %d: expecting error %v, got %v", n+1, test.Err, err)
			} else if test.Err == nil && !reflect.DeepEqual(test.Output, s) {
				t.Errorf("test %d: expecting %#v, got %#v", n+1, test.Output, s)
			}
		}
	}
}

func TestHeaderBusiness(t *testing.T) {
	for n, test := range [...]struct {
		Input   string
		Options []Option
		Output  HeaderBusiness
		Err     error
	}{
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
	} {
		var s HeaderBusiness

		r := NewReader(strings.NewReader(test.Input), test.Options...)

		if lines, err := readLines(r); err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
		} else {
			l := parseLines(lines)

			if err = s.parse(&l, r.options); !errors.Is(err, test.Err) {
				t.Errorf("test %d: expecting error %v, got %v", n+1, test.Err, err)
			} else if test.Err == nil && !reflect.DeepEqual(test.Output, s) {
				t.Errorf("test %d: expecting %#v, got %#v", n+1, test.Output, s)
			}
		}
	}
}

func TestHeaderDataSource(t *testing.T) {
	for n, test := range [...]struct {
		Input   string
		Options []Option
		Output  HeaderDataSource
		Err     error
	}{
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
	} {
		var s HeaderDataSource

		r := NewReader(strings.NewReader(test.Input), test.Options...)

		if lines, err := readLines(r); err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
		} else {
			l := parseLines(lines)

			if err = s.parse(&l, r.options); !errors.Is(err, test.Err) {
				t.Errorf("test %d: expecting error %v, got %v", n+1, test.Err, err)
			} else if test.Err == nil && !reflect.DeepEqual(test.Output, s) {
				t.Errorf("test %d: expecting %#v, got %#v", n+1, test.Output, s)
			}
		}
	}
}

func TestVersion(t *testing.T) {
	for n, test := range [...]struct {
		Input   string
		Options []Option
		Output  Version
		Err     error
	}{
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
	} {
		var s Version

		r := NewReader(strings.NewReader(test.Input), test.Options...)

		if lines, err := readLines(r); err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
		} else {
			l := parseLines(lines)

			if err = s.parse(&l, r.options); !errors.Is(err, test.Err) {
				t.Errorf("test %d: expecting error %v, got %v", n+1, test.Err, err)
			} else if test.Err == nil && !reflect.DeepEqual(test.Output, s) {
				t.Errorf("test %d: expecting %#v, got %#v", n+1, test.Output, s)
			}
		}
	}
}

func TestCharacterSetStructure(t *testing.T) {
	for n, test := range [...]struct {
		Input   string
		Options []Option
		Output  CharacterSetStructure
		Err     error
	}{
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
	} {
		var s CharacterSetStructure

		r := NewReader(strings.NewReader(test.Input), test.Options...)

		if lines, err := readLines(r); err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
		} else {
			l := parseLines(lines)

			if err = s.parse(&l, r.options); !errors.Is(err, test.Err) {
				t.Errorf("test %d: expecting error %v, got %v", n+1, test.Err, err)
			} else if test.Err == nil && !reflect.DeepEqual(test.Output, s) {
				t.Errorf("test %d: expecting %#v, got %#v", n+1, test.Output, s)
			}
		}
	}
}

func TestHeaderPlace(t *testing.T) {
	for n, test := range [...]struct {
		Input   string
		Options []Option
		Output  HeaderPlace
		Err     error
	}{
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
	} {
		var s HeaderPlace

		r := NewReader(strings.NewReader(test.Input), test.Options...)

		if lines, err := readLines(r); err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
		} else {
			l := parseLines(lines)

			if err = s.parse(&l, r.options); !errors.Is(err, test.Err) {
				t.Errorf("test %d: expecting error %v, got %v", n+1, test.Err, err)
			} else if test.Err == nil && !reflect.DeepEqual(test.Output, s) {
				t.Errorf("test %d: expecting %#v, got %#v", n+1, test.Output, s)
			}
		}
	}
}

func TestFamily(t *testing.T) {
	for n, test := range [...]struct {
		Input   string
		Options []Option
		Output  Family
		Err     error
	}{
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
	} {
		var s Family

		r := NewReader(strings.NewReader(test.Input), test.Options...)

		if lines, err := readLines(r); err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
		} else {
			l := parseLines(lines)

			if err = s.parse(&l, r.options); !errors.Is(err, test.Err) {
				t.Errorf("test %d: expecting error %v, got %v", n+1, test.Err, err)
			} else if test.Err == nil && !reflect.DeepEqual(test.Output, s) {
				t.Errorf("test %d: expecting %#v, got %#v", n+1, test.Output, s)
			}
		}
	}
}

func TestVerifiedFamilyEventDetail(t *testing.T) {
	for n, test := range [...]struct {
		Input   string
		Options []Option
		Output  VerifiedFamilyEventDetail
		Err     error
	}{
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
	} {
		var s VerifiedFamilyEventDetail

		r := NewReader(strings.NewReader(test.Input), test.Options...)

		if lines, err := readLines(r); err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
		} else {
			l := parseLines(lines)

			if err = s.parse(&l, r.options); !errors.Is(err, test.Err) {
				t.Errorf("test %d: expecting error %v, got %v", n+1, test.Err, err)
			} else if test.Err == nil && !reflect.DeepEqual(test.Output, s) {
				t.Errorf("test %d: expecting %#v, got %#v", n+1, test.Output, s)
			}
		}
	}
}

func TestFamilyEventDetail(t *testing.T) {
	for n, test := range [...]struct {
		Input   string
		Options []Option
		Output  FamilyEventDetail
		Err     error
	}{
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
	} {
		var s FamilyEventDetail

		r := NewReader(strings.NewReader(test.Input), test.Options...)

		if lines, err := readLines(r); err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
		} else {
			l := parseLines(lines)

			if err = s.parse(&l, r.options); !errors.Is(err, test.Err) {
				t.Errorf("test %d: expecting error %v, got %v", n+1, test.Err, err)
			} else if test.Err == nil && !reflect.DeepEqual(test.Output, s) {
				t.Errorf("test %d: expecting %#v, got %#v", n+1, test.Output, s)
			}
		}
	}
}

func TestAgeStructure(t *testing.T) {
	for n, test := range [...]struct {
		Input   string
		Options []Option
		Output  AgeStructure
		Err     error
	}{
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
			Input: "0 HUSB\n1 AGE 20\n1 AGE 21",
			Err:   ErrContext{"AgeStructure", cAGE, ErrSingleMultiple},
		},
		{ // 4
			Input:   "0 HUSB\n1 AGE 20\n1 AGE 21",
			Options: []Option{AllowMoreThanAllowed},
			Output: AgeStructure{
				Age: "20",
			},
		},
		{ // 5
			Input: "0 HUSB\n1 AGE 20\n1 UNKNOWN\n",
			Err:   ErrContext{"AgeStructure", "UNKNOWN", ErrUnknownTag},
		},
		{ // 6
			Input: "0 HUSB\n1 AGE 20\n1 _UNKNOWN\n",
			Output: AgeStructure{
				Age: "20",
			},
		},
		{ // 7
			Input:   "0 HUSB\n1 AGE 20\n1 UNKNOWN\n",
			Options: []Option{AllowUnknownTags},
			Output: AgeStructure{
				Age: "20",
			},
		},
	} {
		var s AgeStructure

		r := NewReader(strings.NewReader(test.Input), test.Options...)

		if lines, err := readLines(r); err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
		} else {
			l := parseLines(lines)

			if err = s.parse(&l, r.options); !errors.Is(err, test.Err) {
				t.Errorf("test %d: expecting error %v, got %v", n+1, test.Err, err)
			} else if test.Err == nil && !reflect.DeepEqual(test.Output, s) {
				t.Errorf("test %d: expecting %#v, got %#v", n+1, test.Output, s)
			}
		}
	}
}

func TestIndividual(t *testing.T) {
	for n, test := range [...]struct {
		Input   string
		Options []Option
		Output  Individual
		Err     error
	}{
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
	} {
		var s Individual

		r := NewReader(strings.NewReader(test.Input), test.Options...)

		if lines, err := readLines(r); err != nil {
			t.Errorf("test %d: unexpected error: %s", n+1, err)
		} else {
			l := parseLines(lines)

			if err = s.parse(&l, r.options); !errors.Is(err, test.Err) {
				t.Errorf("test %d: expecting error %v, got %v", n+1, test.Err, err)
			} else if test.Err == nil && !reflect.DeepEqual(test.Output, s) {
				t.Errorf("test %d: expecting %#v, got %#v", n+1, test.Output, s)
			}
		}
	}
}
