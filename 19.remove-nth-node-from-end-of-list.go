/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	preHead := &ListNode{}
	preHead.Next = head

	slow, fast := preHead, preHead

	for n >= 0 && fast != nil {
		fast = fast.Next
		n--
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	t := slow.Next.Next
	slow.Next.Next = nil
	slow.Next = t
	return preHead.Next
}

// @lc code=end
