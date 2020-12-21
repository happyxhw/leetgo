/*
 * @lc app=leetcode id=109 lang=golang
 *
 * [109] Convert Sorted List to Binary Search Tree
 */
package leetgo

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	return helper109(head, nil)
}

func helper109(first, second *ListNode) *TreeNode {
	if first == second {
		return nil
	}
	mid, cur := first, first
	for cur != second && cur.Next != second {
		cur = cur.Next.Next
		mid = mid.Next
	}
	var left, right *TreeNode
	root := &TreeNode{Val: mid.Val}
	left = helper109(first, mid)
	right = helper109(mid.Next, second)
	root.Left = left
	root.Right = right
	return root
}

// @lc code=end
