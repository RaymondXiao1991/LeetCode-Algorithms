package Golang

import "testing"

func TestCanConstruct(t *testing.T) {
	for _, test := range []struct {
		ransomNote string
		magazine   string
		want       bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	} {
		got := canConstruct(test.ransomNote, test.magazine)
		if got != test.want {
			t.Errorf("FAIL::CanConstruct(%v, %v) = %v, want: %v", test.ransomNote, test.magazine, got, test.want)
		} else {
			t.Logf("SUCCESS::CanConstruct(%v, %v) = %v", test.ransomNote, test.magazine, test.want)
		}
	}
}
