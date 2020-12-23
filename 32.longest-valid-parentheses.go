/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */
package leetgo

// Input: s = ")()())"
// Output: 4
// Explanation: The longest valid parentheses substring is "()()".

// @lc code=start 5
func longestValidParentheses(s string) int {
	var maxLen int
	stack := []int{-1}
	for i, c := range s {
		n := len(stack)
		top := stack[n-1]
		if top != -1 && c == ')' && s[top] == '(' {
			stack = stack[:n-1]
			last := stack[len(stack)-1]
			if maxLen < i-last {
				maxLen = i - last
			}
		} else {
			stack = append(stack, i)
		}
	}
	return maxLen
}

// @lc code=end
