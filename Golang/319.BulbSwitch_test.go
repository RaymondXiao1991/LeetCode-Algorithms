package Golang

import "testing"

func TestBulbSwitch(t *testing.T) {
	for _, test := range []struct {
		n    int
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
		{16, 4},
	} {
		if bulbSwitch(test.n) != test.want {
			t.Errorf("FAIL::BulbSwitch(%d) = %d", test.n, test.want)
		} else {
			t.Logf("SUCCESS::BulbSwitch(%d) = %d", test.n, test.want)
		}
	}
}
