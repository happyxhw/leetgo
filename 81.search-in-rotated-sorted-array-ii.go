/*
 * @lc app=leetcode id=81 lang=golang
 *
 * [81] Search in Rotated Sorted Array II
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Medium (33.41%)
 * Likes:    1823
 * Dislikes: 527
 * Total Accepted:    282.9K
 * Total Submissions: 846.8K
 * Testcase Example:  '[2,5,6,0,0,1,2]\n0'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 *
 * (i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).
 *
 * You are given a target value to search. If found in the array return true,
 * otherwise return false.
 *
 * Example 1:
 *
 *
 * Input: nums = [2,5,6,0,0,1,2], target = 0
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [2,5,6,0,0,1,2], target = 3
 * Output: false
 *
 * Follow up:
 *
 *
 * This is a follow up problem to Search in Rotated Sorted Array, where nums
 * may contain duplicates.
 * Would this affect the run-time complexity? How and why?
 *
 *
 */
package leetgo

// @lc code=start
func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[start] && nums[mid] == nums[end] {
			if mid-start >= 1 && search(nums[start:mid], target) {
				return true
			}
			if end-mid >= 1 && search(nums[mid+1:end+1], target) {
				return true
			}
			return false
		} else {
			if nums[mid] >= nums[start] {
				if target <= nums[mid] && target >= nums[start] {
					end = mid - 1
				} else {
					start = mid + 1
				}
			} else {
				if target >= nums[mid] && target <= nums[end] {
					start = mid + 1
				} else {
					end = mid - 1
				}
			}
		}
	}
	return false
}

// @lc code=end
