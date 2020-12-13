/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */
package leetgo

// array | dynamic-programming

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	res := make([]int, n)
	res[0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				res[j] = 0
			} else if j > 0 {
				res[j] = res[j] + res[j-1]
			}
		}
	}
	return res[n-1]
}

// @lc code=end
