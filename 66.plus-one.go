/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 *
 * https://leetcode.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (42.64%)
 * Likes:    1971
 * Dislikes: 2832
 * Total Accepted:    744.4K
 * Total Submissions: 1.7M
 * Testcase Example:  '[1,2,3]'
 *
 * Given a non-empty array of decimal digits representing a non-negative
 * integer, increment one to the integer.
 *
 * The digits are stored such that the most significant digit is at the head of
 * the list, and each element in the array contains a single digit.
 *
 * You may assume the integer does not contain any leading zero, except the
 * number 0 itself.
 *
 *
 * Example 1:
 *
 *
 * Input: digits = [1,2,3]
 * Output: [1,2,4]
 * Explanation: The array represents the integer 123.
 *
 *
 * Example 2:
 *
 *
 * Input: digits = [4,3,2,1]
 * Output: [4,3,2,2]
 * Explanation: The array represents the integer 4321.
 *
 *
 * Example 3:
 *
 *
 * Input: digits = [0]
 * Output: [1]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= digits.length <= 100
 * 0 <= digits[i] <= 9
 *
 *
 */
package leetgo

// @lc code=start
func plusOne(digits []int) []int {
	var c, i int
	for i = len(digits) - 1; i >= 0; i-- {
		var t int
		if i == len(digits)-1 {
			t = digits[i] + 1
		} else {
			t = digits[i] + c
		}

		digits[i] = t % 10
		c = t / 10
		if c == 0 {
			break
		}
	}
	if c > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

// @lc code=end
