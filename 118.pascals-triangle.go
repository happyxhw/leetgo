/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 *
 * https://leetcode.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (54.15%)
 * Likes:    2074
 * Dislikes: 123
 * Total Accepted:    443.5K
 * Total Submissions: 817.9K
 * Testcase Example:  '5'
 *
 * Given a non-negative integer numRows, generate the first numRows of Pascal's
 * triangle.
 *
 *
 * In Pascal's triangle, each number is the sum of the two numbers directly
 * above it.
 *
 * Example:
 *
 *
 * Input: 5
 * Output:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1]
 * ]
 *
 *
 */
package leetgo

// @lc code=start
func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	res := make([][]int, numRows)
	res[0] = []int{1}
	for i := 1; i < numRows; i++ {
		t := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				t[j] = 1
			} else {
				t[j] = res[i-1][j-1] + res[i-1][j]
			}
		}

		res[i] = t
	}
	return res
}

// @lc code=end
