package Golang

import (
	"Golang/common"
	"testing"
)

func TestReverseList(t *testing.T) {
	for _, test := range []struct {
		array []int
		want  []int
	}{
		{[]int{1, 2, 2, 3, 2, 4}, []int{4, 2, 3, 2, 2, 1}},
		{[]int{1, 2, 3, 2, 4}, []int{4, 2, 3, 2, 1}},
		{[]int{1, 2, 3, 5, 4}, []int{4, 5, 3, 2, 1}},
		{[]int{1, 3, 4}, []int{4, 3, 1}},
		{[]int{2, 2, 3, 2, 4}, []int{4, 2, 3, 2, 2}},
	} {
		rl := reverseList(makeListNode(test.array))
		rlArray := makeArrayFromListNode(rl)
		if common.SliceEqual(rlArray, test.want) {
			t.Logf("SUCCESS::ReverseList(%v) = %v", test.array, test.want)
		} else {
			t.Errorf("FAIL::ReverseList(%v) = %v, want: %v", test.array, rlArray, test.want)
		}
	}
}

func TestReverseList2(t *testing.T) {
	for _, test := range []struct {
		array []int
		want  []int
	}{
		{[]int{1, 2, 2, 3, 2, 4}, []int{4, 2, 3, 2, 2, 1}},
		{[]int{1, 2, 3, 2, 4}, []int{4, 2, 3, 2, 1}},
		{[]int{1, 2, 3, 5, 4}, []int{4, 5, 3, 2, 1}},
		{[]int{1, 3, 4}, []int{4, 3, 1}},
		{[]int{2, 2, 3, 2, 4}, []int{4, 2, 3, 2, 2}},
	} {
		rl := reverseList2(makeListNode(test.array))
		rlArray := makeArrayFromListNode(rl)
		if common.SliceEqual(rlArray, test.want) {
			t.Logf("SUCCESS::ReverseList(%v) = %v", test.array, test.want)
		} else {
			t.Errorf("FAIL::ReverseList(%v) = %v, want: %v", test.array, rlArray, test.want)
		}
	}
}