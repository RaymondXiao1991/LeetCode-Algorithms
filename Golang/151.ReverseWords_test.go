package Golang

import "testing"

func TestReverseWords(t *testing.T) {
	for _, test := range []struct {
		s    string
		want string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
		{"  Bob    Loves  Alice   ", "Alice Loves Bob"},
		{"Alice does not even like bob", "bob like even not does Alice"},
	} {
		got := reverseWords(test.s)
		if got != test.want {
			t.Errorf("FAIL::ReverseWords(%v) = %v, want:%v", test.s, got, test.want)
		} else {
			t.Logf("SUCCESS::ReverseWords(%v) = %v", test.s, test.want)
		}
	}
}

func TestReverseWords3(t *testing.T) {
	for _, test := range []struct {
		s    string
		want string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
		{"  Bob    Loves  Alice   ", "Alice Loves Bob"},
		{"Alice does not even like bob", "bob like even not does Alice"},
	} {
		got := reverseWords3(test.s)
		if got != test.want {
			t.Errorf("FAIL::ReverseWords(%v) = %v, want:%v", test.s, got, test.want)
		} else {
			t.Logf("SUCCESS::ReverseWords(%v) = %v", test.s, test.want)
		}
	}
}

func TestReverseWords4(t *testing.T) {
	for _, test := range []struct {
		s    string
		want string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
		{"  Bob    Loves  Alice   ", "Alice Loves Bob"},
		{"Alice does not even like bob", "bob like even not does Alice"},
	} {
		got := reverseWords4(test.s)
		if got != test.want {
			t.Errorf("FAIL::ReverseWords(%v) = %v, want:%v", test.s, got, test.want)
		} else {
			t.Logf("SUCCESS::ReverseWords(%v) = %v", test.s, test.want)
		}
	}
}
