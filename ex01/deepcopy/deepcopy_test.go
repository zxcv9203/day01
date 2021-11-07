package deepcopy

import (
	"reflect"
	"testing"
)

func TestDeepCopy(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Run("normal", func(t *testing.T) {
		got, err := DeepCopy(nums, 1, 4)
		if err != nil {
			t.Errorf("Error : %s\n", err)
		}
		for i, value := range got {
			got[i] += value
		}
		want := []int{2, 4, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %#v want : %#v", got, want)
		}
	})
	t.Run("IndexingError", func(t *testing.T) {
		got, err := DeepCopy(nums, 4, 1)
		if got != nil && err != nil {
			t.Errorf("There was no error even though an incorrect value came in\n")
		}
	})
	t.Run("outOfRangeError", func(t *testing.T) {
		got, err := DeepCopy(nums, -3, 1)
		if got != nil && err != nil {
			t.Errorf("There was no error even though an incorrect value came in\n")
		}
	})
}
