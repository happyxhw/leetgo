/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
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
func postorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	p := root

	for p != nil || len(stack) > 0 {
		if p != nil {
			res = append(res, p.Val)
			stack = append(stack, p)
			p = p.Right
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = top.Left
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

// @lc code=end
