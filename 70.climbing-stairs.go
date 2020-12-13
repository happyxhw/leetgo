/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */
package leetgo

// @lc code=start
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	first, second := 1, 1
	for i := 1; i < n; i++ {
		t := second
		second = first + second
		first = t
	}

	return second
}

// @lc code=end
