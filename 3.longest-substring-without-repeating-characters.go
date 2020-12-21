/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
package leetgo

// @lc code=start

func lengthOfLongestSubstring(s string) int {
	indexMap := make(map[rune]int, len(s))
	var maxLen, start int
	for i, item := range s {
		if v, ok := indexMap[item]; ok && v >= start {
			start = v + 1
		}
		indexMap[item] = i
		if maxLen < i-start+1 {
			maxLen = i - start + 1
		}
	}
	return maxLen
}

// @lc code=end
