/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
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
func maxDepth(root *TreeNode) int {
	return helper104(root)
}

func helper104(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := helper104(root.Left) + 1
	right := helper104(root.Right) + 1

	if left > right {
		return left
	}

	return right
}

// @lc code=end
