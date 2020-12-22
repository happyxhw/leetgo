/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */
package leetgo

// @lc code=start
func generateParenthesis(n int) []string {
	var res []string
	return helper022(res, "", 0, 0, n)
}

func helper022(res []string, str string, left, right, n int) []string {
	if left == n && right == n {
		res = append(res, str)
		return res
	}
	if left < n {
		res = helper022(res, str+"(", left+1, right, n)
	}
	if right < left {
		res = helper022(res, str+")", left, right+1, n)
	}
	return res
}

// @lc code=end
