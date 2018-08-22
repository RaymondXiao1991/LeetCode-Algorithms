/*
Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.
Example 1:
	Input: "Hello"
	Output: "hello"

Example 2:
	Input: "here"
	Output: "here"

Example 3:
	Input: "LOVELY"
	Output: "lovely"
*/

package Golang

// ToLowerCase 转小写
func ToLowerCase(str string) string {
	s := []rune(str)
	r := []rune{}
	for i := 0; i < len(s); i++ {
		if 'A' <= s[i] && s[i] <= 'Z' {
			s[i] += 32
		}
		r = append(r, s[i])
	}
	return (string(s))
}
