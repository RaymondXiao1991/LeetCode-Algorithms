/*
剑指 Offer 05. 替换空格

请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

示例 1：
输入：s = "We are happy."
输出："We%20are%20happy."

限制：
0 <= s 的长度 <= 10000

链接：https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof
*/

package Golang

func replaceSpace(s string) string {
	b := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			b = append(b, "%20"...)
		} else {
			b = append(b, s[i])
		}
	}
	return string(b)
}
