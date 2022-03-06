/*
977. 有序数组的平方

给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

示例 1：
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]

示例 2：
输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]

提示：
1 <= nums.length <= 10^4
-10^4 <= nums[i] <= 10^4
nums 已按 非递减顺序 排序

进阶：
请你设计时间复杂度为 O(n) 的算法解决本问题

链接：https://leetcode-cn.com/problems/squares-of-a-sorted-array
*/

package Golang

func sortedSquares(nums []int) []int {
	n := len(nums)
	squares := make([]int, n)
	left, right := 0, n-1
	for i := n - 1; i >= 0; i-- {
		leftSquare, rightSquare := nums[left]*nums[left], nums[right]*nums[right]
		if leftSquare > rightSquare {
			squares[i] = leftSquare
			left++
		} else {
			squares[i] = rightSquare
			right--
		}
	}
	return squares
}
