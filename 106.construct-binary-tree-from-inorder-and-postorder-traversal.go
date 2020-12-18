/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
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
func buildTree(inorder []int, postorder []int) *TreeNode {

	inorderMap := make(map[int]int, len(inorder))
	for i, item := range inorder {
		inorderMap[item] = i
	}
	return helper106(postorder, inorder, inorderMap, len(postorder)-1, 0, len(inorder)-1)
}

func helper106(postorder []int, inorder []int, inorderMap map[int]int, postStart, inStart, inEnd int) *TreeNode {
	if postStart < 0 || inStart > inEnd {
		return nil
	}
	rootVal := postorder[postStart]
	root := TreeNode{Val: rootVal}
	mid := inorderMap[rootVal]

	left := helper106(postorder, inorder, inorderMap, postStart-inEnd+mid-1, inStart, mid-1)
	right := helper106(postorder, inorder, inorderMap, postStart-1, mid+1, inEnd)

	root.Left, root.Right = left, right

	return &root
}

// @lc code=end
