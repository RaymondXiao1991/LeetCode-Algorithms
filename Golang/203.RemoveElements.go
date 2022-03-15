/*
203. 移除链表元素

给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

示例 1：
输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]

示例 2：
输入：head = [], val = 1
输出：[]

示例 3：
输入：head = [7,7,7,7], val = 7
输出：[]

提示：
列表中的节点数目在范围 [0, 10^4] 内
1 <= Node.val <= 50
0 <= val <= 50

链接：https://leetcode-cn.com/problems/remove-linked-list-elements
*/

package Golang

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	// 防止第1个元素就是要删的元素，被删之后没有头部
	fakeHead := &ListNode{
		Next: head,
	}

	for current := fakeHead; current.Next != nil; {
		if current.Next.Val == val {
			current.Next = current.Next.Next // 删除
		} else {
			current = current.Next // 保留
		}
	}

	return fakeHead.Next
}
