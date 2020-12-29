/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */
package leetgo

// @lc code=start
func removeDuplicates_1(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	var j int
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[j] {
			continue
		} else {
			j++
			if j != i {
				nums[j] = nums[i]
			}
		}
	}
	j++
	return j
}

// @lc code=end
