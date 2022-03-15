package Golang

import (
	"fmt"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	for _, test := range []struct {
		head []int
		want []int
	}{
		{[]int{3, 2, 0, -4}, []int{}},
		{[]int{3, 2, 0, -4}, []int{}},
	} {
		got := detectCycle(makeListNode(test.head))
		gotStr := fmt.Sprintf("%v", got)
		t.Logf("%v", gotStr)
		/*
			if gotStr != test.wantStr {
				t.Errorf("FAIL::GenerateMatrix(%d) = %v, want: %v", test.n, got, test.want)
			} else {
				t.Logf("SUCCESS::GenerateMatrix(%d) = %v", test.n, test.want)
			}
		*/
	}
}

func TestDetectCycle2(t *testing.T) {
	for _, test := range []struct {
		head []int
		want []int
	}{
		{[]int{3, 2, 0, -4}, []int{}},
		{[]int{3, 2, 0, -4}, []int{}},
	} {
		got := detectCycle2(makeListNode(test.head))
		gotStr := fmt.Sprintf("%v", got)
		t.Logf("%v", gotStr)
		/*
			if gotStr != test.wantStr {
				t.Errorf("FAIL::GenerateMatrix(%d) = %v, want: %v", test.n, got, test.want)
			} else {
				t.Logf("SUCCESS::GenerateMatrix(%d) = %v", test.n, test.want)
			}
		*/
	}
}
