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

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if n == 1 || head == nil {
		return head
	}

	res, _, _ := helper092(head, 1, m, n)
	return res
}

func helper092(head *ListNode, j, m, n int) (*ListNode, *ListNode, int) {
	if j == n {
		return head, head.Next, 1
	}
	res, last, c := helper092(head.Next, j+1, m, n)
	if c == n-m+1 {
		head.Next = res
		return head, last, c
	}
	head.Next.Next = head
	head.Next = last

	return res, last, c + 1
}

// @lc code=end
