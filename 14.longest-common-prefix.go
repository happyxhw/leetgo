/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */
package leetgo

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var res string
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		flag := true
		for j := 1; j < len(strs); j++ {
			if i < len(strs[j]) && strs[j][i] == c {
				continue
			} else {
				flag = false
				break
			}
		}
		if flag {
			res += string(c)
		} else {
			break
		}
	}
	return res
}

// @lc code=end
