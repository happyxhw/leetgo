/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 *
 * https://leetcode.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (57.07%)
 * Likes:    1460
 * Dislikes: 125
 * Total Accepted:    234.4K
 * Total Submissions: 409.9K
 * Testcase Example:  '3'
 *
 * Given a positive integer n, generate an n x n matrix filled with elements
 * from 1 to n^2 in spiral order.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 3
 * Output: [[1,2,3],[8,9,4],[7,6,5]]
 *
 *
 * Example 2:
 *
 *
 * Input: n = 1
 * Output: [[1]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 20
 *
 *
 */
package leetgo

// @lc code=start
func generateMatrix(n int) [][]int {
	if n == 0 {
		return nil
	}
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	var rowStart, colStart int
	k := 1
	rowEnd, colEnd := n-1, n-1
	for rowStart <= rowEnd && colStart <= colEnd {
		for i := colStart; i <= colEnd; i++ {
			matrix[rowStart][i] = k
			k++
		}
		rowStart++
		for i := rowStart; i <= rowEnd; i++ {
			matrix[i][colEnd] = k
			k++
		}
		colEnd--
		if rowStart <= rowEnd {
			for i := colEnd; i >= colStart; i-- {
				matrix[rowEnd][i] = k
				k++
			}
			rowEnd--
		}
		if colEnd >= colStart {
			for i := rowEnd; i >= rowStart; i-- {
				matrix[i][colStart] = k
				k++
			}
			colStart++
		}
	}
	return matrix
}

// @lc code=end
