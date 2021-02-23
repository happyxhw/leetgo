/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
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
func preorderTraversal(root *TreeNode) []int {
	// var res []int
	// var stack []*TreeNode
	// p := root

	// for p != nil || len(stack) > 0 {
	// 	if p != nil {
	// 		res = append(res, p.Val)
	// 		stack = append(stack, p)
	// 		p = p.Left
	// 	} else {
	// 		top := stack[len(stack)-1]
	// 		stack = stack[:len(stack)-1]
	// 		p = top.Right
	// 	}
	// }
	return helper144(root, []int{})
}

func helper144(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = helper144(root.Left, res)
	res = helper144(root.Right, res)
	return res
}

// @lc code=end
