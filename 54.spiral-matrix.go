/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 *
 * https://leetcode.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (35.14%)
 * Likes:    3224
 * Dislikes: 627
 * Total Accepted:    440.4K
 * Total Submissions: 1.2M
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * Given an m x n matrix, return all elements of the matrix in spiral order.
 *
 *
 * Example 1:
 *
 *
 * Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * Output: [1,2,3,6,9,8,7,4,5]
 *
 *
 * Example 2:
 *
 *
 * Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
 * Output: [1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 *
 * Constraints:
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 10
 * -100 <= matrix[i][j] <= 100
 *
 *
 */
package leetgo

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	rowEnd, colEnd := len(matrix)-1, len(matrix[0])-1
	res := make([]int, 0, (rowEnd+1)*(colEnd+1))
	var rowStart, colStart int

	for rowStart <= rowEnd && colStart <= colEnd {
		for i := colStart; i <= colEnd; i++ {
			res = append(res, matrix[rowStart][i])
		}
		rowStart++
		for i := rowStart; i <= rowEnd; i++ {
			res = append(res, matrix[i][colEnd])
		}
		colEnd--
		if rowStart <= rowEnd {
			for i := colEnd; i >= colStart; i-- {
				res = append(res, matrix[rowEnd][i])
			}
			rowEnd--
		}
		if colEnd >= colStart {
			for i := rowEnd; i >= rowStart; i-- {
				res = append(res, matrix[i][colStart])
			}
			colStart++
		}
	}
	return res
}

// @lc code=end
