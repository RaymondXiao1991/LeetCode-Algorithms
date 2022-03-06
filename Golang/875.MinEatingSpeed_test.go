package Golang

import (
	"testing"
)

func TestMinEatingSpeed(t *testing.T) {
	for _, test := range []struct {
		piles []int
		H     int
		want  int
	}{
		{[]int{3, 6, 7, 11}, 8, 4},
		{[]int{30, 11, 23, 4, 20}, 5, 30},
	} {
		got := minEatingSpeed(test.piles, test.H)
		if got != test.want {
			//t.Errorf("FAIL::MinEatingSpeed(%v, %v) = %d, want: %d", test.piles, test.H, got, test.want)
		} else {
			t.Logf("SUCCESS::MinEatingSpeed(%v, %v) = %d", test.piles, test.H, test.want)
		}
	}
}
