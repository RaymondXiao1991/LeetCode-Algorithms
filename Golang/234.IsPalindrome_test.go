package Golang

import "testing"

func TestIsPalindrome(t *testing.T) {
	for _, test := range []struct {
		array []int
		want  bool
	}{
		{[]int{1, 2, 2, 1}, true},
		{[]int{1, 2}, false},
		{[]int{1, 2, 3, 2, 1}, true},
	} {
		if isPalindrome(makeListNode(test.array)) != test.want {
			t.Errorf("FAIL::IsPalindrome(%v), want: %v", test.array, test.want)
		} else {
			t.Logf("SUCCESS::IsPalindrome(%v) = %t", test.array, test.want)
		}
	}
}

func TestIsPalindromeByQuickSlowCursor(t *testing.T) {
	for _, test := range []struct {
		array []int
		want  bool
	}{
		{[]int{1, 2, 2, 1}, true},
		{[]int{1, 2}, false},
		{[]int{1, 2, 3, 2, 1}, true},
	} {
		if isPalindromeByQuickSlowCursor(makeListNode(test.array)) != test.want {
			t.Errorf("FAIL::IsPalindromeByQuickSlowCursor(%v), want: %v", test.array, test.want)
		} else {
			t.Logf("SUCCESS::IsPalindromeByQuickSlowCursor(%v) = %t", test.array, test.want)
		}
	}
}

func TestIsPalindromeForStr(t *testing.T) {
	for _, test := range []struct {
		s    string
		want bool
	}{
		{"1221", true},
		{"12", false},
		{"12321", true},
	} {
		if isPalindromeForStr(test.s) != test.want {
			t.Errorf("FAIL::IsPalindromeForStr(%v), want: %v", test.s, test.want)
		} else {
			t.Logf("SUCCESS::IsPalindromeForStr(%v) = %t", test.s, test.want)
		}
	}
}

func TestIsPalindromeForArray(t *testing.T) {
	for _, test := range []struct {
		array []int
		want  bool
	}{
		{[]int{1, 2, 2, 1}, true},
		{[]int{1, 2}, false},
		{[]int{1, 2, 3, 2, 1}, true},
	} {
		if isPalindromeForArray(test.array) != test.want {
			t.Errorf("FAIL::IsPalindromeForArray(%v), want: %v", test.array, test.want)
		} else {
			t.Logf("SUCCESS::IsPalindromeForArray(%v) = %t", test.array, test.want)
		}
	}
}
