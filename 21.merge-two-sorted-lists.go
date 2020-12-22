/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := &ListNode{}
	cur := preHead
	for l1 != nil && l2 != nil {
		node := ListNode{Val: l1.Val}
		if l2.Val < l1.Val {
			node.Val = l2.Val
			l2 = l2.Next
		} else {
			l1 = l1.Next
		}
		cur.Next = &node
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}

	return preHead.Next
}

// @lc code=end
