/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */
package leetgo

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var res int
	y := x
	for y != 0 {
		res = res*10 + y%10
		y = y / 10
	}

	return res == x
}

// @lc code=end
