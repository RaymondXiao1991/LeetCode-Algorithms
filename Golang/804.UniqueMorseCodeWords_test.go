package Golang_test

import (
	"Golang"
	"testing"
)

func TestUniqueMorseRepresentations(t *testing.T) {
	for _, test := range []struct {
		Words []string
		want  int
	}{
		{[]string{"gin", "zen", "gig", "msg"}, 2},
	} {
		got := Golang.UniqueMorseRepresentations(test.Words)
		if got != test.want {
			t.Errorf("FAIL::UniqueMorseRepresentations(%v) = %d, want: %d", test.Words, got, test.want)
		} else {
			t.Logf("SUCCESS::UniqueMorseRepresentations(%v) = %d", test.Words, test.want)
		}
	}
}
