package Golang

import (
	"Golang/common"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	for _, test := range []struct {
		nums []int
		want []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	} {
		got := sortedSquares(test.nums)
		if !common.SliceEqual(got, test.want) {
			t.Errorf("FAIL::SortedSquares(%v) = %v, want:%v", test.nums, got, test.want)
		} else {
			t.Logf("SUCCESS::SortedSquares(%v) = %v", test.nums, test.want)
		}
	}
}
