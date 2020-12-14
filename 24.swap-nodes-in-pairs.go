/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
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

// 使用递归的方法
func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := swapPairs1(head.Next.Next)
	t := head.Next
	head.Next = cur
	t.Next = head
	head = t
	return head
}

// 不使用递归
func swapPairs2(head *ListNode) *ListNode {
	pre := &ListNode{}
	pre.Next = head
	cur := pre

	for cur.Next != nil && cur.Next.Next != nil {
		t1 := cur.Next
		t2 := cur.Next.Next

		cur.Next.Next = cur
		cur.Next = t2
		cur = t1
	}

	return pre.Next
}

func swapPairs(head *ListNode) *ListNode {
	preHead := &ListNode{}
	preHead.Next = head

	pre, cur := preHead, head

	for cur != nil && cur.Next != nil {
		t := cur.Next
		cur.Next = cur.Next.Next
		t.Next = cur
		pre.Next = t

		pre = cur
		cur = cur.Next
	}

	return preHead.Next
}

// @lc code=end
