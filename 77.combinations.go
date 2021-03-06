/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 *
 * https://leetcode.com/problems/combinations/description/
 *
 * algorithms
 * Medium (56.44%)
 * Likes:    1899
 * Dislikes: 78
 * Total Accepted:    330.7K
 * Total Submissions: 584K
 * Testcase Example:  '4\n2'
 *
 * Given two integers n and k, return all possible combinations of k numbers
 * out of 1 ... n.
 *
 * You may return the answer in any order.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 4, k = 2
 * Output:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input: n = 1, k = 1
 * Output: [[1]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 20
 * 1 <= k <= n
 *
 *
 */
package leetgo

// @lc code=start
func combine(n int, k int) [][]int {
	return helper077(1, k, n, [][]int{}, []int{})
}

func helper077(start, k, n int, res [][]int, path []int) [][]int {
	if k == 0 {
		t := make([]int, 0, k)
		t = append(t, path...)
		res = append(res, t)
		return res
	}
	for i := start; i <= n; i++ {
		path = append(path, i)
		k--
		res = helper077(i+1, k, n, res, path)
		k++
		path = path[:len(path)-1]
	}
	return res
}

// @lc code=end
