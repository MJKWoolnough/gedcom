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

func testSimpleType[T ~string, U pointerOf[T]](t *testing.T, name string, min, max uint) {
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
			Line:   l("a"),
			Result: "a",
		},
		{ // 4
			Line:   l(strings.Repeat("a", int(max))),
			Result: T(strings.Repeat("a", int(max))),
		},
		{ // 5
			Line:   l(strings.Repeat("a", int(max)+1)),
			Error:  ErrInvalidLength{name, strings.Repeat("a", int(max)+1), min, max},
			Result: "",
		},
		{ // 6
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
	testSimpleType[AddressCity](t, "AddressCity", 1, 60)
}

func TestAddressCountry(t *testing.T) {
	testSimpleType[AddressCountry](t, "AddressCountry", 1, 60)
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
	testSimpleType[AddressLine1](t, "AddressLine1", 1, 60)
}

func TestAddressLine2(t *testing.T) {
	testSimpleType[AddressLine2](t, "AddressLine2", 1, 60)
}

func TestAddressPostalCode(t *testing.T) {
	testSimpleType[AddressPostalCode](t, "AddressPostalCode", 1, 10)
}

func TestAddressState(t *testing.T) {
	testSimpleType[AddressState](t, "AddressState", 1, 60)
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
	testSimpleType[AgeAtEvent](t, "AgeAtEvent", 1, 12)
}

func TestAncestralFileNumber(t *testing.T) {
	testSimpleType[AncestralFileNumber](t, "AncestralFileNumber", 1, 12)
}
