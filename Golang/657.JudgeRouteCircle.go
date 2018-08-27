/*
Initially, there is a Robot at position (0, 0). Given a sequence of its moves, judge if this robot makes a circle, which means it moves back to the original place.
The move sequence is represented by a string. And each move is represent by a character. The valid robot moves are R (Right), L (Left), U (Up) and D (down). The output should be true or false representing whether the robot makes a circle.

Example 1:
	Input: "UD"
	Output: true
Example 2:
	Input: "LL"
	Output: false
*/

package Golang

// JudgeCircle 判断是否又回到原地
func JudgeCircle(moves string) bool {
	x, y := 0, 0
	for _, m := range moves {
		if m == 'U' {
			y++
		} else if m == 'D' {
			y--
		} else if m == 'L' {
			x--
		} else if m == 'R' {
			x++
		}
	}

	if x == 0 && y == 0 {
		return true
	}
	return false
}
