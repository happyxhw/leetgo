/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
 */
package leetgo

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	var res []*TreeNode
	if start == end {
		node := TreeNode{Val: start}
		res = append(res, &node)
		return res
	}
	for k := start; k <= end; k++ {
		leftList := helper(start, k-1)
		rightList := helper(k+1, end)
		for _, left := range leftList {
			for _, right := range rightList {
				root := TreeNode{Val: k}
				root.Left = left
				root.Right = right
				res = append(res, &root)
			}
		}
	}
	return res
}

// @lc code=end
