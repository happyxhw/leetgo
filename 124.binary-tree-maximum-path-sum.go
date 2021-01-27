/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
 *
 * https://leetcode.com/problems/binary-tree-maximum-path-sum/description/
 *
 * algorithms
 * Hard (35.23%)
 * Likes:    4931
 * Dislikes: 364
 * Total Accepted:    453.4K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,2,3]'
 *
 * A path in a binary tree is a sequence of nodes where each pair of adjacent
 * nodes in the sequence has an edge connecting them. A node can only appear in
 * the sequence at most once. Note that the path does not need to pass through
 * the root.
 * 
 * The path sum of a path is the sum of the node's values in the path.
 * 
 * Given the root of a binary tree, return the maximum path sum of any path.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: root = [1,2,3]
 * Output: 6
 * Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 =
 * 6.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: root = [-10,9,20,null,null,15,7]
 * Output: 42
 * Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 +
 * 7 = 42.
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * The number of nodes in the tree is in the range [1, 3 * 10^4].
 * -1000 <= Node.val <= 1000
 * 
 * 
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
func maxPathSum(root *TreeNode) int {
	maxVal := -(1 << 31)
	_, maxVal = helper124(root, maxVal)
	return maxVal
}

func helper124(root *TreeNode, maxVal int) (int, int) {
	if root == nil {
		return 0, maxVal
	}
	maxLeft, maxVal := helper124(root.Left, maxVal)
	if maxLeft < 0 {
		maxLeft = 0
	}
	maxRight, maxVal := helper124(root.Right, maxVal)
	if maxRight < 0 {
		maxRight = 0
	}
	if root.Val + maxLeft + maxRight > maxVal {
		maxVal = root.Val + maxLeft + maxRight
	}
	if maxLeft > maxRight {
		return maxLeft+root.Val, maxVal
	}
	return maxRight+root.Val, maxVal
}

// @lc code=end

