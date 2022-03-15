package Golang

import (
	"Golang/common"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	for _, test := range []struct {
		head []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
	} {
		got := swapPairs(makeListNode(test.head))
		gotListNode := makeArrayFromListNode(got)
		if !common.SliceEqual(gotListNode, test.want) {
			t.Errorf("FAIL::SwapPairs(%v) = %v, want: %v", test.head, gotListNode, test.want)
		} else {
			t.Logf("SUCCESS::SwapPairs(%v) = %v", test.head, test.want)
		}
	}
}
