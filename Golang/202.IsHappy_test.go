package Golang

import "testing"

func TestIsHappy(t *testing.T) {
	for _, test := range []struct {
		n    int
		want bool
	}{
		{19, true},
		{2, false},
	} {
		got := isHappy(test.n)
		if got != test.want {
			t.Errorf("FAIL::IsHappy(%v) = %v, want: %v", test.n, got, test.want)
		} else {
			t.Logf("SUCCESS::IsHappy(%v) = %v", test.n, test.want)
		}
	}
}
