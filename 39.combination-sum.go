/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */
package leetgo

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var temp []int
	res, _ = dfs(candidates, 0, target, res, temp)
	return res
}

func dfs(nums []int, start, target int, res [][]int, temp []int) ([][]int, []int) {
	if target < 0 {
		return res, temp
	}
	if target == 0 {
		tempCopy := make([]int, 0, len(temp))
		tempCopy = append(tempCopy, temp...)
		res = append(res, tempCopy)
	} else {
		for i := start; i < len(nums); i++ {
			temp = append(temp, nums[i])
			res, temp = dfs(nums, i, target-nums[i], res, temp)
			temp = temp[:len(temp)-1]
		}
	}
	return res, temp
}

// @lc code=end
