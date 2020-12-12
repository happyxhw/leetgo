/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */
package leetgo

import "sort"

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	path := make([]int, 0, len(nums))
	res, _ = dfs90(nums, res, path, 0)
	return res
}

func dfs90(nums []int, res [][]int, path []int, start int) ([][]int, []int) {
	copyPath := make([]int, 0, len(path))
	copyPath = append(copyPath, path...)
	res = append(res, copyPath)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		res, path = dfs90(nums, res, path, i+1)
		path = path[:len(path)-1]
	}

	return res, path
}

// @lc code=end
