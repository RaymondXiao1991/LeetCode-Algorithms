package Golang

import (
	"Golang/common"
	"testing"
)

func TestIntersection(t *testing.T) {
	for _, test := range []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
	} {
		got := intersection(test.nums1, test.nums2)
		if !common.SliceEqual(got, test.want) {
			t.Errorf("FAIL::Intersection(%v, %v) = %v, want:%v", test.nums1, test.nums2, got, test.want)
		} else {
			t.Logf("SUCCESS::Intersection(%v, %v) = %v", test.nums1, test.nums2, test.want)
		}
	}
}
