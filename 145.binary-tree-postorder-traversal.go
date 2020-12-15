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
	var post *TreeNode

	res := []int{}
	stack := []*TreeNode{}

	p := root

	for p != nil || len(stack) > 0 {
		if p != nil {
			stack = append(stack, p)
			p = p.Left
		} else {
			n := len(stack)
			p = stack[n-1]
			if p.Right != nil && p.Right != post {
				p = p.Right
			} else {
				stack = stack[0 : n-1]
				res = append(res, p.Val)
				post = p
				p = nil
			}
		}
	}

	return res
}

// @lc code=end
