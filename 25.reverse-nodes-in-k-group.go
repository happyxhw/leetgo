/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
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

// 使用递归解决
func reverseKGroup1(head *ListNode, k int) *ListNode {
	// 前进k步
	var n int
	cur := head
	for n != k && cur != nil {
		n++
		cur = cur.Next
	}

	if n == k {
		cur = reverseKGroup(cur, k)
		for n > 0 {
			t := head.Next
			head.Next = cur
			cur = head
			head = t
			n--
		}
		head = cur

	}
	return head
}

// 不使用递归
func reverseKGroup(head *ListNode, k int) *ListNode {
	preHead := &ListNode{}
	preHead.Next = head
	pre := preHead
	// 计算节点数量
	var n int
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}
	for n/k > 0 {
		cur = pre.Next
		for i := 1; i < k; i++ {
			t := pre.Next
			pre.Next = cur.Next
			cur.Next = cur.Next.Next
			pre.Next.Next = t
		}
		pre = cur
		n -= k
	}

	return preHead.Next
}

// @lc code=end
