/*
25. K个一组翻转链表

给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

进阶：
你可以设计一个只使用常数额外空间的算法来解决此问题吗？
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
*/

package Golang

// 将数组转化为单链表
func makeListNode(array []int) *ListNode {
	if len(array) == 0 {
		return nil
	}

	ln := &ListNode{
		Val: array[0],
	}

	tmp := ln
	for i := 1; i < len(array); i++ {
		tmp.Next = &ListNode{
			Val: array[i],
		}
		tmp = tmp.Next
	}
	return ln
}

// 将单链表转化为数组
func makeArrayFromListNode(head *ListNode) []int {
	if head == nil {
		return nil
	}

	array := []int{}
	for ; head != nil; head = head.Next {
		array = append(array, head.Val)
	}

	return array
}

// 反转区间 [a, b) 的元素
func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	cur := head
	for prev != tail {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return tail, head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	newHead := &ListNode{Next: head}
	pre := newHead

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return newHead.Next
			}
		}
		next := tail.Next
		head, tail := reverse(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		newHead = tail.Next
	}
	return newHead.Next
}

func reverseKGroupInRecursion(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	// 区间 [a, b) 包含 k 个待反转元素
	a, b := &ListNode{Val: head.Val, Next: head.Next}, &ListNode{Val: head.Val, Next: head.Next}
	for i := 0; i < k; i++ {
		// 不足 k 个，不需要反转，base case
		if b == nil {
			return head
		}
		b = b.Next
	}

	// 反转前 k 个元素
	newHead := reverseForRecursion(a, b)
	// 递归反转后续链表并连接起来
	a.Next = reverseKGroup(b, k)
	return newHead
}

/** 反转区间 [a, b) 的元素，注意是左闭右开 */
func reverseForRecursion(head, tail *ListNode) *ListNode {
	pre := &ListNode{}
	cur := &ListNode{}
	nxt := &ListNode{}
	pre = nil
	cur = head
	nxt = head

	// for 终止的条件改一下就行了
	for cur != tail {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	// 返回反转后的头结点
	return pre
}
