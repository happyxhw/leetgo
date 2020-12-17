/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
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
func recoverTree(root *TreeNode) {
	var stack []*TreeNode
	var pre, first, second *TreeNode
	p := root

	for p != nil || len(stack) > 0 {
		if p != nil {
			stack = append(stack, p)
			p = p.Left
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pre != nil && top.Val <= pre.Val {
				if first == nil {
					first, second = pre, top
				} else {
					second = top
					break
				}

			}
			p = top.Right
			pre = top
		}
	}
	first.Val, second.Val = second.Val, first.Val
}

// @lc code=end
