package Golang

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	for _, test := range []struct {
		Arr  []int
		Elem int
		want int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	} {
		got := removeElement(test.Arr, test.Elem)
		if got != test.want {
			t.Errorf("FAIL::RemoveElement(%v, %v) = %d, want:%d", test.Arr, test.Elem, got, test.want)
		} else {
			t.Logf("SUCCESS::RemoveElement(%v, %v) = %d", test.Arr, test.Elem, test.want)
		}
	}
}
