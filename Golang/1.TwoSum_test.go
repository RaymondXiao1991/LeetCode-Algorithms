package Golang

import (
	"Golang/common"
	"testing"
)

func TestTwoSum(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	} {
		got := twoSum(test.nums, test.target)
		if !common.SliceEqual(got, test.want) {
			t.Errorf("FAIL::TwoSum(%v, %v) = %v, want: %v", test.nums, test.target, got, test.want)
		} else {
			t.Logf("SUCCESS::TwoSum(%v, %v) = %v", test.nums, test.target, test.want)
		}
	}
}
