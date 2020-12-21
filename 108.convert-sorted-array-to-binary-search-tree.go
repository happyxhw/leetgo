/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
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
func sortedArrayToBST0(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return &TreeNode{Val: nums[0]}
	}
	mid := n / 2
	var left, right *TreeNode
	root := &TreeNode{Val: nums[mid]}
	if mid > 0 {
		left = sortedArrayToBST0(nums[0:mid])
	}
	if mid < n-1 {
		right = sortedArrayToBST0(nums[mid+1 : n])
	}
	root.Left = left
	root.Right = right

	return root
}

// @lc code=end
