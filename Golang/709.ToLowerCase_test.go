package Golang_test

import (
	"Golang"
	"testing"
)

func TestToLowerCase(t *testing.T) {
	for _, test := range []struct {
		Str  string
		want string
	}{
		{"Hello", "hello"},
		{"here", "here"},
		{"LOVELY", "lovely"},
	} {
		got := Golang.ToLowerCase(test.Str)
		if got != test.want {
			t.Errorf("FAIL::ToLowerCase(%v) = %s, want: %s", test.Str, got, test.want)
		} else {
			t.Logf("SUCCESS::ToLowerCase(%v) = %s", test.Str, test.want)
		}
	}
}
