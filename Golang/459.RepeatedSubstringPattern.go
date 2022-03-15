/*
459. 重复的子字符串

给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。

示例 1:
输入: s = "abab"
输出: true
解释: 可由子串 "ab" 重复两次构成。

示例 2:
输入: s = "aba"
输出: false

示例 3:
输入: s = "abcabcabcabc"
输出: true
解释: 可由子串 "abc" 重复四次构成。 (或子串 "abcabc" 重复两次构成。)

提示：
1 <= s.length <= 10^4
s 由小写英文字母组成

链接：https://leetcode-cn.com/problems/repeated-substring-pattern
*/

package Golang

// 数组长度减去最长相同前后缀的长度相当于是第一个周期的长度，也就是一个周期的长度，如果这个周期可以被整除，就说明整个数组就是这个周期的循环。
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}

	next := getNext(s)
	// 如果 next[len-1] != 0，则说明字符串有最长相同的前后缀（即字符串里的前缀子串和后缀子串相同的最长长度）
	// 最长相等前后缀的长度为：next[n-1]
    // 如果 n%(n-(next[n-1])) == 0 ，则说明 (数组长度-最长相等前后缀的长度) 正好可以被 数组的长度整除，说明有该字符串有重复的子字符串。
	if next[n-1] != 0 && n%(n-(next[n-1])) == 0 {
		return true
	}

	return false
}
