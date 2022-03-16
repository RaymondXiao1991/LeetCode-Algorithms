package Golang

/*
要求给定一个固定数组，一个随机区间范围与一个元素值。返回区间内元素个数。
例子
list = [ 3, 3, 5, 6, 1, 2, 7, 8, 0, 1 ]
start = 4
end = 9
element = 1
表示 查询 list[4:9] 内 1的个数
返回 2
要求时间复杂度在O(logn)
*/
func searchInArray(nums []int, start, end, elem int) int {
	arr := nums[start : end+1]
	sum := 0
	for _, v := range arr {
		if v == elem {
			sum++
		}
	}
	return sum
}

// 二分查找
func binarySearch(arr *[]int, target int, left, right int) int {
	// 在arr[left, right]中查找target
	for left <= right {
		mid := left + ((right - left) >> 1)
		if (*arr)[mid] > target {
			right = mid - 1
		} else if (*arr)[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// 链表旋转最后k个节点
// head: 1->2->3->4->5, k: 3, want: 1->2->5->4->3
func rotateKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 找k点，链表分成两个
	h1 := head
	firstHalf := &ListNode{Next: h1}
	for h1 != nil && k > 0 {
		h1 = h1.Next
		firstHalf.Next = h1.Next
		k--
	}

	// 找到后部分链表
	h2 := head
	for h1 != nil {
		h1, h2 = h1.Next, h2.Next
	}

	// 反转后部分链表
	var tail *ListNode
	for h2 != nil {
		h2.Next, tail, h2 = tail, h2, h2.Next
	}

	// 两部分连接到一起
	firstHalf.Next = tail

	return firstHalf
}
