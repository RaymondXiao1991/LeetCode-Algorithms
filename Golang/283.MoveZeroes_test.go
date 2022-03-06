package Golang

import (
	"Golang/common"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	for _, test := range []struct {
		nums []int
		want []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0, 0, 0, 3, 12}, []int{3, 12, 0, 0, 0}},
		{[]int{0}, []int{0}},
	} {
		moveZeroes(test.nums)
		if !common.SliceEqual(test.nums, test.want) {
			t.Errorf("FAIL::MoveZeroes(%v) = %v, want:%v", test.nums, test.nums, test.want)
		} else {
			t.Logf("SUCCESS::MoveZeroes(%v) = %v", test.nums, test.want)
		}
	}
}
