package Golang

import "testing"

func TestStoneGame(t *testing.T) {
	for _, test := range []struct {
		piles []int
		want  bool
	}{
		{[]int{5, 3, 4, 5}, true},
		{[]int{3, 7, 2, 3}, true},
	} {
		if stoneGame(test.piles) != test.want {
			t.Errorf("FAIL::StoneGame(%v) = %t", test.piles, test.want)
		} else {
			t.Logf("SUCCESS::StoneGame(%v) = %t", test.piles, test.want)
		}
	}
}
