/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 *
 * https://leetcode.com/problems/multiply-strings/description/
 *
 * algorithms
 * Medium (34.57%)
 * Likes:    2120
 * Dislikes: 911
 * Total Accepted:    335.4K
 * Total Submissions: 968.7K
 * Testcase Example:  '"2"\n"3"'
 *
 * Given two non-negative integers num1 and num2 represented as strings, return
 * the product of num1 and num2, also represented as a string.
 *
 * Note: You must not use any built-in BigInteger library or convert the inputs
 * to integer directly.
 *
 *
 * Example 1:
 * Input: num1 = "2", num2 = "3"
 * Output: "6"
 * Example 2:
 * Input: num1 = "123", num2 = "456"
 * Output: "56088"
 *
 *
 * Constraints:
 *
 *
 * 1 <= num1.length, num2.length <= 200
 * num1 and num2 consist of digits only.
 * Both num1 and num2 do not contain any leading zero, except the number 0
 * itself.
 *
 *
 */
package leetgo

// @lc code=start
func multiply(num1 string, num2 string) string {
	sums := make([]byte, len(num1)+len(num2))
	for i := range sums {
		sums[i] = '0'
	}
	for j := len(num2) - 1; j >= 0; j-- {
		var c byte
		for i := len(num1) - 1; i >= 0; i-- {
			t := (sums[i+j+1] - '0') + (num1[i]-'0')*(num2[j]-'0') + c
			sums[i+j+1] = t%10 + '0'
			c = t / 10
		}
		sums[j] += c
	}
	for i := range sums {
		if sums[i] != '0' {
			return string(sums[i:])
		}
	}
	return "0"
}

// @lc code=end
