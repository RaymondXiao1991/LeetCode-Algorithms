package Golang_test

import (
	"Golang"
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
		result := Golang.HammingDistance(test.X, test.Y)
		if result != test.want {
			t.Errorf("FAIL::HammingDistance(%v, %v) = %d, want: %d", test.X, test.Y, result, test.want)
		} else {
			t.Logf("SUCCESS::HammingDistance(%v, %v) = %d", test.X, test.Y, test.want)
		}
	}
}
