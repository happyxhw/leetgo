/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
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
func hasPathSum(root *TreeNode, sum int) bool {
	return helper112(root, sum)
}

func helper112(root *TreeNode, remain int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return remain == root.Val
	}
	if root.Left != nil && helper112(root.Left, remain-root.Val) {
		return true
	}
	if root.Right != nil && helper112(root.Right, remain-root.Val) {
		return true
	}
	return false
}

// @lc code=end
