/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */
package leetgo

// @lc code=start
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	// TODO: 边界条件
	dp[0][0] = true
	for j := 2; j <= n; j++ {
		if p[j-1] == '*' && dp[0][j-2] {
			dp[0][j] = true
		}
	}
	// TODO: 状态转移
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if j >= 2 && p[j-1] == '*' {
				if p[j-2] != s[i-1] && p[j-2] != '.' {
					dp[i][j] = dp[i][j-2]
				} else {
					dp[i][j] = (dp[i-1][j-2] || dp[i-1][j] || dp[i][j-2])
				}
			}
		}
	}

	return dp[m][n]
}

// @lc code=end
