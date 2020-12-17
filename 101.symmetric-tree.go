/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper101(root.Left, root.Right)
}

func helper101(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == nil && right == nil
	}
	if left.Val != right.Val {
		return false
	}
	if !helper101(left.Left, right.Right) {
		return false
	}
	return helper101(left.Right, right.Left)
}

// @lc code=end
