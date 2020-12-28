/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 *
 * https://leetcode.com/problems/set-matrix-zeroes/description/
 *
 * algorithms
 * Medium (43.95%)
 * Likes:    2888
 * Dislikes: 344
 * Total Accepted:    378.3K
 * Total Submissions: 860.6K
 * Testcase Example:  '[[1,1,1],[1,0,1],[1,1,1]]'
 *
 * Given an m x n matrix. If an element is 0, set its entire row and column to
 * 0. Do it in-place.
 *
 * Follow up:
 *
 *
 * A straight forward solution using O(mn) space is probably a bad idea.
 * A simple improvement uses O(m + n) space, but still not the best
 * solution.
 * Could you devise a constant space solution?
 *
 *
 *
 * Example 1:
 *
 *
 * Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
 * Output: [[1,0,1],[0,0,0],[1,0,1]]
 *
 *
 * Example 2:
 *
 *
 * Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
 * Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
 *
 *
 *
 * Constraints:
 *
 *
 * m == matrix.length
 * n == matrix[0].length
 * 1 <= m, n <= 200
 * -2^31 <= matrix[i][j] <= 2^31 - 1
 *
 *
 */
package leetgo

// @lc code=start
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	rowMap := make(map[int]bool)
	colMap := make(map[int]bool)

	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rowMap[i] = true
				colMap[j] = true
			}
		}
	}
	for k := range rowMap {
		for i := 0; i < n; i++ {
			matrix[k][i] = 0
		}
	}
	for k := range colMap {
		for i := 0; i < m; i++ {
			matrix[i][k] = 0
		}
	}
}

// @lc code=end
