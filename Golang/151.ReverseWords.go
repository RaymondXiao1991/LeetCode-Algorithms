/*
151. 翻转字符串里的单词

给你一个字符串 s ，逐个翻转字符串中的所有 单词 。
单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
请你返回一个翻转 s 中单词顺序并用单个空格相连的字符串。

说明：
输入字符串 s 可以在前面、后面或者单词间包含多余的空格。
翻转后单词间应当仅用一个空格分隔。
翻转后的字符串中不应包含额外的空格。

示例 1：
输入：s = "the sky is blue"
输出："blue is sky the"

示例 2：
输入：s = "  hello world  "
输出："world hello"
解释：输入字符串可以在前面或者后面包含多余的空格，但是翻转后的字符不能包括。

示例 3：
输入：s = "a good   example"
输出："example good a"
解释：如果两个单词间有多余的空格，将翻转后单词间的空格减少到只含一个。

示例 4：
输入：s = "  Bob    Loves  Alice   "
输出："Alice Loves Bob"

示例 5：
输入：s = "Alice does not even like bob"
输出："bob like even not does Alice"

提示：
1 <= s.length <= 10^4
s 包含英文大小写字母、数字和空格 ' '
s 中 至少存在一个 单词

进阶：
请尝试使用 O(1) 额外空间复杂度的原地解法。

链接：https://leetcode-cn.com/problems/reverse-words-in-a-string
*/

package Golang

func reverseWords(s string) string {
	// 1.使用双指针删除冗余的空格
	slowIndex, fastIndex := 0, 0
	b := []byte(s)
	// 删除头部冗余空格
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}
	// 删除单词间冗余空格
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}
	// 删除尾部冗余空格
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}

	// 2.反转整个字符串
	reverseStrWithBorder(&b, 0, len(b)-1)

	// 3.反转单个单词  i单词开始位置，j单词结束位置
	i := 0
	for i < len(b) {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		reverseStrWithBorder(&b, i, j-1)
		i = j
		i++
	}
	return string(b)
}

func reverseStrWithBorder(b *[]byte, left, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}

// 分割
func reverseWords3(s string) string {
	var res string
	sbyte := []byte(s)
	var tmp []byte
	for _, v := range sbyte {
		if v == ' ' {
			if len(tmp) == 0 {
				continue
			}

			if res == "" {
				res = string(tmp)
			} else {
				res = string(tmp) + " " + res
			}
			tmp = []byte{}
		} else {
			tmp = append(tmp, v)
		}
	}

	if len(tmp) != 0 {
		if res == "" {
			res = string(tmp)
		} else {
			res = string(tmp) + " " + res
		}
	}

	return res
}

// 双指针
func reverseWords4(s string) string {
	n := len(s)

	// 确定好begin和end，分别为去掉首尾空格后的字符串
	begin, end := 0, n-1
	for begin < n && s[begin] == ' ' {
		begin++
	}
	for end >= 0 && s[end] == ' ' {
		end--
	}

	// 开始用双指针，双指针从后开始，分别指向一个单词的首尾，然后添加进结果res
	var res string
	k := end
	sbyte := []byte(s)
	for k >= begin {
		for k >= begin && sbyte[k] != ' ' {
			k--
		}
		res += string(sbyte[k+1:end+1]) + " "
		for k >= begin && sbyte[k] == ' ' {
			k--
		}
		end = k
	}
	if len(res) > 0 {
		return res[:len(res)-1]
	}

	return res
}
