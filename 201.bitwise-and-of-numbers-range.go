/*
 * @lc app=leetcode id=201 lang=golang
 *
 * [201] Bitwise AND of Numbers Range
 *
 * https://leetcode.com/problems/bitwise-and-of-numbers-range/description/
 *
 * algorithms
 * Medium (39.62%)
 * Likes:    1282
 * Dislikes: 141
 * Total Accepted:    167.4K
 * Total Submissions: 422.4K
 * Testcase Example:  '5\n7'
 *
 * Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND
 * of all numbers in this range, inclusive.
 *
 * Example 1:
 *
 *
 * Input: [5,7]
 * Output: 4
 *
 *
 * Example 2:
 *
 *
 * Input: [0,1]
 * Output: 0
 */
package leetgo

// @lc code=start
func rangeBitwiseAnd(m int, n int) int {
	var i uint32

	for m != n {
		m = m >> 1
		n = n >> 1
		i++
	}

	return m << i
}

// @lc code=end
