package Golang

import (
	"testing"
)

func TestNumJewelsInStones(t *testing.T) {
	for _, test := range []struct {
		J    string
		S    string
		want int
	}{
		{"aA", "a", 1},
		{"aA", "A", 1},
		{"aA", "aA", 2},
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
		{"Bb", "aAAbbbb", 4},
	} {
		got := NumJewelsInStones(test.J, test.S)
		if got != test.want {
			t.Errorf("FAIL::NumJewelsInStones(%v, %v) = %d, want: %d", test.J, test.S, got, test.want)
		} else {
			t.Logf("SUCCESS::NumJewelsInStones(%v, %v) = %d", test.J, test.S, test.want)
		}
	}
}
