package Golang

import (
	"testing"
)

func TestFourSumCount(t *testing.T) {
	for _, test := range []struct {
		nums1 []int
		nums2 []int
		nums3 []int
		nums4 []int
		want  int
	}{
		{[]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}, 2},
		{[]int{0}, []int{0}, []int{0}, []int{0}, 1},
	} {
		got := fourSumCount(test.nums1, test.nums2, test.nums3, test.nums4)
		if got != test.want {
			t.Errorf("FAIL::FourSumCount(%v, %v, %v, %v) = %v, want: %v", test.nums1, test.nums2, test.nums3, test.nums4, got, test.want)
		} else {
			t.Logf("SUCCESS::FourSumCount(%v, %v, %v, %v) = %v", test.nums1, test.nums2, test.nums3, test.nums4, test.want)
		}
	}
}
