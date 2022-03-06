package Golang

import "testing"

func TestMinWindow(t *testing.T) {
	for _, test := range []struct {
		s    string
		t    string
		want string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
	} {
		got := minWindow(test.s, test.t)
		if got != test.want {
			t.Errorf("FAIL::MinWindow(%v, %v) = %v, want: %v", test.s, test.t, got, test.want)
		} else {
			t.Logf("SUCCESS::MinWindow(%v, %v) = %v", test.s, test.t, test.want)
		}
	}
}
