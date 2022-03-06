package Golang

import "testing"

func TestTotalFruit(t *testing.T) {
	for _, test := range []struct {
		fruits []int
		want   int
	}{
		{[]int{1, 2, 1}, 3},
		{[]int{0, 1, 2, 2}, 3},
		{[]int{1, 2, 3, 2, 2}, 4},
		{[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}, 5},
	} {
		got := totalFruit(test.fruits)
		if got != test.want {
			t.Errorf("FAIL::TotalFruit(%v) = %d, want:%d", test.fruits, got, test.want)
		} else {
			t.Logf("SUCCESS::TotalFruit(%v) = %d", test.fruits, test.want)
		}
	}
}
