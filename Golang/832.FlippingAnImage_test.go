package Golang_test

import (
	"Golang"
	"common"
	"testing"
)

func TestFlipAndInvertImage(t *testing.T) {
	for _, test := range []struct {
		Imgs [][]int
		want [][]int
	}{
		{[][]int{[]int{1, 1, 0}, []int{1, 0, 1}, []int{0, 0, 0}}, [][]int{[]int{1, 0, 0}, []int{0, 1, 0}, []int{1, 1, 1}}},
	} {
		result := Golang.FlipAndInvertImage(test.Imgs)
		for i, r := range result {
			if common.SliceEqual(r, test.want[i]) {
				//t.Errorf("FAIL::FlipAndInvertImage(%v) = %v, want: %v", test.Imgs, r, w)
				t.Logf("SUCCESS::FlipAndInvertImage(%v) = %v", r, test.want[i])
			}
		}
	}
}
