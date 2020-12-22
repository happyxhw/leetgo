/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */
package leetgo

// 1 2 3 3 2 1
// @lc code=start
func longestPalindrome(s string) string {
	var maxLen int
	var index int
	for i := range s {
		m := helper005(s, i, i)
		n := helper005(s, i, i+1)
		t := m
		if n > t {
			t = n
		}
		if maxLen < t {
			index = i
			maxLen = t
		}
	}
	start := index - (maxLen-1)/2
	end := index + maxLen/2
	return s[start : end+1]
}

func helper005(s string, i, j int) int {
	x, y := i, j
	for x >= 0 && y < len(s) && s[x] == s[y] {
		x--
		y++
	}
	return y - x - 1
}

// @lc code=end
