package Golang

import (
	"Golang/common"
	"testing"
)

func TestSearchInArray(t *testing.T) {
	for _, test := range []struct {
		list    []int
		start   int
		end     int
		element int
		want    int
	}{
		{[]int{3, 3, 5, 6, 1, 2, 7, 8, 0, 1}, 4, 9, 1, 2},
	} {
		got := searchInArray(test.list, test.start, test.end, test.element)
		if got != test.want {
			t.Errorf("FAIL::SearchInArray(%v, %v, %v, %v) = %d, want: %d", test.list, test.start, test.end, test.element, got, test.want)
		} else {
			t.Logf("SUCCESS::SearchInArray(%v, %v, %v, %v) = %d", test.list, test.start, test.end, test.element, test.want)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	for _, test := range []struct {
		arr    []int
		target int
		left   int
		right  int
		want   int
	}{
		{[]int{1, 2, 5, 8, 11, 13}, 11, 0, 5, 4},
	} {
		got := binarySearch(&test.arr, test.target, test.left, test.right)
		if got != test.want {
			t.Errorf("FAIL::SearchInArray(%v, %v, %v, %v) = %d, want: %d", test.arr, test.target, test.left, test.right, got, test.want)
		} else {
			t.Logf("SUCCESS::SearchInArray(%v, %v, %v, %v) = %d", test.arr, test.target, test.left, test.right, test.want)
		}
	}
}

func TestRotateKthFromEnd(t *testing.T) {
	for _, test := range []struct {
		head []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, []int{1, 2, 5, 4, 3}},
	} {
		got := rotateKthFromEnd(makeListNode(test.head), test.k)
		if common.SliceEqual(makeArrayFromListNode(got), test.want) {
			t.Errorf("FAIL::RotateLastK(%v, %v) = %v, want: %d", test.head, test.k, got, test.want)
		} else {
			t.Logf("SUCCESS::RotateLastK(%v, %v) = %v", test.head, test.k, test.want)
		}
	}
}
