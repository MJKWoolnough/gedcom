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
			Err:   ErrContext{"Header", "Source", ErrRequiredMissing},
		},
		{ // 2
			Input:   "0 HEADER\n",
			Options: []Option{AllowMissingRequired},
		},
		{ // 3
			Input: "0 HEADER\n1 SOUR id",
			Err:   ErrContext{"Header", "Submitter", ErrRequiredMissing},
		},
		{ // 4
			Input: "0 HEADER\n1 SOUR id\n1 SUBM submitter",
			Err:   ErrContext{"Header", "Version", ErrRequiredMissing},
		},
		{ // 5
			Input: "0 HEADER\n1 SOUR id\n1 SUBM submitter\n1 GEDC\n2 VERS 5.5\n2 FORM LINEAGE-LINKED",
			Err:   ErrContext{"Header", "CharacterSet", ErrRequiredMissing},
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
			Err:   ErrContext{"Header", cGEDC, ErrContext{"Version", "VersionNumber", ErrRequiredMissing}},
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
			Err:   ErrContext{"Header", cPLAC, ErrContext{"HeaderPlace", "PlaceHierarchy", ErrRequiredMissing}},
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
