/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
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
func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	res = helper107(root, res, 0)
	start, end := 0, len(res)-1
	for start < end {
		res[start], res[end] = res[end], res[start]
		start++
		end--
	}
	return res
}

func helper107(root *TreeNode, res [][]int, depth int) [][]int {
	if root == nil {
		return res
	}
	if len(res) == depth {
		res = append(res, []int{})
	}
	res[depth] = append(res[depth], root.Val)

	res = helper107(root.Left, res, depth+1)
	res = helper107(root.Right, res, depth+1)
	return res
}

// @lc code=end
