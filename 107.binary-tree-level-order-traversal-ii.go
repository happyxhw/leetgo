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

	return res
}

func helper107(root *TreeNode, res [][]int, depth int) ([][]int, int) {
	if root == nil {
		return res, depth - 1
	}
	if len(res) == depth {
		res[depth] = []int{}
	}

	res, _ = helper107(root.Left, res, depth+1)
	res, _ = helper107(root.Right, res, depth+1)
}

// @lc code=end
