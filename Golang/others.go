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
	arr := nums[start:end+1]
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
