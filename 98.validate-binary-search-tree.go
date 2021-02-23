/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 */
package leetgo

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 利用二叉树的中序遍历解决
// 栈
func isValidBST0(root *TreeNode) bool {
	var stack []*TreeNode
	var pre *TreeNode
	p := root

	for p != nil || len(stack) > 0 {
		if p != nil {
			stack = append(stack, p)
			p = p.Left
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pre != nil && top.Val <= pre.Val {
				return false
			}
			p = top.Right
			pre = top
		}
	}
	return true
}

// 递归
func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	_, ok := helperValidBST(root, pre)
	return ok
}

func helperValidBST(root, pre *TreeNode) (*TreeNode, bool) {
	if root == nil {
		return pre, true
	}
	var ok bool
	if pre, ok = helperValidBST(root.Left, pre); !ok {
		return pre, false
	}
	if pre != nil && root.Val <= pre.Val {
		return pre, false
	}
	pre = root
	if pre, ok = helperValidBST(root.Right, pre); !ok {
		return pre, false
	}
	return pre, true
}

// @lc code=end
