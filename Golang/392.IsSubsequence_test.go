package Golang

import "testing"

func TestIsSubsequence(t *testing.T) {
	for _, test := range []struct {
		s    string
		t    string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	} {
		if isSubsequence(test.s, test.t) != test.want {
			t.Errorf("IsSubsequence(%v, %v) = %t", test.s, test.t, test.want)
		} else {
			t.Logf("SUCCESS::IsSubsequence(%v, %v) = %t", test.s, test.t, test.want)
		}
	}
}

func TestIsSubsequenceInBinary(t *testing.T) {
	for _, test := range []struct {
		s    string
		t    string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	} {
		if isSubsequenceInBinary(test.s, test.t) != test.want {
			t.Errorf("FAIL::IsSubsequenceInBinary(%v, %v) = %t", test.s, test.t, test.want)
		} else {
			t.Logf("SUCCESS::IsSubsequenceInBinary(%v, %v) = %t", test.s, test.t, test.want)
		}
	}
}
