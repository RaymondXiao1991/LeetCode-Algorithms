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
	i, sum := start, 0
	for start < end {
		if nums[i] == elem {
			sum++
		}
		i++
	}
	return sum
}
