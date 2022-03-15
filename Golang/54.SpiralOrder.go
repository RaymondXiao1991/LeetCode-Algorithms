/*
54. 螺旋矩阵

给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

示例 1：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

示例 2：
输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

提示：
m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100

链接：https://leetcode-cn.com/problems/spiral-matrix
*/

package Golang

func spiralOrder(matrix [][]int) []int {
	order := []int{}

	// 每一圈
	for i := 0; i < len(matrix[0])/2+1; i++ {
		// 上
		for top := i; top < len(matrix[0])-i-1 && len(matrix)-1-i >= i; top++ {
			order = append(order, matrix[i][top])
		}
		// 右
		for right := i; right < len(matrix)-i && i <= len(matrix[0])-1-i; right++ {
			order = append(order, matrix[right][len(matrix[0])-i-1])
		}
		// 下
		for bottom := len(matrix[0]) - i - 2; bottom > i && len(matrix)-1-i > i; bottom-- {
			order = append(order, matrix[len(matrix)-1-i][bottom])
		}
		// 左
		for left := len(matrix) - 1 - i; left > i && i < len(matrix[0])-1-i; left-- {
			order = append(order, matrix[left][i])
		}
	}

	return order
}
