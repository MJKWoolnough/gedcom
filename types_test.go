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
	Options options
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

		if err := s.parse(&test.Line, test.Options); !errors.Is(err, test.Error) {
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
			Line:   l(strings.Repeat("a", 61)),
			Error:  ErrInvalidLength{"AddressCity", strings.Repeat("a", 61), 1, 60},
			Result: AddressCity(""),
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
			Line:   l(strings.Repeat("a", 61)),
			Error:  ErrInvalidLength{"AddressCountry", strings.Repeat("a", 61), 1, 60},
			Result: AddressCountry(""),
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
			Line:   l("a", lt(strings.Repeat("b", 61), cCONT)),
			Error:  ErrContext{"AddressLine", cCONT, ErrInvalidLength{"AddressLine", strings.Repeat("b", 61), 1, 60}},
			Result: AddressLine("a\n"),
		},
		{ // 6
			Line:   l("a", lt("b", cCONT)),
			Result: AddressLine("a\nb"),
		},
		{ // 7
			Line:   l("a", lt("b", cCONT), lt("c", cCONT)),
			Result: AddressLine("a\nb\nc"),
		},
		{ // 8
			Line:   l("a", lt("", cCONC)),
			Error:  ErrContext{"AddressLine", cCONC, ErrInvalidLength{"AddressLine", "", 1, 60}},
			Result: AddressLine("a"),
		},
		{ // 9
			Line:   l("a", lt(strings.Repeat("b", 61), cCONC)),
			Error:  ErrContext{"AddressLine", cCONC, ErrInvalidLength{"AddressLine", strings.Repeat("b", 61), 1, 60}},
			Result: AddressLine("a"),
		},
		{ // 10
			Line:   l("a", lt("b", cCONC)),
			Result: AddressLine("ab"),
		},
	})
}
