/*
 * @lc app=leetcode id=60 lang=golang
 *
 * [60] Permutation Sequence
 *
 * https://leetcode.com/problems/permutation-sequence/description/
 *
 * algorithms
 * Hard (39.05%)
 * Likes:    1956
 * Dislikes: 352
 * Total Accepted:    215.8K
 * Total Submissions: 552K
 * Testcase Example:  '3\n3'
 *
 * The set [1, 2, 3, ..., n] contains a total of n! unique permutations.
 *
 * By listing and labeling all of the permutations in order, we get the
 * following sequence for n = 3:
 *
 *
 * "123"
 * "132"
 * "213"
 * "231"
 * "312"
 * "321"
 *
 *
 * Given n and k, return the k^th permutation sequence.
 *
 *
 * Example 1:
 * Input: n = 3, k = 3
 * Output: "213"
 * Example 2:
 * Input: n = 4, k = 9
 * Output: "2314"
 * Example 3:
 * Input: n = 3, k = 1
 * Output: "123"
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 9
 * 1 <= k <= n!
 *
 *
 */
package leetgo

// @lc code=start
func getPermutation(n int, k int) string {
	used := make([]bool, n)
	res, _ := helper060(1, k, n, []byte{}, used)
	return string(res)
}

func helper060(start, m, n int, s []byte, used []bool) ([]byte, int) {
	if start > n {
		if m == 1 {
			return s, m
		}
		m--
		return nil, m
	}
	for i := 1; i <= n; i++ {
		if used[i-1] {
			continue
		}
		used[i-1] = true
		s = append(s, byte(i)+'0')
		var res []byte
		res, m = helper060(start+1, m, n, s, used)
		if len(res) > 0 {
			return res, m
		}
		s = s[:len(s)-1]
		used[i-1] = false
	}
	return nil, m
}

// @lc code=end
