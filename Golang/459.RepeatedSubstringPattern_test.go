package Golang

import "testing"

func TestRepeatedSubstringPattern(t *testing.T) {
	for _, test := range []struct {
		s    string
		want bool
	}{
		{"abab", true},
		{"aba", false},
		{"abcabcabcabc", true},
	} {
		got := repeatedSubstringPattern(test.s)
		if got != test.want {
			t.Errorf("FAIL::RepeatedSubstringPattern(%v) = %v, want: %v", test.s, got, test.want)
		} else {
			t.Logf("SUCCESS::RepeatedSubstringPattern(%v) = %v", test.s, test.want)
		}
	}
}
