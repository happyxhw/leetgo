/*
 * @lc app=leetcode id=115 lang=golang
 *
 * [115] Distinct Subsequences
 *
 * https://leetcode.com/problems/distinct-subsequences/description/
 *
 * algorithms
 * Hard (39.19%)
 * Likes:    1648
 * Dislikes: 62
 * Total Accepted:    151.7K
 * Total Submissions: 386.6K
 * Testcase Example:  '"rabbbit"\n"rabbit"'
 *
 * Given two strings s and t, return the number of distinct subsequences of s
 * which equals t.
 *
 * A string's subsequence is a new string formed from the original string by
 * deleting some (can be none) of the characters without disturbing the
 * relative positions of the remaining characters. (i.e., "ACE" is a
 * subsequence of "ABCDE" while "AEC" is not).
 *
 * It's guaranteed the answer fits on a 32-bit signed integer.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "rabbbit", t = "rabbit"
 * Output: 3
 * Explanation:
 * As shown below, there are 3 ways you can generate "rabbit" from S.
 * rabbbit
 * rabbbit
 * rabbbit
 *
 *
 * Example 2:
 *
 *
 * Input: s = "babgbag", t = "bag"
 * Output: 5
 * Explanation:
 * As shown below, there are 5 ways you can generate "bag" from S.
 * babgbag
 * babgbag
 * babgbag
 * babgbag
 * babgbag
 *
 *
 * Constraints:
 *
 *
 * 0 <= s.length, t.length <= 1000
 * s and t consist of English letters.
 *
 *
 */
package leetgo

// @lc code=start
func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1
	for i := 1; i <= m; i++ {
		dp[i][0] = 1
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[m][n]
}

// @lc code=end
