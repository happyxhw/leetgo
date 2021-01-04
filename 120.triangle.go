/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 *
 * https://leetcode.com/problems/triangle/description/
 *
 * algorithms
 * Medium (45.29%)
 * Likes:    2515
 * Dislikes: 292
 * Total Accepted:    277.3K
 * Total Submissions: 611.6K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * Given a triangle array, return the minimum path sum from top to bottom.
 *
 * For each step, you may move to an adjacent number of the row below. More
 * formally, if you are on index i on the current row, you may move to either
 * index i or index i + 1 on the next row.
 *
 *
 * Example 1:
 *
 *
 * Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
 * Output: 11
 * Explanation: The triangle looks like:
 * ⁠  2
 * ⁠ 3 4
 * ⁠6 5 7
 * 4 1 8 3
 * The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined
 * above).
 *
 *
 * Example 2:
 *
 *
 * Input: triangle = [[-10]]
 * Output: -10
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= triangle.length <= 200
 * triangle[0].length == 1
 * triangle[i].length == triangle[i - 1].length + 1
 * -10^4 <= triangle[i][j] <= 10^4
 *
 *
 *
 * Follow up: Could you do this using only O(n) extra space, where n is the
 * total number of rows in the triangle?
 */
package leetgo

// @lc code=start
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	n := len(triangle)
	dp := make([]int, n)
	dp[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		tmp := make([]int, i)
		for k := 0; k < i; k++ {
			tmp[k] = dp[k]
		}
		for j, t := range triangle[i] {
			if j == 0 {
				dp[j] = tmp[0] + t
			} else if j == i {
				dp[j] = tmp[i-1] + t
			} else {
				m := tmp[j-1]
				if m > tmp[j] {
					m = tmp[j]
				}
				dp[j] = m + t
			}
		}
	}
	res := dp[0]
	for _, item := range dp {
		if res > item {
			res = item
		}
	}
	return res
}

// @lc code=end
