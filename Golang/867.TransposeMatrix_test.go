package Golang

import (
	"Golang/common"
	"testing"
)

func TestTranspose(t *testing.T) {
	for _, test := range []struct {
		A    [][]int
		want [][]int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, [][]int{{1, 4}, {2, 5}, {3, 6}}},
	} {
		Transpose(test.A)
		got := Transpose(test.A)
		for i, r := range got {
			if common.SliceEqual(r, test.want[i]) {
				t.Logf("SUCCESS::Transpose(%v) = %v", r, test.want[i])
			}
		}
	}
}
