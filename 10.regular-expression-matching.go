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
	/*
		'*' 不会单独出现，只会出现 .* 或 x* 或 a*b*cc*
		空字符串匹配 p：j >= 2 && p[j-1] == '*' && dp[0][j-2]
	*/
	for j := 2; j <= n; j += 2 {
		if p[j-1] == '*' && dp[0][j-2] {
			dp[0][j] = true
		} else {
			break
		}
	}
	// TODO: 状态转移
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// xxxxxxA -> yyyyyA || yyyyy.
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
				// xxxxxBA -> yyyyy*
			} else if j >= 2 && p[j-1] == '*' {
				switch {
				// xxxxxBA -> yyyyyA* || xxxxxBA -> yyyyy.*:
				// [xxxxxB]A -> [yyyyy]A*: dp[i-1][j-2]  *=0
				// [xxxxxB]A -> [yyyyyA*]A: dp[i-1][j-2] *=1
				// [xxxxxBA] -> [yyyyy]A*: dp[i-1][j-2]  *=0
				case p[j-2] == s[i-1] || p[j-2] == '.':
					dp[i][j] = dp[i-1][j-2] || dp[i-1][j] || dp[i][j-2] || dp[i][j-1]
				// xxxxxBA -> yyyyyC*: dp[i][j-2]
				default:
					dp[i][j] = dp[i][j-2]
				}
			} else {
				dp[i][j] = false
			}
		}
	}

	return dp[m][n]
}

// @lc code=end
