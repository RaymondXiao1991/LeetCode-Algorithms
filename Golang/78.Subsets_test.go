package Golang

import (
	"fmt"
	"testing"
)

func TestSubsets(t *testing.T) {
	for _, test := range []struct {
		Arr     []int
		want    [][]int
		wantStr string
	}{
		{[]int{}, [][]int{{}}, "[[]]"},
		{[]int{1}, [][]int{{}, {1}}, "[[] [1]]"},
		{[]int{1, 2}, [][]int{{}, {1}, {1, 2}, {2}}, "[[] [1] [1 2] [2]]"},
		{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}, "[[] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3]]"},
	} {
		got := subsets(test.Arr)
		gotStr := fmt.Sprintf("%v", got)
		if gotStr != test.wantStr {
			t.Errorf("FAIL::Subsets(%v) = %v", test.Arr, test.want)
		} else {
			t.Logf("SUCCESS::Subsets(%v) = %v", test.Arr, test.want)
		}
	}
}
