package Golang

import (
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	for _, test := range []struct {
		target int
		nums   []int
		want   int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{4, []int{1, 4, 4}, 1},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
	} {
		got := minSubArrayLen(test.target, test.nums)
		if got != test.want {
			t.Errorf("FAIL::MinSubArrayLen(%v, %v) = %v, want:%v", test.target, test.nums, got, test.want)
		} else {
			t.Logf("SUCCESS::MinSubArrayLen(%v, %v) = %v", test.target, test.nums, test.want)
		}
	}
}
