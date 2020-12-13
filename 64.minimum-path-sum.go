/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */
package leetgo

// @lc code=start
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := make([]int, n)
	res[0] = grid[0][0]
	for j := 1; j < n; j++ {
		res[j] = res[j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				res[j] += grid[i][j]
				continue
			}
			if res[j] > res[j-1] {
				res[j] = res[j-1] + grid[i][j]
			} else {
				res[j] += grid[i][j]
			}
		}
	}
	return res[len(res)-1]
}

// @lc code=end
