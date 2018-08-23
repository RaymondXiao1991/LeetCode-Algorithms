package Golang_test

import (
	"Golang"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	for _, test := range []struct {
		Arr  []int
		Elem int
		want int
	}{
		{[]int{1, 2, 2, 3, 2, 4}, 2, 3},
		{[]int{1, 2, 3, 2, 4}, 2, 3},
		{[]int{1, 2, 3, 2, 4}, 2, 3},
		{[]int{1, 3, 4}, 2, 3},
		{[]int{2, 2, 3, 2, 4}, 2, 2},
	} {
		if Golang.RemoveElement(test.Arr, test.Elem) != test.want {
			t.Errorf("RemoveElement(%v, %v) = %d", test.Arr, test.Elem, test.want)
		} else {
			t.Logf("RemoveElement(%v, %v) = %d", test.Arr, test.Elem, test.want)
		}
	}
}
