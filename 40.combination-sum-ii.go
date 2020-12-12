/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */
package leetgo

import (
	"fmt"
	"sort"
)

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var tempPath []int
	sort.Ints(candidates)
	fmt.Println(candidates)
	res, _ = dfs2(candidates, res, tempPath, 0, target)
	return res
}

func dfs2(nums []int, res [][]int, tempPath []int, start, remain int) ([][]int, []int) {
	if remain < 0 {
		return res, tempPath
	}
	if remain == 0 {
		copyPath := make([]int, 0, len(tempPath))
		copyPath = append(copyPath, tempPath...)
		res = append(res, copyPath)
	} else {
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			tempPath = append(tempPath, nums[i])
			res, tempPath = dfs2(nums, res, tempPath, i+1, remain-nums[i])
			tempPath = tempPath[:len(tempPath)-1]
		}
	}
	return res, tempPath
}

// @lc code=end
