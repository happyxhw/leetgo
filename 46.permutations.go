/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */
package leetgo

// @lc code=start
func permute(nums []int) [][]int {
	var res [][]int
	path := make([]int, 0, len(nums))
	used := make([]bool, len(nums))
	res = dfs46(nums, res, path, used)
	return res
}

func dfs46(nums []int, res [][]int, path []int, used []bool) [][]int {
	if len(path) == len(nums) {
		copyPath := make([]int, 0, len(nums))
		copyPath = append(copyPath, path...)
		res = append(res, copyPath)
		return res
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		res = dfs46(nums, res, path, used)
		path = path[:len(path)-1]
		used[i] = false
	}
	return res
}

// @lc code=end
