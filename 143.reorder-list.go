/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 *
 * https://leetcode.com/problems/reorder-list/description/
 *
 * algorithms
 * Medium (40.56%)
 * Likes:    2885
 * Dislikes: 141
 * Total Accepted:    306.3K
 * Total Submissions: 755K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given a singly linked list L: L0→L1→…→Ln-1→Ln,
 * reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…
 *
 * You may not modify the values in the list's nodes, only nodes itself may be
 * changed.
 *
 * Example 1:
 *
 *
 * Given 1->2->3->4, reorder it to 1->4->2->3.
 *
 * Example 2:
 *
 *
 * Given 1->2->3->4->5, reorder it to 1->5->2->4->3.
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
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	mid, last := head, head

	for last.Next != nil && last.Next.Next != nil {
		mid = mid.Next
		last = last.Next.Next
	}

	var pre *ListNode
	cur := mid.Next

	for cur != nil {
		t := cur.Next
		cur.Next = pre
		pre = cur
		cur = t
	}
	mid.Next = pre

	first := head

	for first != mid {
		mid.Next = pre.Next
		pre.Next = first.Next

		first.Next = pre
		first = pre.Next
		pre = mid.Next
	}
}

// @lc code=end
