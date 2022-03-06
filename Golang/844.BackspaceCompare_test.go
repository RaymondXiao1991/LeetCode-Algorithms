package Golang

import (
	"testing"
)

func TestBackspaceCompare(t *testing.T) {
	for _, test := range []struct {
		s    string
		t    string
		want bool
	}{
		{"ab#c", "ad#c", true},
		{"ab##", "c#d#", true},
		{"a#c", "b", false},
	} {
		got := backspaceCompare(test.s, test.t)
		if got != test.want {
			t.Errorf("FAIL::BackspaceCompare(%v, %v) = %t, want:%t", test.s, test.t, got, test.want)
		} else {
			t.Logf("SUCCESS::BackspaceCompare(%v, %v) = %t", test.s, test.t, test.want)
		}
	}
}
