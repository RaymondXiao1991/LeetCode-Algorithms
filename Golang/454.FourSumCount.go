/*
454. 四数相加 II

给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：
0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0

示例 1：
输入：nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
输出：2
解释：
两个元组如下：
1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0

示例 2：
输入：nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
输出：1

提示：
n == nums1.length
n == nums2.length
n == nums3.length
n == nums4.length
1 <= n <= 200
-2^28 <= nums1[i], nums2[i], nums3[i], nums4[i] <= 2^28

链接：https://leetcode-cn.com/problems/4sum-ii
*/

package Golang

// 方法一：分组 + 哈希表
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count, hashMap := 0, map[int]struct{}{}

	// 先求出 A 和 B 任意两数之和 sumAB
	for _, v1 := range nums1 {
		// 以 sumAB 为 key，sumAB 出现的次数为 value，存入 hashmap 中
		for _, v2 := range nums2 {
			hashMap[v1+v2] = struct{}{}
		}
	}

	// 然后计算 C 和 D 中任意两数之和的相反数 sumCD，在 hashmap 中查找是否存在 key 为 sumCD。
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			if _, ok := hashMap[-v3-v4]; ok {
				count++
			}
		}
	}

	return count
}
