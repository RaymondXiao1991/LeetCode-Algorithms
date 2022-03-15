package Golang

import (
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	for _, test := range []struct {
		headA []int
		headB []int
		want  []int
	}{
		{[]int{4, 1, 8, 4, 5}, []int{5, 0, 1, 8, 4, 5}, []int{1, 2, 3, 5}},
	} {
		got := getIntersectionNode(makeListNode(test.headA), makeListNode(test.headB))
		gotListNode := makeArrayFromListNode(got)
		t.Logf("Intersected at :%v", gotListNode)
	}
}
