/*
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

进阶：
如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？

致谢：
特别感谢 @pbrother 添加此问题并且创建所有测试用例。

示例 1：
输入：s = "abc", t = "ahbgdc"
输出：true

示例 2：
输入：s = "axc", t = "ahbgdc"
输出：false

提示：
0 <= s.length <= 100
0 <= t.length <= 10^4
两个字符串都只由小写字符组成。

链接：https://leetcode-cn.com/problems/is-subsequence
*/

package Golang

// 双指针法
func isSubsequence(s string, t string) bool {
	m, n := len(s), len(t)
	i, j := 0, 0
	for i < m && j < n {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return i == m
}

// 二分查找法
func isSubsequenceInBinary(s string, t string) bool {
	// 构造一个二维表，记录每个字母出现的位置
	indexs := make([][]int, 26) //也可以直接256 asc码表的个数
	for i := 0; i < len(t); i++ {
		indexs[int(t[i]-'a')] = append(indexs[int(t[i]-'a')], i)
	}
	// 下面遍历s
	target := 0
	for j := 0; j < len(s); j++ {
		if len(indexs[s[j]-'a']) == 0 {
			return false
		}
		pos := findLeftBound(indexs[s[j]-'a'], target) // 找到满足>=target的index的坐标
		if pos == len(indexs[s[j]-'a']) {
			return false
		}
		target = indexs[s[j]-'a'][pos] + 1
	}
	return true
}

func findLeftBound(indexes []int, target int) int {
	lo, hi, mid := 0, len(indexes), 0
	for lo < hi {
		mid = lo + (hi-lo)>>1
		if indexes[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
