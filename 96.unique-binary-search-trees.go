/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */
package leetgo

// 动态规划 和 二叉树
// @lc code=start
func numTrees(n int) int {
	res := make([]int, n+1)
	res[0], res[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			res[i] += res[j] * res[i-j-1]
		}
	}

	return res[n]
}

// @lc code=end

// res[n] = res[m] + res[n]
