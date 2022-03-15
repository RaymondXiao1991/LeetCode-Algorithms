package Golang

import "testing"

func TestIsAnagram(t *testing.T) {
	for _, test := range []struct {
		s    string
		t    string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	} {
		got := isAnagram(test.s, test.t)
		if got != test.want {
			t.Errorf("FAIL::IsAnagram(%s, %s) = %v, want: %v", test.s, test.t, got, test.want)
		} else {
			t.Logf("SUCCESS::IsAnagram(%s, %s) = %v", test.s, test.t, test.want)
		}
	}
}
