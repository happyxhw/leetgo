/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */
package leetgo

// @lc code=start
var digitMap = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	var res []string
	var path []byte
	res = helper017(digits, res, path, 0)
	return res
}

func helper017(digits string, res []string, path []byte, start int) []string {
	if start >= len(digits) {
		return res
	}
	for _, c := range digitMap[digits[start]] {
		path = append(path, c)
		if len(path) == len(digits) {
			res = append(res, string(path))
		}
		res = helper017(digits, res, path, start+1)
		path = path[:len(path)-1]
	}
	return res
}

// @lc code=end
