/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
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
func flatten(root *TreeNode) {
	_, _ = helper114(root, root)
}

func helper114(root, tail *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, tail
	}
	if root.Left == nil && root.Right == nil {
		return root, root
	}
	left, leftTail := helper114(root.Left, tail)
	right, rightTail := helper114(root.Right, tail)
	if left != nil {
		root.Right = left
		leftTail.Right = right
	}
	tail = rightTail
	if right == nil {
		tail = leftTail
	}
	root.Left = nil
	return root, tail
}

// @lc code=end
