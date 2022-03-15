/*
19. 删除链表的倒数第 N 个结点

给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

示例 1：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

示例 2：
输入：head = [1], n = 1
输出：[]

示例 3：
输入：head = [1,2], n = 1
输出：[1]

提示：
链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

进阶：你能尝试使用一趟扫描实现吗？

链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
*/

package Golang

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fakeHead := &ListNode{
		Next: head,
	}

	count := 0
	var recurseFn func(*ListNode)
	recurseFn = func(node *ListNode) {
		if node == nil {
			return
		}

		recurseFn(node.Next)

		count++
		if count == n+1 {
			node.Next = node.Next.Next
		}
	}

	recurseFn(fakeHead)
	return fakeHead.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	fakeHead := &ListNode{
		Next: head,
	}
	prev := fakeHead
	cur := head
	i := 1
	for cur != nil {
		cur = cur.Next
		if i > n {
			prev = prev.Next
		}
		i++
	}
	prev.Next = prev.Next.Next
	return fakeHead.Next
}
