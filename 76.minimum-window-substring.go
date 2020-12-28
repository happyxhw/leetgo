/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 *
 * https://leetcode.com/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (35.55%)
 * Likes:    5691
 * Dislikes: 388
 * Total Accepted:    472.6K
 * Total Submissions: 1.3M
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * Given two strings s and t, return the minimum window in s which will contain
 * all the characters in t. If there is no such window in s that covers all
 * characters in t, return the empty string "".
 *
 * Note that If there is such a window, it is guaranteed that there will always
 * be only one unique minimum window in s.
 *
 *
 * Example 1:
 * Input: s = "ADOBECODEBANC", t = "ABC"
 * Output: "BANC"
 * Example 2:
 * Input: s = "a", t = "a"
 * Output: "a"
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length, t.length <= 10^5
 * s and t consist of English letters.
 *
 *
 *
 * Follow up: Could you find an algorithm that runs in O(n) time?
 */
package leetgo

// @lc code=start
func minWindow(s string, t string) string {
	var res string
	tmpMap := make(map[byte]int)
	for i := range t {
		tmpMap[t[i]]++
	}

	left, cnt, minLen := 0, 0, 1<<31
	for i := 0; i < len(s); i++ {
		if tmpMap[s[i]] > 0 {
			cnt++
		}
		tmpMap[s[i]]--
		for cnt == len(t) {
			if minLen > i+1-left {
				minLen = i + 1 - left
				res = s[left : left+minLen]
			}
			if tmpMap[s[left]] >= 0 {
				cnt--
			}
			tmpMap[s[left]]++
			left++
		}
	}
	return res
}

// @lc code=end
