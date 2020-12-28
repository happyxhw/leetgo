/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 *
 * https://leetcode.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (34.71%)
 * Likes:    1677
 * Dislikes: 2152
 * Total Accepted:    646.4K
 * Total Submissions: 1.9M
 * Testcase Example:  '4'
 *
 * Given a non-negative integer x, compute and return the square root of x.
 *
 * Since the return type is an integer, the decimal digits are truncated, and
 * only the integer part of the result is returned.
 *
 *
 * Example 1:
 *
 *
 * Input: x = 4
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: x = 8
 * Output: 2
 * Explanation: The square root of 8 is 2.82842..., and since the decimal part
 * is truncated, 2 is returned.
 *
 *
 * Constraints:
 *
 *
 * 0 <= x <= 2^31 - 1
 *
 *
 */
package leetgo

// @lc code=start
func mySqrt(x int) int {
	y := float64(x)
	ans := y
	delta := 0.1

	for ans*ans-y > delta || ans*ans-y < -delta {
		ans = (ans + y/ans) / 2
	}
	return int(ans)
}

func mySqrt_1(x int) int {
	y := float64(x)
	delta := 0.00001

	return helper069(0.0, y, y, delta)
}

func helper069(left, right, x, delta float64) int {
	mid := left + (right-left)/2

	if mid*mid-x < delta && mid*mid-x > -delta {
		return int(mid)
	}

	if mid*mid > x {
		return helper069(left, mid, x, delta)
	} else {
		return helper069(mid, right, x, delta)
	}
}

// @lc code=end
