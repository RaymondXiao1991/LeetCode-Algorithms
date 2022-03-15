package Golang

import "testing"

func TestCandy(t *testing.T) {
	for _, test := range []struct {
		ratings []int
		want    int
	}{
		{[]int{1, 0, 2}, 5},
		{[]int{1, 2, 2}, 4},
	} {
		got := candy(test.ratings)
		if got != test.want {
			t.Errorf("FAIL::Candy(%v) = %v, want:%v", test.ratings, got, test.want)
		} else {
			t.Logf("SUCCESS::Candy(%v) = %v", test.ratings, test.want)
		}
	}
}
