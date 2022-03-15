package common

// 快速排序是一种"分治法"，将原本的问题分解成两个子问题——比基准值小的数和比基准值大的数，然后再分别解决这两个子问题。
// 解决子问题的时候会再次使用快速排序，只有在子问题里只剩下一个数字的时候，排序才算完成。

func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left < right {
		mid := partition(nums, left, right)
		quickSort(nums, left, mid-1)
		quickSort(nums, mid+1, right)
	}
}

// 返回调整完毕之后的base的位置，便于之后对两边的数组进行快排
// 一轮排序完成后，基准值左边的都小于基准值，基准值右边的都大于基准值
// 排序就分成了两个子问题，分别再对基准值左边和右边的数据用快排算法进行排序
// 直到子问题里只剩下一个数字，再也无法分解出子问题的时候，整个序列的排序才算完成
func partition(nums []int, left, right int) int {
	base := nums[left] // 选择一个基准值
	for left < right {
		for left < right && nums[right] >= base { // 大于基准值的往左移
			right--
		}
		// 填补left位置空值
		// right指针值 < base， right值移到left位置
		// right位置值空
		nums[left] = nums[right]

		for left < right && nums[left] <= base { // 小于基准值的往右移
			left++
		}
		// 填补right位置空值
		// left指针值 > base，left值移到right位置
		// left位置值空
		nums[right] = nums[left]
	}
	// base填补left位置的空值
	nums[left] = base

	return left
}

// 优化，这里有两个基本的优化方案：
// 基准值的选择，越靠近中位数越好
// 使用为递归减少递归的层数

func QuickSort2(nums []int) []int {
	quickSort2(nums, 0, len(nums)-1)
	return nums
}

func quickSort2(nums []int, left, right int) {
	for left < right {
		mid := partition2(nums, left, right)
		quickSort2(nums, left, mid-1)
		left = mid + 1
	}
}

func partition2(nums []int, left, right int) int {
	base := threeSumMedian(nums[left], nums[(left+right)>>1], nums[right])
	nums[left], base = base, nums[left]

	for left < right {
		for base < nums[right] && left < right {
			right--
		}
		nums[left] = nums[right]

		for nums[left] <= base && left < right {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = base

	return left
}

func threeSumMedian(a, b, c int) int {
	if a < b {
		a, b = b, a
	}

	if c > a {
		return a
	}

	if c > b {
		return c
	}

	return b
}
