package Golang

import "testing"

func TestSearchInsert(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1}, 0, 0},
	} {
		got := searchInsert(test.nums, test.target)
		if got != test.want {
			t.Errorf("FAIL::SearchInsert(%v, %v) = %d, want: %d", test.nums, test.target, got, test.want)
		} else {
			t.Logf("SUCCESS::SearchInsert(%v, %v) = %d", test.nums, test.target, test.want)
		}
	}
}
