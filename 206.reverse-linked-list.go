/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
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

// 使用递归
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}

// 不使用递归-1
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	preHead := &ListNode{}
	preHead.Next = head

	cur := preHead.Next

	for cur.Next != nil {
		t := preHead.Next
		preHead.Next = cur.Next
		cur.Next = cur.Next.Next
		preHead.Next.Next = t
	}

	return preHead.Next
}

// 不使用递归-2
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode
	for head != nil {
		t := head.Next
		head.Next = pre
		pre = head
		head = t
	}
	return pre
}

// @lc code=end
