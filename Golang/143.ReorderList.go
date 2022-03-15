/*
143. 重排链表

给定一个单链表 L 的头节点 head ，单链表 L 表示为：
L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：
L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1：
输入：head = [1,2,3,4]
输出：[1,4,2,3]

示例 2：
输入：head = [1,2,3,4,5]
输出：[1,5,2,4,3]

提示：
链表的长度范围为 [1, 5 * 10^4]
1 <= node.val <= 1000

链接：https://leetcode-cn.com/problems/reorder-list
*/

package Golang

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	// 找中点，链表分成两个：快指针走完链表时，慢指针刚好只走了一半
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	// 第二个链表倒置
	var secondHalf *ListNode
	for slow != nil {
		slow.Next, secondHalf, slow = secondHalf, slow, slow.Next
	}

	// 链表节点依次连接
	current := head
	for current != secondHalf && current.Next != secondHalf {
		/*
			next1, next2 := current.Next, secondHalf.Next
			current.Next, secondHalf.Next = secondHalf, next1
			current, secondHalf = next1, next2
		*/
		current.Next, secondHalf.Next, current, secondHalf = secondHalf, current.Next, current.Next, secondHalf.Next
	}
}
