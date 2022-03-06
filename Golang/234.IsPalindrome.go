/*
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

示例 1：
输入：head = [1,2,2,1]
输出：true

示例 2：
输入：head = [1,2]
输出：false

提示：
链表中节点数目在范围[1, 105] 内
0 <= Node.val <= 9

链接：https://leetcode-cn.com/problems/palindrome-linked-list
*/

package Golang

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 将单链表中所有值拷贝到数组，然后首尾比较
func isPalindrome(head *ListNode) bool {
	vals := []int{}
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}

	n := len(vals)
	for i, v := range vals[:n/2] {
		if v != vals[n-i-1] {
			return false
		}
	}

	return true
}

func isPalindromeForArray(array []int) bool {
	if array == nil {
		return true
	}

	n := len(array)
	for i, v := range array[:n/2] {
		if v != array[n-i-1] {
			return false
		}
	}

	return true
}

func isPalindromeForStr(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 快慢指针
func isPalindromeByQuickSlowCursor(head *ListNode) bool {
	var prev *ListNode = nil
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next // 快指针一次走2步
		// 反转慢指针走过的路径，以下4行代码与"反转单链表"一模一样
		next := slow.Next // 临时保存下一个节点
		slow.Next = prev  // slow反指向前一个节点
		prev = slow       // pre后移一位
		slow = next       // slow后移一位
	}
	// 长度为偶数时，fast最终必不为 null
	// 长度为奇数时，当fast.next为null时，fast必为null
	if fast != nil {
		slow = slow.Next
	}
	// 判断 pre 和 slow 是否为回文
	for prev != nil && slow != nil {
		if prev.Val != slow.Val {
			return false
		}
		prev, slow = prev.Next, slow.Next
	}
	return true
}
