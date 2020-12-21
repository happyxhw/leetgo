/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	cur := pre
	var c int
	for l1 != nil || l2 != nil {
		s := 0
		if l1 != nil {
			s = l1.Val
			l1 = l1.Next
		}
		t := 0
		if l2 != nil {
			t = l2.Val
			l2 = l2.Next
		}
		sum := s + t + c
		c = sum / 10
		node := ListNode{Val: sum % 10}
		cur.Next = &node
		cur = cur.Next
	}
	if c > 0 {
		cur.Next = &ListNode{Val: 1}
	}
	return pre.Next
}

// @lc code=end
