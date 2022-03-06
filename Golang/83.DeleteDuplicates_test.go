package Golang

import (
	"Golang/common"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	for _, test := range []struct {
		Arr  []int
		want []int
	}{
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{1, 1, 2, 3, 3}, []int{1, 2, 3}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
	} {
		if common.SliceEqual(makeArrayFromListNode(deleteDuplicates(makeListNode(test.Arr))), test.want) {
			t.Logf("SUCCESS::DeleteDuplicates(%v) = %v", test.Arr, test.want)
		} else {
			t.Errorf("FAIL::DeleteDuplicates(%v) = %v, want:%v", test.Arr, makeArrayFromListNode(deleteDuplicates(makeListNode(test.Arr))), test.want)
		}
	}
}
