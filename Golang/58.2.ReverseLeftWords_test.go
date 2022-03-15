package Golang

import "testing"

func TestReverseLeftWords(t *testing.T) {
	for _, test := range []struct {
		s    string
		n    int
		want string
	}{
		{"abcdefg", 2, "cdefgab"},
		{"lrloseumgh", 6, "umghlrlose"},
	} {
		got := reverseLeftWords(test.s, test.n)
		if got != test.want {
			t.Errorf("FAIL::ReverseLeftWords(%v, %d) = %v, want: %v", test.s, test.n, got, test.want)
		} else {
			t.Logf("SUCCESS::ReverseLeftWords(%v, %d) = %v", test.s, test.n, test.want)
		}
	}
}
