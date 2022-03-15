package Golang

import (
	"Golang/common"
	"testing"
)

func TestReverseString(t *testing.T) {
	for _, test := range []struct {
		s    []byte
		want []byte
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	} {
		originS := common.DeepCopy(test.s)
		reverseString(test.s)
		got := test.s
		if !common.SliceEqual(got, test.want) {
			t.Errorf("FAIL::ReverseString(%v) = %v, want:%v", originS, got, test.want)
		} else {
			t.Logf("SUCCESS::ReverseString(%v) = %v", originS, test.want)
		}
	}
}
