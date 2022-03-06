/*
34. 在排序数组中查找元素的第一个和最后一个位置

给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。

进阶：
你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？

示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

示例 2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]

示例 3：
输入：nums = [], target = 0
输出：[-1,-1]

提示：
0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums 是一个非递减数组
-109 <= target <= 109

链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
*/

package Golang

import (
	"sort"
	"sync"
)

func searchRange(nums []int, target int) []int {
	leftFn := func(nums []int, target int) int {
		left, right := 0, len(nums)
		for left < right {
			mid := left + ((right - left) >> 1)
			if nums[mid] > target {
				right = mid
			} else {
				left = mid + 1
			}
		}
		return left
	}

	leftMost := leftFn(nums, target-1)
	if leftMost == len(nums) || nums[leftMost] != target {
		return []int{-1, -1}
	}

	rightMost := leftFn(nums, target)
	return []int{leftMost, rightMost - 1}
}

func searchRangeByUseSort(nums []int, target int) []int {
	leftMost := sort.SearchInts(nums, target)
	if leftMost == len(nums) || nums[leftMost] != target {
		return []int{-1, -1}
	}

	rightMost := sort.SearchInts(nums, target+1) - 1
	return []int{leftMost, rightMost}
}

func searchRangeByUseWG(nums []int, target int) []int {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	n := len(nums)
	left, right := -1, -1

	go func() {
		defer wg.Done()

		for i := 0; i < n; i++ {
			if nums[i] == target {
				left = i
				break
			}
		}
	}()

	go func() {
		defer wg.Done()

		for i := n - 1; i >= 0; i-- {
			if nums[i] == target {
				right = i
				break
			}
		}
	}()

	wg.Wait()

	return []int{left, right}
}
