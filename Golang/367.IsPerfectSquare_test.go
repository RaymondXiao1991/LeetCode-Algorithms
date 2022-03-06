package Golang

import "testing"

func TestIsPerfectSquare(t *testing.T) {
	for _, test := range []struct {
		x    int
		want bool
	}{
		{1, true},
		{16, true},
		{9, true},
		{14, false},
		{10, false},
	} {
		got := isPerfectSquare(test.x)
		if got != test.want {
			t.Errorf("FAIL::IsPerfectSquare(%v) = %v, want: %v", test.x, got, test.want)
		} else {
			t.Logf("SUCCESS::IsPerfectSquare(%v) = %v", test.x, test.want)
		}
	}
}
