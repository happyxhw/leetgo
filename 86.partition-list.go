/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 *
 * https://leetcode.com/problems/partition-list/description/
 *
 * algorithms
 * Medium (42.81%)
 * Likes:    1715
 * Dislikes: 348
 * Total Accepted:    240.5K
 * Total Submissions: 561.7K
 * Testcase Example:  '[1,4,3,2,5,2]\n3'
 *
 * Given a linked list and a value x, partition it such that all nodes less
 * than x come before nodes greater than or equal to x.
 *
 * You should preserve the original relative order of the nodes in each of the
 * two partitions.
 *
 * Example:
 *
 *
 * Input: head = 1->4->3->2->5->2, x = 3
 * Output: 1->2->2->4->3->5
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
func partition_0(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	preHead := &ListNode{}
	preHead.Next = head
	slow, fast := preHead, preHead
	for fast != nil && fast.Next != nil {
		if fast.Next.Val < x {
			if fast != slow {
				t := fast.Next
				fast.Next = fast.Next.Next
				t.Next = slow.Next
				slow.Next = t
			} else {
				fast = fast.Next
			}
			slow = slow.Next
		} else {
			fast = fast.Next
		}

	}
	return preHead.Next
}

// @lc code=end
