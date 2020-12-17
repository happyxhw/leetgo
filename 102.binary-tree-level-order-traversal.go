/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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
func levelOrder1(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		var tmp []int
		var tmpQueue []*TreeNode
		for _, item := range queue {
			tmp = append(tmp, item.Val)
			if item.Left != nil {
				tmpQueue = append(tmpQueue, item.Left)
			}
			if item.Right != nil {
				tmpQueue = append(tmpQueue, item.Right)
			}
		}
		queue = tmpQueue
		res = append(res, tmp)
	}
	return res
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	return helper102(root, res, 1)
}

func helper102(root *TreeNode, res [][]int, depth int) [][]int {
	if root == nil {
		return res
	}
	if len(res) < depth {
		res = append(res, []int{})
	}
	res[depth-1] = append(res[depth-1], root.Val)
	res = helper102(root.Left, res, depth+1)
	res = helper102(root.Right, res, depth+1)
	return res
}

// @lc code=end
