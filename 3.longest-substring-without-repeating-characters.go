/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
package leetgo

func lengthOfLongestSubstring(s string) int {
	indexMap := make(map[rune]int, len(s))
	maxLen, start := 0, 0

	for i, v := range s {
		if index, ok := indexMap[v]; ok && index >= start {
			start = index + 1
		}
		indexMap[v] = i
		maxLen = max(maxLen, i-start+1)
	}
	return maxLen
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end
