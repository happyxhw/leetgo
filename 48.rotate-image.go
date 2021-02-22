/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */
package leetgo

// @lc code=start
func rotate_48(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	m := len(matrix)
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			t := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = t
		}
	}
	for i := 0; i < m; i++ {
		left, right := 0, m-1
		for left < right {
			t := matrix[i][left]
			matrix[i][left] = matrix[i][right]
			matrix[i][right] = t
			left++
			right--
		}
	}
}

// @lc code=end
