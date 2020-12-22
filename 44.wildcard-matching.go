/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 */
package leetgo

// @lc code=start
func isMatch0(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = true
		} else {
			break
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = (dp[i-1][j] || dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

// @lc code=end
