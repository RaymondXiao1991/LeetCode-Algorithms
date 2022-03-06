package Golang

import "testing"

func TestCanWinNim(t *testing.T) {
	for _, test := range []struct {
		n    int
		want bool
	}{
		{4, false},
		{1, true},
		{2, true},
	} {
		if canWinNim(test.n) != test.want {
			t.Errorf("FAIL::CanWinNim(%v) = %t", test.n, test.want)
		} else {
			t.Logf("SUCCESS::CanWinNim(%v) = %t", test.n, test.want)
		}
	}
}
