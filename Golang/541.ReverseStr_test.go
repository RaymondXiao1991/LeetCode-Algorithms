package Golang

import (
	"testing"
)

func TestReverseStr(t *testing.T) {
	for _, test := range []struct {
		s    string
		k    int
		want string
	}{
		{"abcdefg", 2, "bacdfeg"},
		{"abcd", 2, "bacd"},
	} {
		got := reverseStr(test.s, test.k)
		if got != test.want {
			t.Errorf("FAIL::ReverseStr(%v, %d) = %v, want:%v", test.s, test.k, got, test.want)
		} else {
			t.Logf("SUCCESS::ReverseStr(%v, %d) = %v", test.s, test.k, test.want)
		}
	}
}
