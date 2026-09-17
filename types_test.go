package gedcom

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type typeI interface {
	parse(l *Line, o options) error
}

type typeTests[T any] struct {
	Line    Line
	Options []Option
	Result  T
	Error   error
}

type pointerOf[T any] interface {
	typeI
	*T
}

func testType[T any, U pointerOf[T]](t *testing.T, tests []typeTests[T]) {
	t.Helper()

	for n, test := range tests {
		var s U

		s = reflect.New(reflect.TypeOf(s).Elem()).Interface().(U)

		var opts options

		for _, opt := range test.Options {
			opt(&opts)
		}

		if err := s.parse(&test.Line, opts); !errors.Is(err, test.Error) {
			fmt.Printf("%#v\n%#v\n", err, test.Error)
			t.Errorf("test %d: expecting error %v, got %v", n+1, test.Error, err)
		} else if !reflect.DeepEqual(s, &test.Result) {
			t.Errorf("test %d: expecting value %v, got %v", n+1, test.Result, *s)
		}
	}
}

func testSimpleType[T ~string, U pointerOf[T]](t *testing.T, min, max uint) {
	name := reflect.TypeOf(new(T)).Elem().Name()

	var tests []typeTests[T]

	if min > 0 {
		tests = []typeTests[T]{
			{
				Line:   l(""),
				Error:  ErrInvalidLength{name, "", min, max},
				Result: "",
			},
			{
				Line:    l(""),
				Options: []Option{AllowWrongLength},
				Result:  "",
			},
			{
				Line:    l(strings.Repeat("a", int(min)-1)),
				Options: []Option{AllowWrongLength},
				Result:  T(strings.Repeat("a", int(min)-1)),
			},
		}
	}

	tests = append(tests,
		typeTests[T]{
			Line:   l(strings.Repeat("a", int(min))),
			Result: T(strings.Repeat("a", int(min))),
		},
		typeTests[T]{
			Line:   l(strings.Repeat("a", int(max))),
			Result: T(strings.Repeat("a", int(max))),
		},
		typeTests[T]{
			Line:   l(strings.Repeat("a", int(max)+1)),
			Error:  ErrInvalidLength{name, strings.Repeat("a", int(max)+1), min, max},
			Result: "",
		},
		typeTests[T]{
			Line:    l(strings.Repeat("a", int(max)+1)),
			Options: []Option{AllowWrongLength},
			Result:  T(strings.Repeat("a", int(max)+1)),
		},
	)

	testType[T, U](t, tests)
}

func testOptions[T ~string, U pointerOf[T]](t *testing.T, options ...T) {
	name := reflect.TypeOf(new(T)).Elem().Name()
	tests := []typeTests[T]{
		{
			Line:   l(""),
			Error:  ErrInvalidValue{name, ""},
			Result: "",
		},
		{
			Line:    l(""),
			Options: []Option{IgnoreInvalidValue},
			Result:  "",
		},
		{
			Line:   l("a"),
			Error:  ErrInvalidValue{name, "a"},
			Result: "",
		},
		{
			Line:    l("a"),
			Options: []Option{IgnoreInvalidValue},
			Result:  "",
		},
	}

	for _, opt := range options {
		for i := range 1 << len(opt) {
			var sb strings.Builder

			for n, c := range opt {
				if c < 'A' || c > 'Z' {
					sb.WriteByte(byte(c))

					continue
				}

				if i&(1<<n) == 0 {
					sb.WriteByte(byte(c) | 0x20)
				} else {
					sb.WriteByte(byte(c) & 0xDF)
				}
			}

			tests = append(tests, typeTests[T]{
				Line:   l(sb.String()),
				Result: opt,
			})
		}
	}

	testType[T, U](t, tests)
}

func testMultiLine[T ~string, U pointerOf[T]](t *testing.T, max uint) {
	name := reflect.TypeOf(new(T)).Elem().Name()

	testType[T, U](t, []typeTests[T]{
		{ // 1
			Line:   l(""),
			Error:  ErrInvalidLength{name, "", 1, max},
			Result: T(""),
		},
		{ // 2
			Line:   l(strings.Repeat("a", int(max)+1)),
			Error:  ErrInvalidLength{name, strings.Repeat("a", int(max)+1), 1, max},
			Result: T(""),
		},
		{ // 3
			Line:   l("a"),
			Result: T("a"),
		},
		{ // 4
			Line:   l("a", lt("", cCONT)),
			Error:  ErrContext{name, cCONT, ErrInvalidLength{name, "", 1, max}},
			Result: T("a\n"),
		},
		{ // 5
			Line:   l("a", lt(strings.Repeat("b", int(max)), cCONT)),
			Result: T("a\n" + strings.Repeat("b", int(max))),
		},
		{ // 6
			Line:   l("a", lt(strings.Repeat("b", int(max)+1), cCONT)),
			Error:  ErrContext{name, cCONT, ErrInvalidLength{name, strings.Repeat("b", int(max)+1), 1, max}},
			Result: T("a\n"),
		},
		{ // 7
			Line:    l("a", lt(strings.Repeat("b", int(max)+1), cCONT)),
			Options: []Option{AllowWrongLength},
			Result:  T("a\n" + strings.Repeat("b", int(max)+1)),
		},
		{ // 8
			Line:   l("a", lt("b", cCONT)),
			Result: T("a\nb"),
		},
		{ // 9
			Line:   l("a", lt("b", cCONT), lt("c", cCONT)),
			Result: T("a\nb\nc"),
		},
		{ // 10
			Line:   l("a", lt("", cCONC)),
			Error:  ErrContext{name, cCONC, ErrInvalidLength{name, "", 1, max}},
			Result: T("a"),
		},
		{ // 11
			Line:   l("a", lt(strings.Repeat("b", int(max)+1), cCONC)),
			Error:  ErrContext{name, cCONC, ErrInvalidLength{name, strings.Repeat("b", int(max)+1), 1, max}},
			Result: T("a"),
		},
		{ // 12
			Line:   l("a", lt("b", cCONC)),
			Result: T("ab"),
		},
	})
}

func testRange[T ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint, U pointerOf[T]](t *testing.T, max T) {
	name := reflect.TypeOf(new(T)).Elem().Name()
	maxLength := uint(len(strconv.FormatUint(uint64(max), 10)))

	var minErr, largeErr error

	if uint(len(strconv.FormatUint(uint64(max)+1, 10))) > maxLength {
		largeErr = ErrInvalidLength{name, strconv.FormatUint(uint64(max)+1, 10), 1, maxLength}
	} else {
		largeErr = ErrInvalidValue{name, strconv.FormatUint(uint64(max)+1, 10)}
	}

	if max < 10 {
		minErr = ErrInvalidLength{name, "-1", 1, maxLength}
	} else {
		minErr = ErrInvalidValue{name, "-1"}
	}

	testType[T, U](t, []typeTests[T]{
		{ // 1
			Line:   l(""),
			Error:  ErrInvalidLength{name, "", 1, maxLength},
			Result: 0,
		},
		{ // 2
			Line:    l(""),
			Options: []Option{AllowWrongLength, IgnoreInvalidValue},
			Result:  0,
		},
		{ // 3
			Line:   l("a"),
			Error:  ErrInvalidValue{name, "a"},
			Result: 0,
		},
		{ // 4
			Line:    l("a"),
			Options: []Option{IgnoreInvalidValue},
			Result:  0,
		},
		{ // 5
			Line:   l("1"),
			Result: 1,
		},
		{ // 6
			Line:   l("-1"),
			Error:  minErr,
			Result: 0,
		},
		{ // 7
			Line:   l(strconv.FormatUint(uint64(max), 10)),
			Result: max,
		},
		{ // 8
			Line:   l(strconv.FormatUint(uint64(max)+1, 10)),
			Error:  largeErr,
			Result: 0,
		},
		{ // 9
			Line:   l("1" + strings.Repeat("0", int(maxLength))),
			Error:  ErrInvalidLength{name, "1" + strings.Repeat("0", int(maxLength)), 1, maxLength},
			Result: 0,
		},
	})
}

func l(v string, subs ...Line) Line {
	return Line{
		line: line{
			value: v,
		},
		Sub: subs,
	}
}

func lt(v, t string, subs ...Line) Line {
	return Line{
		line: line{
			tag:   t,
			value: v,
		},
		Sub: subs,
	}
}

func TestAddressCity(t *testing.T) {
	testSimpleType[AddressCity](t, 1, 60)
}

func TestAddressCountry(t *testing.T) {
	testSimpleType[AddressCountry](t, 1, 60)
}

func TestAddressLine(t *testing.T) {
	testMultiLine[AddressLine](t, 60)
}

func TestAddressLine1(t *testing.T) {
	testSimpleType[AddressLine1](t, 1, 60)
}

func TestAddressLine2(t *testing.T) {
	testSimpleType[AddressLine2](t, 1, 60)
}

func TestAddressPostalCode(t *testing.T) {
	testSimpleType[AddressPostalCode](t, 1, 10)
}

func TestAddressState(t *testing.T) {
	testSimpleType[AddressState](t, 1, 60)
}

func TestAdoptedBy(t *testing.T) {
	testOptions[AdoptedBy](t, cHUSB, cWIFE, cBOTH)
}

func TestAgeAtEvent(t *testing.T) {
	testSimpleType[AgeAtEvent](t, 1, 12)
}

func TestAncestralFileNumber(t *testing.T) {
	testSimpleType[AncestralFileNumber](t, 1, 12)
}

func TestApprovedSystemID(t *testing.T) {
	testSimpleType[ApprovedSystemID](t, 1, 20)
}

func TestAttributeType(t *testing.T) {
	testOptions[AttributeType](t, cCAST, cEDUC, cNATI, cOCCU, cPROP, cRELI, cRESI, cTITL)
}

func TestAutomatedRecordID(t *testing.T) {
	testSimpleType[AutomatedRecordID](t, 1, 12)
}

func TestCasteName(t *testing.T) {
	testSimpleType[CasteName](t, 1, 90)
}

func TestCauseOfEvent(t *testing.T) {
	testSimpleType[CauseOfEvent](t, 1, 90)
}

func TestCertaintyAssessment(t *testing.T) {
	testType(t, []typeTests[CertaintyAssessment]{
		{ // 1
			Line:   l(""),
			Error:  ErrInvalidValue{"CertaintyAssessment", ""},
			Result: 0,
		},
		{ // 2
			Line:    l("CertaintyAssessment"),
			Options: []Option{IgnoreInvalidValue},
			Result:  0,
		},
		{ // 3
			Line:   l("a"),
			Error:  ErrInvalidValue{"CertaintyAssessment", "a"},
			Result: 0,
		},
		{ // 4
			Line:    l("a"),
			Options: []Option{IgnoreInvalidValue},
			Result:  0,
		},
		{ // 5
			Line:   l("0"),
			Result: 0,
		},
		{ // 6
			Line:   l("1"),
			Result: 1,
		},
		{ // 7
			Line:   l("2"),
			Result: 2,
		},
		{ // 8
			Line:   l("3"),
			Result: 3,
		},
		{ // 9
			Line:   l("4"),
			Error:  ErrInvalidValue{"CertaintyAssessment", "4"},
			Result: 0,
		},
	})
}

func TestChangeDate(t *testing.T) {
	testSimpleType[ChangeDate](t, 10, 11)
}

func TestCharacterSet(t *testing.T) {
	testOptions[CharacterSet](t, cANSEL, cUNICODE, cASCII)
}

func TestCopyrightGedcomFile(t *testing.T) {
	testSimpleType[CopyrightGedcomFile](t, 1, 90)
}

func TestCopyrightSourceData(t *testing.T) {
	testSimpleType[CopyrightSourceData](t, 1, 90)
}

func TestCountOfChildren(t *testing.T) {
	testRange[CountOfChildren](t, 255)
}

func TestCountOfMarriages(t *testing.T) {
	testRange[CountOfMarriages](t, 255)
}

func TestDate(t *testing.T) {
	testSimpleType[Date](t, 3, 35)
}

func TestDateApproximated(t *testing.T) {
	testSimpleType[DateApproximated](t, 4, 35)
}

func TestDateCalendar(t *testing.T) {
	testSimpleType[DateCalendar](t, 4, 35)
}

func TestDateCalendarEscape(t *testing.T) {
	testSimpleType[DateCalendarEscape](t, 4, 15)
}

func TestDateExact(t *testing.T) {
	testSimpleType[DateExact](t, 10, 11)
}

func TestDateFren(t *testing.T) {
	testSimpleType[DateFren](t, 4, 35)
}

func TestDateHebr(t *testing.T) {
	testSimpleType[DateHebr](t, 4, 35)
}

func TestDateJuln(t *testing.T) {
	testSimpleType[DateJuln](t, 4, 35)
}

func TestDateLDSOrd(t *testing.T) {
	testSimpleType[DateLDSOrd](t, 4, 35)
}

func TestDatePeriod(t *testing.T) {
	testSimpleType[DatePeriod](t, 7, 35)
}

func TestDatePhrase(t *testing.T) {
	testSimpleType[DatePhrase](t, 1, 35)
}

func TestDateRange(t *testing.T) {
	testSimpleType[DateRange](t, 8, 35)
}

func TestDateValue(t *testing.T) {
	testSimpleType[DateValue](t, 1, 35)
}

func TestDay(t *testing.T) {
	testRange[Day](t, 99)
}

func TestDescriptiveTitle(t *testing.T) {
	testSimpleType[DescriptiveTitle](t, 1, 248)
}

func TestDigit(t *testing.T) {
	testRange[Digit](t, 9)
}

func TestEncodedMultimediaLine(t *testing.T) {
	testSimpleType[EncodedMultimediaLine](t, 0, 87)
}

func TestEntryRecordingDate(t *testing.T) {
	testSimpleType[EntryRecordingDate](t, 1, 90)
}

func TestEventAttributeType(t *testing.T) {
	testSimpleType[EventAttributeType](t, 1, 15)
}

func TestEventDescriptor(t *testing.T) {
	testSimpleType[EventDescriptor](t, 1, 90)
}

func TestEventTypeCitedFrom(t *testing.T) {
	testSimpleType[EventTypeCitedFrom](t, 1, 15)
}

func TestEventTypeFamily(t *testing.T) {
	testOptions[EventTypeFamily](t, cANUL, cCENS, cDIV, cDIVF, cENGA, cMARR, cMARB, cMARC, cMARL, cMARS, cEVEN)
}

func TestEventTypeIndividual(t *testing.T) {
	testOptions[EventTypeIndividual](t, cADOP, cBIRT, cBAPM, cBARM, cBASM, cBLES, cBURI, cCENS, cCHR, cCHRA, cCONF, cCREM, cDEAT, cEMIG, cFCOM, cGRAD, cIMMI, cNATU, cORDN, cRETI, cPROB, cWILL, cEVEN)
}

func TestEventsRecorded(t *testing.T) {
	testSimpleType[EventsRecorded](t, 1, 90)
}

func TestFileName(t *testing.T) {
	testSimpleType[FileName](t, 1, 90)
}

func TestContentDescription(t *testing.T) {
	testMultiLine[ContentDescription](t, 248)
}

func TestForm(t *testing.T) {
	testSimpleType[Form](t, 14, 20)
}

func TestGenerationsOfAncestors(t *testing.T) {
	testRange[GenerationsOfAncestors](t, 9999)
}

func TestGenerationsOfDescendants(t *testing.T) {
	testRange[GenerationsOfDescendants](t, 9999)
}

func TestLanguageID(t *testing.T) {
	testSimpleType[LanguageID](t, 1, 15)
}

func TestLanguageOfText(t *testing.T) {
	testSimpleType[LanguageOfText](t, 1, 15)
}

func TestLanguagePreference(t *testing.T) {
	testSimpleType[LanguagePreference](t, 1, 90)
}

func TestLDSBaptismDateStatus(t *testing.T) {
	testOptions[LDSBaptismDateStatus](t, cCHILD, cCLEARED, cINFANT, cPRE1970, cQUALIFIED, cSTILLBORN, cSUBMITTED, cUNCLEARED)
}

func TestLDSChildSealingDateStatus(t *testing.T) {
	testOptions[LDSChildSealingDateStatus](t, cBIC, cCLEARED, cCOMPLETED, cDNS, cPRE1970, cQUALIFIED, cSTILLBORN, cSUBMITTED, cUNCLEARED)
}

func TestLDSEndowmentDateStatus(t *testing.T) {
	testOptions[LDSEndowmentDateStatus](t, cCHILD, cCLEARED, cCOMPLETED, cINFANT, cPRE1970, cQUALIFIED, cSTILLBORN, cSUBMITTED, cUNCLEARED)
}

func TestLDSSpouseSealingDateStatus(t *testing.T) {
	testOptions[LDSSpouseSealingDateStatus](t, cCANCELED, cCLEARED, cCOMPLETED, cDNS, cDNSCAN, cPRE1970, cQUALIFIED, cSUBMITTED, cUNCLEARED)
}

func TestMonth(t *testing.T) {
	testOptions[Month](t, cJAN, cFEB, cMAR, cAPR, cMAY, cJUN, cJUL, cAUG, cSEP, cOCT, cNOV, cDEC)
}

func TestMonthFren(t *testing.T) {
	testOptions[MonthFren](t, cVEND, cBRUM, cFRIM, cNIVO, cPLUV, cVENT, cGERM, cFLOR, cPRAI, cMESS, cTHER, cFRUC, cCOMP)
}

func TestMonthHebr(t *testing.T) {
	testOptions[MonthHebr](t, cTSH, cCSH, cKSL, cTVT, cSHV, cADR, cADS, cNSN, cIYR, cSVN, cTMZ, cAAV, cELL)
}

func TestMultimediaFileReference(t *testing.T) {
	testSimpleType[MultimediaFileReference](t, 1, 30)
}

func TestMultimediaFormat(t *testing.T) {
	testOptions[MultimediaFormat](t, cbmp, cgif, cjpeg, cole, cpcx, ctiff, cwav)
}

func TestNameOfBusiness(t *testing.T) {
	testSimpleType[NameOfBusiness](t, 1, 90)
}

func TestNameOfFamilyFile(t *testing.T) {
	testSimpleType[NameOfFamilyFile](t, 1, 20)
}

func TestNameOfProduct(t *testing.T) {
	testSimpleType[NameOfProduct](t, 1, 90)
}

func TestNameOfRepository(t *testing.T) {
	testSimpleType[NameOfRepository](t, 1, 90)
}

func TestNameOfSourceData(t *testing.T) {
	testSimpleType[NameOfSourceData](t, 1, 90)
}

func TestNamePersonal(t *testing.T) {
	testSimpleType[NamePersonal](t, 1, 120)
}

func TestNamePiece(t *testing.T) {
	testSimpleType[NamePiece](t, 1, 90)
}

func TestNamePieceGiven(t *testing.T) {
	testSimpleType[NamePieceGiven](t, 1, 120)
}

func TestNamePieceNickname(t *testing.T) {
	testSimpleType[NamePieceNickname](t, 1, 30)
}

func TestNamePiecePrefix(t *testing.T) {
	testSimpleType[NamePiecePrefix](t, 1, 30)
}

func TestNamePieceSuffix(t *testing.T) {
	testSimpleType[NamePieceSuffix](t, 1, 30)
}

func TestNamePieceSurname(t *testing.T) {
	testSimpleType[NamePieceSurname](t, 1, 120)
}

func TestNamePieceSurnamePrefix(t *testing.T) {
	testSimpleType[NamePieceSurnamePrefix](t, 1, 30)
}

func TestNationalIDNumber(t *testing.T) {
	testSimpleType[NationalIDNumber](t, 1, 30)
}

func TestNationalOrTribalOrigin(t *testing.T) {
	testSimpleType[NationalOrTribalOrigin](t, 1, 120)
}

func TestNewTag(t *testing.T) {
	testSimpleType[NewTag](t, 1, 15)
}

func TestNobilityTypeTitle(t *testing.T) {
	testSimpleType[NobilityTypeTitle](t, 1, 120)
}

func TestNumber(t *testing.T) {
	testRange[Number](t, 999999999)
}

func TestOccupation(t *testing.T) {
	testSimpleType[Occupation](t, 1, 90)
}

func TestOrdinanceProcessFlag(t *testing.T) {
	testOptions[OrdinanceProcessFlag](t, cyes, cno)
}

func TestPedigreeLinkageType(t *testing.T) {
	testOptions[PedigreeLinkageType](t, cadopted, cbirth, cfoster, csealing)
}

func TestPhoneNumber(t *testing.T) {
	testSimpleType[PhoneNumber](t, 1, 25)
}

func TestPhysicalDescription(t *testing.T) {
	testSimpleType[PhysicalDescription](t, 1, 248)
}

func TestPlaceHierarchy(t *testing.T) {
	testSimpleType[PlaceHierarchy](t, 1, 120)
}

func TestPlaceLivingOrdinance(t *testing.T) {
	testSimpleType[PlaceLivingOrdinance](t, 1, 120)
}

func TestPlaceValue(t *testing.T) {
	testSimpleType[PlaceValue](t, 1, 120)
}

func TestPossessions(t *testing.T) {
	testSimpleType[Possessions](t, 1, 248)
}

func TestPublicationDate(t *testing.T) {
	testSimpleType[PublicationDate](t, 10, 11)
}

func TestReceivingSystemName(t *testing.T) {
	testSimpleType[ReceivingSystemName](t, 1, 20)
}

func TestRecordIdentifier(t *testing.T) {
	testSimpleType[RecordIdentifier](t, 1, 18)
}

func TestRecordType(t *testing.T) {
	testOptions[RecordType](t, cFAM, cINDI, cNOTE, cOBJE, cREPO, cSOUR, cSUBM, cSUBN)
}
