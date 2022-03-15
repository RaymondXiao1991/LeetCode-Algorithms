package Golang

import (
	"Golang/common"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	for _, test := range []struct {
		head []int
		n    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 1, []int{1}},
	} {
		got := removeNthFromEnd(makeListNode(test.head), test.n)
		gotListNode := makeArrayFromListNode(got)
		if !common.SliceEqual(gotListNode, test.want) {
			t.Errorf("FAIL::RemoveNthFromEnd(%v, %d) = %v, want: %v", test.head, gotListNode, test.n, test.want)
		} else {
			t.Logf("SUCCESS::RemoveNthFromEnd(%v, %d) = %v", test.head, test.n, test.want)
		}
	}
}

func TestRemoveNthFromEnd2(t *testing.T) {
	for _, test := range []struct {
		head []int
		n    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 1, []int{1}},
	} {
		got := removeNthFromEnd2(makeListNode(test.head), test.n)
		gotListNode := makeArrayFromListNode(got)
		if !common.SliceEqual(gotListNode, test.want) {
			t.Errorf("FAIL::RemoveNthFromEnd(%v, %d) = %v, want: %v", test.head, gotListNode, test.n, test.want)
		} else {
			t.Logf("SUCCESS::RemoveNthFromEnd(%v, %d) = %v", test.head, test.n, test.want)
		}
	}
}
