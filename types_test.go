package gedcom

import (
	"errors"
	"reflect"
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
