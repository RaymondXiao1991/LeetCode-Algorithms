package Golang_test

import (
	"Golang"
	"testing"
)

func TestPeakIndexInMountainArray(t *testing.T) {
	for _, test := range []struct {
		Peak []int
		want int
	}{
		{[]int{0, 1, 0}, 1},
		{[]int{0, 2, 1, 0}, 1},
		{[]int{2, 3, 1, 0}, 1},
	} {
		got := Golang.PeakIndexInMountainArray(test.Peak)
		if got != test.want {
			t.Errorf("FAIL::PeakIndexInMountainArray(%v) = %d, want: %d", test.Peak, got, test.want)
		} else {
			t.Logf("SUCCESS::PeakIndexInMountainArray(%v) = %d", test.Peak, test.want)
		}
	}
}
