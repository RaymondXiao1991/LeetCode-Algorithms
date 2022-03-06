/*
76. 最小覆盖子串

给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
注意：
对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。

示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"

示例 2：
输入：s = "a", t = "a"
输出："a"

示例 3:
输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。

提示：
1 <= s.length, t.length <= 10^5
s 和 t 由英文字母组成

进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？

链接：https://leetcode-cn.com/problems/minimum-window-substring
*/

package Golang

/* 模板
int[] arrayT = new int[52];
for (char c : t.toCharArray()) {
    arrayT[getIndex(c)]++;
}
int l = 0;
int r = 0;
int valid = 0;
while (r < s.length()) {
    char cur = s.charAt(r);
    r++;
    // 逻辑代码,窗口内数据更新

    while(needLeftMove()){
        // 将要移除的元素
        char d = s.charAt(l);
        l++;
        // 逻辑代码,窗口内数据更新
    }
}

return "";
*/

func minWindow(s string, t string) string {
	n := len(s)
	minLen := n + 1
	start := n                                    // 结果子串的起始位置
	targetCharWithMissingNumMap := map[rune]int{} // 存储目标字符和对应的缺失个数
	missingType := 0                              // 当前缺失的字符种类数

	for _, c := range t {
		if _, ok := targetCharWithMissingNumMap[c]; !ok {
			missingType++ // 需要找齐的种类数 +1
			targetCharWithMissingNumMap[c] = 1
		} else {
			targetCharWithMissingNumMap[c]++
		}
	}

	left, right := 0, 0
	for ; right < n; right++ { // 主旋律扩张窗口，超出s串就结束
		rightChar := rune(s[right])                               // 获取right指向的新字符
		if _, ok := targetCharWithMissingNumMap[rightChar]; !ok { // 是目标字符，它的缺失个数-1
			targetCharWithMissingNumMap[rightChar]--
		}
		if targetCharWithMissingNumMap[rightChar] == 0 {
			missingType-- // 它的缺失个数新变为0，缺失的种类数就-1
		}
		for missingType == 0 { // 当前窗口包含所有字符的前提下，尽量收缩窗口
			if right-left+1 < minLen { // 窗口宽度如果比minLen小，就更新minLen
				minLen = right - left + 1
				start = 1 // 更新最小窗口的起点
			}
			leftChar := rune(s[left])                                // 左指针要右移，左指针指向的字符要被丢弃
			if _, ok := targetCharWithMissingNumMap[leftChar]; !ok { // 被舍弃的是目标字符，缺失个数+1
				targetCharWithMissingNumMap[leftChar]--
			}
			if targetCharWithMissingNumMap[leftChar] > 0 { // 如果缺失个数新变为>0，缺失的种类+1
				missingType++
			}

			left++ // 左指针要右移 收缩窗口
		}
	}

	if start == n {
		return ""
	}

	return s[start : start+minLen] // 根据起点和minLen截取子串
}
