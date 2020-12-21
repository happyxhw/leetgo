/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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
func isBalanced(root *TreeNode) bool {
	_, ret := helper110(root)
	return ret
}

func helper110(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	left, ret := helper110(root.Left)
	if !ret {
		return 0, false
	}
	right, ret := helper110(root.Right)
	if !ret {
		return 0, false
	}
	if left-right < -1 || left-right > 1 {
		return 0, false
	}
	if left > right {
		return left + 1, true
	}
	return right + 1, true
}

// @lc code=end
