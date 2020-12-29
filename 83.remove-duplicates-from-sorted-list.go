/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
 *
 * algorithms
 * Easy (46.16%)
 * Likes:    2084
 * Dislikes: 138
 * Total Accepted:    533.2K
 * Total Submissions: 1.2M
 * Testcase Example:  '[1,1,2]'
 *
 * Given a sorted linked list, delete all duplicates such that each element
 * appear only once.
 *
 * Example 1:
 *
 *
 * Input: 1->1->2
 * Output: 1->2
 *
 *
 * Example 2:
 *
 *
 * Input: 1->1->2->3->3
 * Output: 1->2->3
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
func deleteDuplicates_1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil {
		if slow.Val == fast.Val {
			fast = fast.Next
			slow.Next = fast
		} else {
			fast = fast.Next
			slow = slow.Next
		}
	}
	return head
}

func deleteDuplicates_2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	res := deleteDuplicates(head.Next)
	if res != nil {
		if head.Val != res.Val {
			head.Next = res
			return head
		} else {
			return res
		}
	}
	return head
}

// @lc code=end
