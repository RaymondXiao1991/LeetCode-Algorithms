/*
The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

Given two integers x and y, calculate the Hamming distance.

Note:
0 ≤ x, y < 2(31).

Example:
	Input: x = 1, y = 4
	Output: 2
Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑
The above arrows point to positions where the corresponding bits are different.
*/

package Golang

// HammingDistance 两数比特位不同的个数
func HammingDistance(x int, y int) int {
	z := x ^ y
	c := 0
	for c = 0; z > 0; z >>= 1 {
		c += z & 1
	}
	return c
}
