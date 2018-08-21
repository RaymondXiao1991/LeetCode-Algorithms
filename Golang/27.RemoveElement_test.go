package Golang_test

import (
	"Golang"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	for _, test := range []struct {
		arr  []int
		elem int
		want int
	}{
		{[]int{1, 2, 2, 3, 2, 4}, 2, 3},
		{[]int{1, 2, 3, 2, 4}, 2, 3},
		{[]int{1, 2, 3, 2, 4}, 2, 3},
		{[]int{1, 3, 4}, 2, 3},
		{[]int{2, 2, 3, 2, 4}, 2, 2},
	} {
		if Golang.RemoveElement(test.arr, test.elem) != test.want {
			t.Errorf("RemoveElement(%v, %v) = %d", test.arr, test.elem, test.want)
		} else {
			t.Logf("Equal(%v, %v) = %d", test.arr, test.elem, test.want)
		}
	}
}
