package Golang

import (
	"Golang/common"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	for _, test := range []struct {
		arr  []int
		val  int
		want []int
	}{
		{[]int{1, 2, 6, 3, 4, 5, 6}, 6, []int{1, 2, 3, 4, 5}},
		{[]int{}, 1, []int{}},
		{[]int{7, 7, 7, 7}, 7, []int{}},
	} {
		got := removeElements(makeListNode(test.arr), test.val)
		if !common.SliceEqual(makeArrayFromListNode(got), test.want) {
			t.Errorf("FAIL::RemoveElements(%v, %d) = %v, want:%v", test.arr, test.val, got, test.want)
		} else {
			t.Logf("SUCCESS::RemoveElements(%v, %d) = %v", test.arr, test.val, test.want)
		}
	}
}
