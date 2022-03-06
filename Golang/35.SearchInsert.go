/*
35. 搜索插入位置

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
请必须使用时间复杂度为 O(log n) 的算法。

示例 1:
输入: nums = [1,3,5,6], target = 5
输出: 2

示例 2:
输入: nums = [1,3,5,6], target = 2
输出: 1

示例 3:
输入: nums = [1,3,5,6], target = 7
输出: 4

示例 4:
输入: nums = [1,3,5,6], target = 0
输出: 0

示例 5:
输入: nums = [1], target = 0
输出: 0

提示:
1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 为无重复元素的升序排列数组
-104 <= target <= 104

链接：https://leetcode-cn.com/problems/search-insert-position
*/

package Golang

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)
	// 在区间 nums[left, right] 里查找第 1 个大于等于 target 的元素的下标
	// 当 left 与 right 重合的时候，搜索终止
	for left < right {
		mid := left + ((right - left) >> 1)
		if nums[mid] >= target {
			// 下一轮搜索的区间是 [left, mid]
			right = mid
		} else {
			// 下一轮搜索的区间是 [mid+1, right]
			left = mid + 1
		}
	}
	return left
}
