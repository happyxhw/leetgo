/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 *
 * https://leetcode.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (37.16%)
 * Likes:    2664
 * Dislikes: 186
 * Total Accepted:    400.2K
 * Total Submissions: 1.1M
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n3'
 *
 * Write an efficient algorithm that searches for a value in an m x n matrix.
 * This matrix has the following properties:
 *
 *
 * Integers in each row are sorted from left to right.
 * The first integer of each row is greater than the last integer of the
 * previous row.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 100
 * -10^4 <= matrix[i][j], target <= 10^4
 *
 *
 */
package leetgo

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	return helper074(matrix, target)
}

func helper074(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	rowStart, rowEnd, colStart, colEnd := 0, m-1, 0, n-1
	for rowStart <= rowEnd && colStart <= colEnd {
		rowMid := (rowStart + rowEnd) / 2
		switch {
		case matrix[rowMid][0] > target:
			rowEnd = rowMid - 1
		case matrix[rowMid][n-1] < target:
			rowStart = rowMid + 1
		case matrix[rowMid][0] <= target && matrix[rowMid][n-1] >= target:
			colMid := (colStart + colEnd) / 2
			if matrix[rowMid][colMid] == target {
				return true
			} else if matrix[rowMid][colMid] > target {
				colEnd = colMid - 1
			} else {
				colStart = colMid + 1
			}
		}
	}
	return false
}

// @lc code=end
