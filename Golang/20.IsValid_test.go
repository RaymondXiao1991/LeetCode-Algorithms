package Golang

import "testing"

func TestIsValid(t *testing.T) {
	for _, test := range []struct {
		s    string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	} {
		if isValid(test.s) != test.want {
			t.Errorf("FAIL::IsValid(%v), want: %v", test.s, test.want)
		} else {
			t.Logf("SUCCESS::IsValid(%v) = %t", test.s, test.want)
		}
	}
}
