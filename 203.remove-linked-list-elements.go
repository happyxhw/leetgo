/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 *
 * https://leetcode.com/problems/remove-linked-list-elements/description/
 *
 * algorithms
 * Easy (39.18%)
 * Likes:    2453
 * Dislikes: 119
 * Total Accepted:    436.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,6,3,4,5,6]\n6'
 *
 * Remove all elements from a linked list of integers that have value val.
 *
 * Example:
 *
 *
 * Input:  1->2->6->3->4->5->6, val = 6
 * Output: 1->2->3->4->5
 *
 *
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
func removeElements(head *ListNode, val int) *ListNode {
	pre := &ListNode{Val: -1}
	pre.Next = head

	res := pre

	for pre.Next != nil {
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
			continue
		}
		pre = pre.Next
	}

	return res.Next
}

// @lc code=end
