/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 *
 * https://leetcode.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (46.47%)
 * Likes:    2376
 * Dislikes: 312
 * Total Accepted:    544.6K
 * Total Submissions: 1.2M
 * Testcase Example:  '"11"\n"1"'
 *
 * Given two binary strings a and b, return their sum as a binary string.
 *
 *
 * Example 1:
 * Input: a = "11", b = "1"
 * Output: "100"
 * Example 2:
 * Input: a = "1010", b = "1011"
 * Output: "10101"
 *
 *
 * Constraints:
 *
 *
 * 1 <= a.length, b.length <= 10^4
 * a and b consist only of '0' or '1' characters.
 * Each string does not contain leading zeros except for the zero itself.
 *
 *
 */
package leetgo

import "strconv"

// @lc code=start
func addBinary(a string, b string) string {
	res := ""

	i, j := len(a)-1, len(b)-1
	var carry uint8
	for i >= 0 || j >= 0 {
		var first, second uint8
		if i >= 0 {
			first = a[i] - uint8(48)
		}
		if j >= 0 {
			second = b[j] - uint8(48)
		}

		sum := carry + first + second
		carry = sum / 2

		res = strconv.Itoa(int(sum%2)) + res

		if i >= 0 {
			i--
		}
		if j >= 0 {
			j--
		}
	}
	if carry > 0 {
		res = "1" + res
	}
	return res
}

// @lc code=end
