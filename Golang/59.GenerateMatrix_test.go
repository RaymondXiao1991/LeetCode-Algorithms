package Golang

import (
	"fmt"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	for _, test := range []struct {
		n       int
		want    [][]int
		wantStr string
	}{
		{3, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}, "[[1 2 3] [8 9 4] [7 6 5]]"},
		{1, [][]int{{1}}, "[[1]]"},
	} {
		got := generateMatrix(test.n)
		gotStr := fmt.Sprintf("%v", got)
		if gotStr != test.wantStr {
			t.Errorf("FAIL::GenerateMatrix(%d) = %v, want: %v", test.n, got, test.want)
		} else {
			t.Logf("SUCCESS::GenerateMatrix(%d) = %v", test.n, test.want)
		}
	}
}
