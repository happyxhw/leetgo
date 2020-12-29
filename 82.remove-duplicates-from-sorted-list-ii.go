/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (37.85%)
 * Likes:    2244
 * Dislikes: 118
 * Total Accepted:    281.5K
 * Total Submissions: 743.7K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * Given a sorted linked list, delete all nodes that have duplicate numbers,
 * leaving only distinct numbers from the original list.
 *
 * Return the linked list sorted as well.
 *
 * Example 1:
 *
 *
 * Input: 1->2->3->3->4->4->5
 * Output: 1->2->5
 *
 *
 * Example 2:
 *
 *
 * Input: 1->1->1->2->3
 * Output: 2->3
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
func deleteDuplicates(head *ListNode) *ListNode {
	preHead := &ListNode{}
	preHead.Next = head
	slow, fast := preHead, head

	for fast != nil {
		for fast.Next != nil && fast.Val == fast.Next.Val {
			fast = fast.Next
		}
		if slow.Next != fast {
			slow.Next = fast.Next
			fast = fast.Next
		} else {
			fast = fast.Next
			slow = slow.Next
		}
	}
	return preHead.Next
}

// @lc code=end
