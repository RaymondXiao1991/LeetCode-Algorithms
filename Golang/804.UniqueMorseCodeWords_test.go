package Golang_test

import (
	"Golang"
	"testing"
)

func TestUniqueMorseRepresentations(t *testing.T) {
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
		result := Golang.UniqueMorseRepresentations(test.J, test.S)
		if result != test.want {
			t.Errorf("FAIL::UniqueMorseRepresentations(%v, %v) = %d, want: %d", test.J, test.S, result, test.want)
		} else {
			t.Logf("SUCCESS::UniqueMorseRepresentations(%v, %v) = %d", test.J, test.S, test.want)
		}
	}
}
