package Golang

import "testing"

func TestShipWithinDays(t *testing.T) {
	for _, test := range []struct {
		weights []int
		days    int
		want    int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 15},
		{[]int{3, 2, 2, 4, 1, 4}, 3, 6},
		{[]int{1, 2, 3, 1, 1}, 4, 3},
	} {
		got := shipWithinDays(test.weights, test.days)
		if got != test.want {
			t.Errorf("FAIL::ShipWithinDays(%v, %v) = %d, want: %d", test.weights, test.days, got, test.want)
		} else {
			t.Logf("SUCCESS::ShipWithinDays(%v, %v) = %d", test.weights, test.days, test.want)
		}
	}
}
