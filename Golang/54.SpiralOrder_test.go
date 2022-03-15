package Golang

import (
	"Golang/common"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	for _, test := range []struct {
		matrix [][]int
		want   []int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	} {
		got := spiralOrder(test.matrix)
		if !common.SliceEqual(got, test.want) {
			t.Errorf("FAIL::SpiralOrder(%v) = %v, want: %v", test.matrix, got, test.want)
		} else {
			t.Logf("SUCCESS::SpiralOrder(%v) = %v", test.matrix, test.want)
		}
	}
}
