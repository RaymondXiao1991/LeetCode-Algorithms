package Golang

import (
	"Golang/common"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	/*
		for _, test := range []struct {
			head []int
			k    int
			want []int
		} {
			{[]int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
			{[]int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
			{[]int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5}},
			{[]int{1}, 1, []int{1}},
		} {
			head := &ListNode{Val:1, Next: nil}
			head.Next = &ListNode{Val: 2, Next: nil}
			head.Next.Next = &ListNode{Val: 3, Next: nil}
			head.Next.Next.Next = &ListNode{Val: 4, Next: nil}
			if got != test.want {
				t.Errorf("FAIL::LRUCache(%v, %v) = %d, want: %d", test.Key, test.Val, got, test.want)
			} else {
				t.Logf("SUCCESS::LRUCache(%v, %v) = %d", test.Key, test.Val, test.want)
			}
		}
	*/
	got := reverseKGroup(
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
		2)
	t.Logf("%v", got)
}

func TestMakeListNode(t *testing.T) {
	for _, test := range []struct {
		array []int
		want  bool
	}{
		{[]int{1}, true},
		{[]int{1, 3, 5}, true},
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 3, 1, 4, 5}, true},
	} {
		got := makeArrayFromListNode(makeListNode(test.array))
		if common.SliceEqual(got, test.array) {
			t.Logf("SUCCESS::MakeListNode(%v) = %t", test.array, test.want)
		} else {
			t.Errorf("FAIL::MakeListNode(%v) = %t", test.array, test.want)
		}
	}
}
