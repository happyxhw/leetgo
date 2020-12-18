/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */
package leetgo

// @lc code=start
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	cur, pre := make([]int, n+1), make([]int, n+1)
	for j := 1; j < n+1; j++ {
		pre[j] = j
	}
	for i := 1; i <= m; i++ {
		cur[0] = i
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				cur[j] = pre[j-1]
			} else {
				cur[j] = min(pre[j-1], min(pre[j], cur[j-1])) + 1
			}
		}
		for i := range pre {
			pre[i] = 0
		}
		pre, cur = cur, pre
	}

	return pre[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
