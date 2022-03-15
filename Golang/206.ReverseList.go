/*
206. 反转链表

给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

示例 1：
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

示例 2：
输入：head = [1,2]
输出：[2,1]

示例 3：
输入：head = []
输出：[]

提示：
链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000

进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？

链接：https://leetcode-cn.com/problems/reverse-linked-list
*/

package Golang

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 在遍历链表时，将当前节点的 next 指针改为指向前一个节点。
// 由于节点没有引用其前一个节点，因此必须事先存储其前一个节点。
// 在更改引用之前，还需要存储后一个节点。最后返回新的头引用。
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		// next := cur.Next
		// cur.Next = prev
		// prev = cur
		// cur = next
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}

func reverseList2(head *ListNode) *ListNode {
	var tail *ListNode
	for head != nil {
		head.Next, tail, head = tail, head, head.Next
	}
	return tail
}
