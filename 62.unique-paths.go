/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */
package leetgo

// array | dynamic-programming

// @lc code=start
func uniquePaths(m int, n int) int {
	res := make([]int, n)
	res[0] = 1
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			res[j] = res[j] + res[j-1]
		}
	}
	return res[n-1]
}

// @lc code=end
