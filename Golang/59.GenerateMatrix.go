/*
59. 螺旋矩阵 II

给你一个正整数 n ，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

示例 1：
输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]

示例 2：
输入：n = 1
输出：[[1]]

提示：
1 <= n <= 20

链接：https://leetcode-cn.com/problems/spiral-matrix-ii
*/

package Golang

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	left, right, top, bottom := 0, n-1, 0, n-1
	num, tar := 1, n*n
	for num <= tar {
		for i := left; i <= right; i++ {
			matrix[top][i] = num // left to right.
			num++
		}

		top++

		for i := top; i <= bottom; i++ {
			matrix[i][right] = num // top to bottom.
			num++
		}

		right--

		for i := right; i >= left; i-- {
			matrix[bottom][i] = num // right to left.
			num++
		}

		bottom--

		for i := bottom; i >= top; i-- {
			matrix[i][left] = num // bottom to top.
			num++
		}

		left++
	}

	return matrix
}
