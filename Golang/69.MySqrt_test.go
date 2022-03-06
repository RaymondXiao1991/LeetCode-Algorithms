package Golang

import (
	"testing"
)

func TestMySqrt(t *testing.T) {
	for _, test := range []struct {
		x    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 2},
		{6, 2},
		{7, 2},
		{8, 2},
		{9, 3},
		{10, 3},
		{11, 3},
		{12, 3},
		{13, 3},
		{14, 3},
		{15, 3},
		{16, 4},
	} {
		got := mySqrt(test.x)
		if got != test.want {
			t.Errorf("FAIL::MySqrt(%v) = %d, want: %d", test.x, got, test.want)
		} else {
			t.Logf("SUCCESS::MySqrt(%v) = %d", test.x, test.want)
		}
	}
}
