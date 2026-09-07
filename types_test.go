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
	testType(t, []typeTests[AddressCity]{
		{ // 1
			Line:   l(""),
			Error:  ErrInvalidLength{"AddressCity", "", 1, 60},
			Result: AddressCity(""),
		},
		{ // 2
			Line:   l("a"),
			Result: AddressCity("a"),
		},
		{ // 3
			Line:   l(strings.Repeat("a", 60)),
			Result: AddressCity(strings.Repeat("a", 60)),
		},
		{ // 4
			Line:   l(strings.Repeat("a", 61)),
			Error:  ErrInvalidLength{"AddressCity", strings.Repeat("a", 61), 1, 60},
			Result: AddressCity(""),
		},
		{ // 5
			Line:    l(strings.Repeat("a", 61)),
			Options: []Option{AllowWrongLength},
			Result:  AddressCity(strings.Repeat("a", 61)),
		},
	})
}

func TestAddressCountry(t *testing.T) {
	testType(t, []typeTests[AddressCountry]{
		{ // 1
			Line:   l(""),
			Error:  ErrInvalidLength{"AddressCountry", "", 1, 60},
			Result: AddressCountry(""),
		},
		{ // 2
			Line:   l("a"),
			Result: AddressCountry("a"),
		},
		{ // 3
			Line:   l(strings.Repeat("a", 60)),
			Result: AddressCountry(strings.Repeat("a", 60)),
		},
		{ // 4
			Line:   l(strings.Repeat("a", 61)),
			Error:  ErrInvalidLength{"AddressCountry", strings.Repeat("a", 61), 1, 60},
			Result: AddressCountry(""),
		},
		{ // 5
			Line:    l(strings.Repeat("a", 61)),
			Options: []Option{AllowWrongLength},
			Result:  AddressCountry(strings.Repeat("a", 61)),
		},
	})
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
	testType(t, []typeTests[AddressLine1]{
		{ // 1
			Line:   l(""),
			Error:  ErrInvalidLength{"AddressLine1", "", 1, 60},
			Result: AddressLine1(""),
		},
		{ // 2
			Line:   l("a"),
			Result: AddressLine1("a"),
		},
		{ // 3
			Line:   l(strings.Repeat("a", 60)),
			Result: AddressLine1(strings.Repeat("a", 60)),
		},
		{ // 4
			Line:   l(strings.Repeat("a", 61)),
			Error:  ErrInvalidLength{"AddressLine1", strings.Repeat("a", 61), 1, 60},
			Result: AddressLine1(""),
		},
		{ // 5
			Line:    l(strings.Repeat("a", 61)),
			Options: []Option{AllowWrongLength},
			Result:  AddressLine1(strings.Repeat("a", 61)),
		},
	})
}

func TestAddressLine2(t *testing.T) {
	testType(t, []typeTests[AddressLine2]{
		{ // 1
			Line:   l(""),
			Error:  ErrInvalidLength{"AddressLine2", "", 1, 60},
			Result: AddressLine2(""),
		},
		{ // 2
			Line:   l("a"),
			Result: AddressLine2("a"),
		},
		{ // 3
			Line:   l(strings.Repeat("a", 60)),
			Result: AddressLine2(strings.Repeat("a", 60)),
		},
		{ // 4
			Line:   l(strings.Repeat("a", 61)),
			Error:  ErrInvalidLength{"AddressLine2", strings.Repeat("a", 61), 1, 60},
			Result: AddressLine2(""),
		},
		{ // 5
			Line:    l(strings.Repeat("a", 61)),
			Options: []Option{AllowWrongLength},
			Result:  AddressLine2(strings.Repeat("a", 61)),
		},
	})
}
