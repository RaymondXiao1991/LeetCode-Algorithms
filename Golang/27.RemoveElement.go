/*Given an array and a value, remove all instances of that > value in place and

 *return the new length.

 *The order of elements can be changed. It doesn't matter what you leave

 *beyond the new length.
 */

package Golang

// RemoveElement 移除元素
func RemoveElement(arr []int, elem int) int {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == elem {
			continue
		}
		arr[j] = arr[i]
		j++
	}
	return j
}
