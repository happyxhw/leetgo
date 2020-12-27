/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
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
func mergeKLists(lists []*ListNode) *ListNode {
	return helper023(lists, 0, len(lists)-1)
}

func helper023(lists []*ListNode, start, end int) *ListNode {
	if start > end {
		return nil
	}
	if start == end {
		return lists[start]
	}
	mid := (start + end) / 2
	left := helper023(lists, start, mid)
	right := helper023(lists, mid+1, end)
	return merge2(left, right)
}

func merge2(l1 *ListNode, l2 *ListNode) *ListNode {
	preHead := &ListNode{}
	cur := preHead
	for l1 != nil && l2 != nil {
		node := ListNode{Val: l1.Val}
		if l2.Val < l1.Val {
			node.Val = l2.Val
			l2 = l2.Next
		} else {
			l1 = l1.Next
		}
		cur.Next = &node
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}

	return preHead.Next
}

// @lc code=end
