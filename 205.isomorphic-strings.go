/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 *
 * https://leetcode.com/problems/isomorphic-strings/description/
 *
 * algorithms
 * Easy (40.47%)
 * Likes:    1848
 * Dislikes: 449
 * Total Accepted:    344.1K
 * Total Submissions: 850.3K
 * Testcase Example:  '"egg"\n"add"'
 *
 * Given two strings s and t, determine if they are isomorphic.
 *
 * Two strings s and t are isomorphic if the characters in s can be replaced to
 * get t.
 *
 * All occurrences of a character must be replaced with another character while
 * preserving the order of characters. No two characters may map to the same
 * character, but a character may map to itself.
 *
 *
 * Example 1:
 * Input: s = "egg", t = "add"
 * Output: true
 * Example 2:
 * Input: s = "foo", t = "bar"
 * Output: false
 * Example 3:
 * Input: s = "paper", t = "title"
 * Output: true
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 5 * 10^4
 * t.length == s.length
 * s and t consist of any valid ascii character.
 *
 *
 */
package leetgo

// @lc code=start
func isIsomorphic(s string, t string) bool {
	d1, d2 := make(map[byte]int), make(map[byte]int)

	for i := range s {
		if _, ok := d1[s[i]]; !ok {
			d1[s[i]] = i + 1
		}
		if _, ok := d2[t[i]]; !ok {
			d2[t[i]] = i + 1
		}

		if d1[s[i]] != d2[t[i]] {
			return false
		}
	}

	return true
}

// @lc code=end
