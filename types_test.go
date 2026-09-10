package gedcom

import (
	"errors"
	"reflect"
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
			t.Errorf("test %d: expecting error %v, got %v", n+1, test.Error, err)
		} else if !reflect.DeepEqual(s, &test.Result) {
			t.Errorf("test %d: expecting value %v, got %v", n+1, test.Result, *s)
		}
	}
}

func testSimpleType[T ~string, U pointerOf[T]](t *testing.T, min, max uint) {
	name := reflect.TypeOf(new(T)).Elem().Name()

	testType[T, U](t, []typeTests[T]{
		{ // 1
			Line:   l(""),
			Error:  ErrInvalidLength{name, "", min, max},
			Result: "",
		},
		{ // 2
			Line:    l(""),
			Options: []Option{AllowWrongLength},
			Result:  "",
		},
		{ // 3
			Line:    l(strings.Repeat("a", int(min)-1)),
			Options: []Option{AllowWrongLength},
			Result:  T(strings.Repeat("a", int(min)-1)),
		},
		{ // 4
			Line:   l(strings.Repeat("a", int(min))),
			Result: T(strings.Repeat("a", int(min))),
		},
		{ // 5
			Line:   l(strings.Repeat("a", int(max))),
			Result: T(strings.Repeat("a", int(max))),
		},
		{ // 6
			Line:   l(strings.Repeat("a", int(max)+1)),
			Error:  ErrInvalidLength{name, strings.Repeat("a", int(max)+1), min, max},
			Result: "",
		},
		{ // 7
			Line:    l(strings.Repeat("a", int(max)+1)),
			Options: []Option{AllowWrongLength},
			Result:  T(strings.Repeat("a", int(max)+1)),
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
	testType(t, []typeTests[AddressLine]{
		{ // 1
			Line:   l(""),
			Error:  ErrInvalidLength{"AddressLine", "", 1, 60},
			Result: AddressLine(""),
		},
		{ // 2
			Line:   l(strings.Repeat("a", 61)),
			Error:  ErrInvalidLength{"AddressLine", strings.Repeat("a", 61), 1, 60},
			Result: AddressLine(""),
		},
		{ // 3
			Line:   l("a"),
			Result: AddressLine("a"),
		},
		{ // 4
			Line:   l("a", lt("", cCONT)),
			Error:  ErrContext{"AddressLine", cCONT, ErrInvalidLength{"AddressLine", "", 1, 60}},
			Result: AddressLine("a\n"),
		},
		{ // 5
			Line:   l("a", lt(strings.Repeat("b", 60), cCONT)),
			Result: AddressLine("a\n" + strings.Repeat("b", 60)),
		},
		{ // 6
			Line:   l("a", lt(strings.Repeat("b", 61), cCONT)),
			Error:  ErrContext{"AddressLine", cCONT, ErrInvalidLength{"AddressLine", strings.Repeat("b", 61), 1, 60}},
			Result: AddressLine("a\n"),
		},
		{ // 7
			Line:    l("a", lt(strings.Repeat("b", 61), cCONT)),
			Options: []Option{AllowWrongLength},
			Result:  AddressLine("a\n" + strings.Repeat("b", 61)),
		},
		{ // 8
			Line:   l("a", lt("b", cCONT)),
			Result: AddressLine("a\nb"),
		},
		{ // 9
			Line:   l("a", lt("b", cCONT), lt("c", cCONT)),
			Result: AddressLine("a\nb\nc"),
		},
		{ // 10
			Line:   l("a", lt("", cCONC)),
			Error:  ErrContext{"AddressLine", cCONC, ErrInvalidLength{"AddressLine", "", 1, 60}},
			Result: AddressLine("a"),
		},
		{ // 11
			Line:   l("a", lt(strings.Repeat("b", 61), cCONC)),
			Error:  ErrContext{"AddressLine", cCONC, ErrInvalidLength{"AddressLine", strings.Repeat("b", 61), 1, 60}},
			Result: AddressLine("a"),
		},
		{ // 12
			Line:   l("a", lt("b", cCONC)),
			Result: AddressLine("ab"),
		},
	})
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
	testType(t, []typeTests[AdoptedBy]{
		{ // 1
			Line:   l(""),
			Error:  ErrInvalidValue{"AdoptedBy", ""},
			Result: "",
		},
		{ // 2
			Line:    l(""),
			Options: []Option{IgnoreInvalidValue},
			Result:  "",
		},
		{ // 3
			Line:   l("a"),
			Error:  ErrInvalidValue{"AdoptedBy", "a"},
			Result: "",
		},
		{ // 4
			Line:    l("a"),
			Options: []Option{IgnoreInvalidValue},
			Result:  "",
		},
		{ // 5
			Line:   l("HUSB"),
			Result: cHUSB,
		},
		{ // 6
			Line:   l("Wife"),
			Result: cWIFE,
		},
		{ // 7
			Line:   l("both"),
			Result: cBOTH,
		},
	})
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
	testType(t, []typeTests[AttributeType]{
		{ // 1
			Line:   l(""),
			Error:  ErrInvalidValue{"AttributeType", ""},
			Result: "",
		},
		{ // 2
			Line:    l(""),
			Options: []Option{IgnoreInvalidValue},
			Result:  "",
		},
		{ // 3
			Line:   l("a"),
			Error:  ErrInvalidValue{"AttributeType", "a"},
			Result: "",
		},
		{ // 4
			Line:    l("a"),
			Options: []Option{IgnoreInvalidValue},
			Result:  "",
		},
		{ // 5
			Line:   l("CAST"),
			Result: cCAST,
		},
		{ // 6
			Line:   l("educ"),
			Result: cEDUC,
		},
		{ // 7
			Line:   l("NatI"),
			Result: cNATI,
		},
		{ // 8
			Line:   l("oCcU"),
			Result: cOCCU,
		},
		{ // 9
			Line:   l("PrOp"),
			Result: cPROP,
		},
		{ // 10
			Line:   l("reLI"),
			Result: cRELI,
		},
		{ // 11
			Line:   l("REsi"),
			Result: cRESI,
		},
		{ // 12
			Line:   l("titl"),
			Result: cTITL,
		},
	})
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
	testType(t, []typeTests[CharacterSet]{
		{ // 1
			Line:   l(""),
			Error:  ErrInvalidValue{"CharacterSet", ""},
			Result: "",
		},
		{ // 2
			Line:    l(""),
			Options: []Option{IgnoreInvalidValue},
			Result:  "",
		},
		{ // 3
			Line:   l("a"),
			Error:  ErrInvalidValue{"CharacterSet", "a"},
			Result: "",
		},
		{ // 4
			Line:    l("a"),
			Options: []Option{IgnoreInvalidValue},
			Result:  "",
		},
		{ // 5
			Line:   l("ANSEL"),
			Result: cANSEL,
		},
		{ // 6
			Line:   l("Unicode"),
			Result: cUNICODE,
		},
		{ // 7
			Line:   l("ascii"),
			Result: cASCII,
		},
	})
}

func TestCopyrightGedcomFile(t *testing.T) {
	testSimpleType[CopyrightGedcomFile](t, 1, 90)
}

func TestCopyrightSourceData(t *testing.T) {
	testSimpleType[CopyrightSourceData](t, 1, 90)
}

func TestCountOfChildren(t *testing.T) {
	testType(t, []typeTests[CountOfChildren]{
		{ // 1
			Line:   l(""),
			Error:  ErrInvalidLength{"CountOfChildren", "", 1, 3},
			Result: 0,
		},
		{ // 2
			Line:    l(""),
			Options: []Option{AllowWrongLength, IgnoreInvalidValue},
			Result:  0,
		},
		{ // 3
			Line:   l("a"),
			Error:  ErrInvalidValue{"CountOfChildren", "a"},
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
			Error:  ErrInvalidValue{"CountOfChildren", "-1"},
			Result: 0,
		},
		{ // 7
			Line:   l("255"),
			Result: 255,
		},
		{ // 8
			Line:   l("256"),
			Error:  ErrInvalidValue{"CountOfChildren", "256"},
			Result: 0,
		},
		{ // 9
			Line:   l("1000"),
			Error:  ErrInvalidLength{"CountOfChildren", "1000", 1, 3},
			Result: 0,
		},
	})
}

func TestCountOfMarriages(t *testing.T) {
	testType(t, []typeTests[CountOfMarriages]{
		{ // 1
			Line:   l(""),
			Error:  ErrInvalidLength{"CountOfMarriages", "", 1, 3},
			Result: 0,
		},
		{ // 2
			Line:    l(""),
			Options: []Option{AllowWrongLength, IgnoreInvalidValue},
			Result:  0,
		},
		{ // 3
			Line:   l("a"),
			Error:  ErrInvalidValue{"CountOfMarriages", "a"},
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
			Error:  ErrInvalidValue{"CountOfMarriages", "-1"},
			Result: 0,
		},
		{ // 7
			Line:   l("255"),
			Result: 255,
		},
		{ // 8
			Line:   l("256"),
			Error:  ErrInvalidValue{"CountOfMarriages", "256"},
			Result: 0,
		},
		{ // 9
			Line:   l("1000"),
			Error:  ErrInvalidLength{"CountOfMarriages", "1000", 1, 3},
			Result: 0,
		},
	})
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
