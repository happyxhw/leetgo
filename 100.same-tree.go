/*
 * @lc app=leetcode id=100 lang=golang
 *
 * [100] Same Tree
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
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return q == nil && p == nil
	}
	if !isSameTree(p.Left, q.Left) {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Right, q.Right)
}

// @lc code=end
