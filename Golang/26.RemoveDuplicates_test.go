package Golang

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	for _, test := range []struct {
		Arr  []int
		want int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	} {
		if removeDuplicates(test.Arr) != test.want {
			t.Errorf("FAIL::RemoveDuplicates(%v) = %d", test.Arr, test.want)
		} else {
			t.Logf("SUCCESS::RemoveDuplicates(%v) = %d", test.Arr, test.want)
		}
	}
}
