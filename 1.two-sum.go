/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
package leetgo

// @lc code=start
func twoSum_1(nums []int, target int) []int {
	sumMap := make(map[int]int, len(nums))
	for i, item := range nums {
		if v, ok := sumMap[target-item]; ok {
			return []int{v, i}
		} else {
			sumMap[item] = i
		}
	}
	return []int{}
}

// @lc code=end
