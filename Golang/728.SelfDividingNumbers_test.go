package Golang_test

import (
	"Golang"
	"common"
	"testing"
)

func TestSelfDividingNumbers(t *testing.T) {
	for _, test := range []struct {
		Left  int
		Right int
		want  []int
	}{
		{1, 22, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}},
		{1, 11, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11}},
		{1, 16, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15}},
	} {
		got := Golang.SelfDividingNumbers(test.Left, test.Right)
		if !common.SliceEqual(got, test.want) {
			t.Errorf("FAIL::SelfDividingNumbers(%v, %v) = %v, want: %v", test.Left, test.Right, got, test.want)
		} else {
			t.Logf("SUCCESS::SelfDividingNumbers(%v, %v) = %d", test.Left, test.Right, test.want)
		}
	}
}
