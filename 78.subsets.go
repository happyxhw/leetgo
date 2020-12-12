/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */
package leetgo

// @lc code=start
func subsets(nums []int) [][]int {
	var res [][]int
	tempPath := make([]int, 0, len(nums))
	res, _ = dfs78(nums, res, tempPath, 0)
	return res
}

func dfs78(nums []int, res [][]int, tempPath []int, start int) ([][]int, []int) {

	copyPath := make([]int, 0, len(tempPath))
	copyPath = append(copyPath, tempPath...)
	res = append(res, copyPath)

	for i := start; i < len(nums); i++ {
		tempPath = append(tempPath, nums[i])
		res, tempPath = dfs78(nums, res, tempPath, i+1)
		tempPath = tempPath[:len(tempPath)-1]
	}
	return res, tempPath

}

// @lc code=end
