package Golang

import "testing"

func TestSearch(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	} {
		got := search(test.nums, test.target)
		if got != test.want {
			t.Errorf("FAIL::Search(%v, %v) = %d, want: %d", test.nums, test.target, got, test.want)
		} else {
			t.Logf("SUCCESS::Search(%v, %v) = %d", test.nums, test.target, test.want)
		}
	}
}
