package Golang

import (
	"testing"
)

func TestHammingDistance(t *testing.T) {
	for _, test := range []struct {
		X    int
		Y    int
		want int
	}{
		{1, 4, 2},
	} {
		got := HammingDistance(test.X, test.Y)
		if got != test.want {
			t.Errorf("FAIL::HammingDistance(%v, %v) = %d, want: %d", test.X, test.Y, got, test.want)
		} else {
			t.Logf("SUCCESS::HammingDistance(%v, %v) = %d", test.X, test.Y, test.want)
		}
	}
}
