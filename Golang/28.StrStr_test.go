package Golang

import "testing"

func TestStrStr(t *testing.T) {
	for _, test := range []struct {
		haystack string
		needle   string
		want     int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"", "", 0},
	} {
		got := strStr(test.haystack, test.needle)
		if got != test.want {
			t.Errorf("FAIL::StrStr(%s, %s) = %v, want: %v", test.haystack, test.needle, got, test.want)
		} else {
			t.Logf("SUCCESS::StrStr(%s, %s) = %v", test.haystack, test.needle, test.want)
		}
	}
}
