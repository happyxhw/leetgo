/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */
package leetgo

// @lc code=start

func reverse(x int) int {
	var res int
	maxInt32 := 1 << 31
	for x != 0 {
		res = res*10 + x%10
		x = x / 10
		if res > maxInt32 || res < -(maxInt32-1) {
			return 0
		}
	}
	return res
}

// @lc code=end
