/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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

//  preorder = [3,9,20,15,7]
//  inorder = [9,3,15,20,7]
func buildTree1(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int, len(inorder))
	for i, item := range inorder {
		inorderMap[item] = i
	}
	return helper105(preorder, inorder, inorderMap, 0, 0, len(inorder)-1)
}

func helper105(preorder []int, inorder []int, inorderMap map[int]int, preStart, inStart, inEnd int) *TreeNode {
	if preStart >= len(preorder) || inStart > inEnd {
		return nil
	}
	rootVal := preorder[preStart]
	root := TreeNode{Val: rootVal}
	mid := inorderMap[rootVal]

	left := helper105(preorder, inorder, inorderMap, preStart+1, inStart, mid-1)
	right := helper105(preorder, inorder, inorderMap, preStart+mid-inStart+1, mid+1, inEnd)

	root.Left, root.Right = left, right

	return &root
}

// @lc code=end
