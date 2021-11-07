package find

import (
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	t.Run("normalCase", func(t *testing.T) {
		got := Find(100, 1, 2, 0, 3, 111, 314234, 41231, -333, 9393)
		want := []int{-333, 0, 1, 2, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%#v %#v", got, want)
		}
	})
	t.Run("no nums", func(t *testing.T) {
		got := Find(100)
		want := []int{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%#v %#v", got, want)
		}
	})
}
