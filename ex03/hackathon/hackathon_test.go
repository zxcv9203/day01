package hackathon

import (
	"reflect"
	"testing"
)

func TestMatch(t *testing.T) {
	got := [][][]int{
		Match([]int{0, 1, 2}),
		Match([]int{2, 1, 3, 0, 5, 4, 0, 1}),
		Match([]int{2, 0, 3, 1}),
		Match([]int{0, 0, 0, 0, 0, 0}),
	}
	want := [][][]int{
		{{0}, {1}, {2}},
		{{0, 2, 3}, {1}, {4, 5}, {6, 7}},
		{{0, 1, 2, 3}},
		{{0}, {1, 2, 3, 4, 5}},
	}
	for i := range want {
		if !reflect.DeepEqual(got[i], want[i]) {
			t.Errorf("got : %#v want : %#v", got[i], want[i])
		}
	}

}
