/*
20. 有效的括号

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

示例 1：
输入：s = "()"
输出：true

示例 2：
输入：s = "()[]{}"
输出：true

示例 3：
输入：s = "(]"
输出：false

示例 4：
输入：s = "([)]"
输出：false

示例 5：
输入：s = "{[]}"
输出：true

提示：
1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成

链接：https://leetcode-cn.com/problems/valid-parentheses
*/

package Golang

func isValid(s string) bool {
	n := len(s)

	if len(s)%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []byte{}
	for i := 0; i < n; i++ {
		if p, ok := pairs[s[i]]; ok { // 遇到右符号后，开始检查与最后一个入栈的左符号是否对应
			if len(stack) == 0 || stack[len(stack)-1] != p {
				return false
			}
			// 出栈
			stack = stack[:len(stack)-1]
		} else { // 收集左括号入栈
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
