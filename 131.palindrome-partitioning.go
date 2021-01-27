/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 *
 * https://leetcode.com/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (51.27%)
 * Likes:    2902
 * Dislikes: 93
 * Total Accepted:    288.3K
 * Total Submissions: 559.8K
 * Testcase Example:  '"aab"'
 *
 * Given a string s, partition s such that every substring of the partition is
 * a palindrome. Return all possible palindrome partitioning of s.
 *
 * A palindrome string is a string that reads the same backward as forward.
 *
 *
 * Example 1:
 * Input: s = "aab"
 * Output: [["a","a","b"],["aa","b"]]
 * Example 2:
 * Input: s = "a"
 * Output: [["a"]]
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 16
 * s contains only lowercase English letters.
 *
 *
 */
package leetgo

// @lc code=start
func partition(s string) [][]string {
	var res [][]string
	return helper131(s, res, []string{})
}

func helper131(s string, res [][]string, path []string) [][]string {
	if len(s) == 0 {
		if len(path) > 0 {
			tmpPath := make([]string, 0, len(path))
			tmpPath = append(tmpPath, path...)
			res = append(res, tmpPath)
			return res
		}
	}
	for i := 1; i <= len(s); i++ {
		if isValidPalindrome(s[:i]) {
			path = append(path, s[:i])
			res = helper131(s[i:], res, path)
			path = path[:len(path)-1]
		}
	}
	return res
}

func isValidPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// @lc code=end
