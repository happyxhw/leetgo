/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */
package leetgo

// @lc code=start
func isValid(s string) bool {
	stack := make([]rune, 0, len(s))
	for _, c := range s {
		switch c {
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		case '(':
			stack = append(stack, ')')
		default:
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top != c {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// @lc code=end
