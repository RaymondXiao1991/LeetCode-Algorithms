package Golang

import (
	"fmt"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	for _, test := range []struct {
		n    int
		want [][]int
	}{
		{3, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
		{1, [][]int{{1}}},
	} {
		got := generateMatrix(test.n)
		gotStr := fmt.Sprintf("%v", got)
		t.Logf("%v", gotStr)
		/*
			if got != test.want {
				t.Errorf("FAIL::GenerateMatrix(%d) = %d, want: %d", test.nums, test.target, got, test.want)
			} else {
				t.Logf("SUCCESS::GenerateMatrix(%d) = %d", test.nums, test.target, test.want)
			}
		*/
	}
}
