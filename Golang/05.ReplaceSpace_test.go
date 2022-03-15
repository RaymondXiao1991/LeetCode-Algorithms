package Golang

import "testing"

func TestReplaceSpace(t *testing.T) {
	for _, test := range []struct {
		s    string
		want string
	}{
		{"We are happy.", "We%20are%20happy."},
	} {
		got := replaceSpace(test.s)
		if got != test.want {
			t.Errorf("FAIL::ReplaceSpace(%v) = %v, want:%v", test.s, got, test.want)
		} else {
			t.Logf("SUCCESS::ReplaceSpace(%v) = %v", test.s, test.want)
		}
	}
}
