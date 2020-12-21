/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
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
func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	var path []int

	res = helper113(root, sum, res, path)
	return res
}

func helper113(root *TreeNode, remain int, res [][]int, path []int) [][]int {
	if root == nil {
		return res
	}
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		if remain == root.Val {
			tmpPath := make([]int, 0, len(path))
			tmpPath = append(tmpPath, path...)
			res = append(res, tmpPath)
		}
		return res
	}
	res = helper113(root.Left, remain-root.Val, res, path)
	res = helper113(root.Right, remain-root.Val, res, path)
	return res
}

// @lc code=end
