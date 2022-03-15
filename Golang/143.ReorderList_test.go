package Golang

import (
	"Golang/common"
	"testing"
)

func TestReorderList(t *testing.T) {
	for _, test := range []struct {
		head []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 4, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3}},
	} {
		origin := test.head
		got := makeListNode(test.head)
		reorderList(got)
		gotArray := makeArrayFromListNode(got)
		if !common.SliceEqual(gotArray, test.want) {
			t.Errorf("FAIL::ReorderList(%v) = %v, want: %v", origin, gotArray, test.want)
		} else {
			t.Logf("SUCCESS::ReorderList(%v) = %v", origin, test.want)
		}
	}
}
