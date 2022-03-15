/* ************************************************************************
> File Name:     main.go
> Author:        程序员Carl
> 微信公众号:    代码随想录
> Created Time:  二  3/15 15:32:01 2022
> Description:   
************************************************************************/

package main

import "fmt"


func QuickSort(nums []int) []int {
    quickSort(nums, 0, len(nums)-1)
    return nums
}

func quickSort(nums []int, left, right int) {
    if left < right {
        mid := partition(nums, left, right)
        fmt.Println("mid:", mid)
        fmt.Println("quickSort左半边前:", nums)
        quickSort(nums, left, mid-1)
        fmt.Println("quickSort左半边后:", nums)
        quickSort(nums, mid+1, right)
        fmt.Println("quickSort右半边后:", nums)
    }
}

func partition(nums []int, left, right int) int {
    base := nums[left]
    for left < right {
        for left < right && nums[right] >= base {
            right--
        }
        fmt.Println("right指针左移前:", nums)
        nums[left] = nums[right]
        fmt.Println("right指针左移后:", nums)

        for left < right && nums[left] <= base {
            left++
        }
        fmt.Println("left指针右移前:", nums)
        nums[right] = nums[left]
        fmt.Println("left指针右移后:", nums)
    }
    fmt.Println("基准值复位前:", nums)
    nums[left] = base
    fmt.Println("基准值复位后:", nums)

    return left
}

func main() {
    got := QuickSort2([]int{2,6,7,1,9,10})
    fmt.Println("got:", got)
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
        fmt.Println("mid:", mid)
        fmt.Println("quickSort左半边前:", nums)
		quickSort2(nums, left, mid-1)
        fmt.Println("quickSort左半边后:", nums)
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
        fmt.Println("right指针左移前:", nums)
		nums[left] = nums[right]
        fmt.Println("right指针左移后:", nums)

		for nums[left] <= base && left < right {
			left++
		}
        fmt.Println("left指针右移前:", nums)
		nums[right] = nums[left]
        fmt.Println("left指针右移后:", nums)
	}
    fmt.Println("基准值复位前:", nums)
	nums[left] = base
    fmt.Println("基准值复位后:", nums)

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
