package common

import "testing"

func TestQuickSort(t *testing.T) {
	for _, test := range []struct {
		nums []int
		want []int
	}{
		{[]int{2, 5, 3, 7, 4, 5, 8, 1, 4, 0}, []int{0, 1, 2, 3, 4, 4, 5, 5, 7, 8}},
		{[]int{7, 3, 5, 2, 8, 4, 5, 0, 4, 1}, []int{0, 1, 2, 3, 4, 4, 5, 5, 7, 8}},
	} {
		got := QuickSort(test.nums)
		if !SliceEqual(got, test.want) {
			t.Errorf("FAIL::QuickSort(%v) = %v, want: %v", test.nums, got, test.want)
		} else {
			t.Logf("SUCCESS::QuickSort(%v) = %v", test.nums, test.want)
		}
	}
}

func TestQuickSort2(t *testing.T) {
	for _, test := range []struct {
		nums []int
		want []int
	}{
		{[]int{2, 5, 3, 7, 4, 5, 8, 1, 4, 0}, []int{0, 1, 2, 3, 4, 4, 5, 5, 7, 8}},
		{[]int{7, 3, 5, 2, 8, 4, 5, 0, 4, 1}, []int{0, 1, 2, 3, 4, 4, 5, 5, 7, 8}},
	} {
		got := QuickSort2(test.nums)
		if !SliceEqual(got, test.want) {
			t.Errorf("FAIL::QuickSort(%v) = %v, want: %v", test.nums, got, test.want)
		} else {
			t.Logf("SUCCESS::QuickSort(%v) = %v", test.nums, test.want)
		}
	}
}
