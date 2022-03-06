package Golang

import (
	"Golang/common"
	"testing"
)

func TestSearchRange(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
	} {
		got := searchRange(test.nums, test.target)
		if !common.SliceEqual(got, test.want) {
			t.Errorf("FAIL::SearchRange(%v, %v) = %d, want: %d", test.nums, test.target, got, test.want)
		} else {
			t.Logf("SUCCESS::SearchRange(%v, %v) = %d", test.nums, test.target, test.want)
		}
	}
}

func TestSearchRangeByUseSort(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
	} {
		got := searchRangeByUseSort(test.nums, test.target)
		if !common.SliceEqual(got, test.want) {
			t.Errorf("FAIL::SearchRangeByUseSort(%v, %v) = %d, want: %d", test.nums, test.target, got, test.want)
		} else {
			t.Logf("SUCCESS::SearchRangeByUseSort(%v, %v) = %d", test.nums, test.target, test.want)
		}
	}
}

func TestSearchRangeByUseWG(t *testing.T) {
	for _, test := range []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
	} {
		got := searchRangeByUseWG(test.nums, test.target)
		if !common.SliceEqual(got, test.want) {
			t.Errorf("FAIL::SearchRangeByUseWG(%v, %v) = %d, want: %d", test.nums, test.target, got, test.want)
		} else {
			t.Logf("SUCCESS::SearchRangeByUseWG(%v, %v) = %d", test.nums, test.target, test.want)
		}
	}
}
