package Golang_test

import (
	"Golang"
	"testing"
)

func TestJudgeCircle(t *testing.T) {
	for _, test := range []struct {
		Moves string
		want  bool
	}{
		{"UD", true},
		{"LL", false},
		{"UDLR", true},
	} {
		got := Golang.JudgeCircle(test.Moves)
		if got != test.want {
			t.Errorf("FAIL::JudgeCircle(%v) = %t, want: %t", test.Moves, got, test.want)
		} else {
			t.Logf("SUCCESS::JudgeCircle(%v) = %t", test.Moves, test.want)
		}
	}
}
