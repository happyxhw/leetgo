/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */
package leetgo

import "sort"

// @lc code=start
func permuteUnique(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	path := make([]int, 0, len(nums))
	used := make([]bool, len(nums))
	return dfs47(nums, res, path, used)
}

func dfs47(nums []int, res [][]int, path []int, used []bool) [][]int {
	if len(path) == len(nums) {
		copyPath := make([]int, 0, len(nums))
		copyPath = append(copyPath, path...)
		res = append(res, copyPath)
		return res
	}
	for i := 0; i < len(nums); i++ {
		if used[i] || (i > 0 && nums[i] == nums[i-1] && used[i-1]) {
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		res = dfs47(nums, res, path, used)
		path = path[:len(path)-1]
		used[i] = false
	}
	return res
}

// @lc code=end
