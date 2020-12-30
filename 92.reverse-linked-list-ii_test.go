/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 *
 * https://leetcode.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (40.07%)
 * Likes:    3107
 * Dislikes: 164
 * Total Accepted:    314.9K
 * Total Submissions: 785.7K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * Reverse a linked list from position m to n. Do it in one-pass.
 *
 * Note: 1 ≤ m ≤ n ≤ length of list.
 *
 * Example:
 *
 *
 * Input: 1->2->3->4->5->NULL, m = 2, n = 4
 * Output: 1->4->3->2->5->NULL
 *
 *
 */
package leetgo

import (
	"fmt"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	head := &ListNode{Val: 1}
	a := &ListNode{Val: 2}
	head.Next = a
	b := &ListNode{Val: 3}
	a.Next = b
	c := &ListNode{Val: 4}
	b.Next = c
	d := &ListNode{Val: 5}
	c.Next = d

	head = reverseBetween(head, 1, 4)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
