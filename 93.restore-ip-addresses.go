/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 *
 * https://leetcode.com/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (37.03%)
 * Likes:    1570
 * Dislikes: 542
 * Total Accepted:    219K
 * Total Submissions: 591.1K
 * Testcase Example:  '"25525511135"'
 *
 * Given a string s containing only digits, return all possible valid IP
 * addresses that can be obtained from s. You can return them in any order.
 *
 * A valid IP address consists of exactly four integers, each integer is
 * between 0 and 255, separated by single dots and cannot have leading zeros.
 * For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses and
 * "0.011.255.245", "192.168.1.312" and "192.168@1.1" are invalid IP
 * addresses.
 *
 *
 * Example 1:
 * Input: s = "25525511135"
 * Output: ["255.255.11.135","255.255.111.35"]
 * Example 2:
 * Input: s = "0000"
 * Output: ["0.0.0.0"]
 * Example 3:
 * Input: s = "1111"
 * Output: ["1.1.1.1"]
 * Example 4:
 * Input: s = "010010"
 * Output: ["0.10.0.10","0.100.1.0"]
 * Example 5:
 * Input: s = "101023"
 * Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
 *
 *
 * Constraints:
 *
 *
 * 0 <= s.length <= 3000
 * s consists of digits only.
 *
 *
 */

package leetgo

import "strconv"

// @lc code=start
func restoreIpAddresses(s string) []string {
	return helper093(s, "", 4, []string{})
}

func helper093(s, tmp string, k int, res []string) []string {
	if k == 0 {
		if s == "" {
			res = append(res, tmp)
		} else {
			return res
		}
	} else {
		for i := 1; i <= 3; i++ {
			if len(s) >= i && isValidIp(s[:i]) {
				if k == 1 {
					res = helper093(s[i:], tmp+s[:i], k-1, res)
				} else {
					res = helper093(s[i:], tmp+s[:i]+".", k-1, res)
				}
			}
		}
	}
	return res
}

func isValidIp(s string) bool {
	if len(s) == 0 || len(s) > 3 || (len(s) > 1 && s[0] == '0') {
		return false
	}
	t, _ := strconv.ParseInt(s, 10, 64)
	res := int(t)
	return res <= 255 && res >= 0
}

// @lc code=end
