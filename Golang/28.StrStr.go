/*
28. 实现 strStr()

实现 strStr() 函数。
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。

说明：
当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。

示例 1：
输入：haystack = "hello", needle = "ll"
输出：2

示例 2：
输入：haystack = "aaaaa", needle = "bba"
输出：-1

示例 3：
输入：haystack = "", needle = ""
输出：0

提示：
0 <= haystack.length, needle.length <= 5 * 10^4
haystack 和 needle 仅由小写英文字符组成

链接：https://leetcode-cn.com/problems/implement-strstr
*/

package Golang

// KMP算法

func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}

	next := getNext(needle)

	j := 0 // 模式串的起始位置
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}

	return -1
}

// 得到Next数组，next 数组记录的就是最长相同前后缀
func getNext(s string) []int {
	n := len(s)
	next := make([]int, n)
	next[0] = 0 // 第0个为0

	j := 0                   // j游标负责比对前缀后缀是否相同，其值最终记录到next数组中作为每位的回退点
	for i := 1; i < n; i++ { // i游标负责遍历整个模式串
		for j > 0 && s[i] != s[j] { // j要保证大于0，因为下面有取j-1作为数组下标的操作
			j = next[j-1] // 找前一位的对应的回退位置
		}
		if s[i] == s[j] { // 匹配，j和i同时向后移动
			j++ // i的增加已在for循环里
		}
		next[i] = j
	}

	return next
}
