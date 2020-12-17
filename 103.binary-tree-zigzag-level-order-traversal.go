/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	return helper103(root, res, 1)
}

func helper103(root *TreeNode, res [][]int, depth int) [][]int {
	if root == nil {
		return res
	}
	if len(res) < depth {
		res = append(res, []int{})
	}
	if depth%2 != 0 {
		res[depth-1] = append(res[depth-1], root.Val)
	} else {
		res[depth-1] = append([]int{root.Val}, res[depth-1]...)
	}

	res = helper103(root.Left, res, depth+1)
	res = helper103(root.Right, res, depth+1)

	return res
}

// @lc code=end
