/*
 * @lc app=leetcode id=154 lang=golang
 *
 * [154] Find Minimum in Rotated Sorted Array II
 *
 * https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Hard (42.00%)
 * Likes:    1427
 * Dislikes: 267
 * Total Accepted:    235K
 * Total Submissions: 559.5K
 * Testcase Example:  '[1,3,5]'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 *
 * (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
 *
 * Find the minimum element.
 *
 * The array may contain duplicates.
 *
 * Example 1:
 *
 *
 * Input: [1,3,5]
 * Output: 1
 *
 * Example 2:
 *
 *
 * Input: [2,2,2,0,1]
 * Output: 0
 *
 * Note:
 *
 *
 * This is a follow up problem to Find Minimum in Rotated Sorted Array.
 * Would allow duplicates affect the run-time complexity? How and why?
 *
 *
 */
package leetgo

// @lc code=start
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	if nums[left] < nums[right] {
		return nums[left]
	}
	res := nums[left]
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < res {
			res = nums[mid]
		}
		if nums[mid] == nums[left] && nums[mid] == nums[right] {
			if mid != left {
				t := findMin(nums[left:mid])
				if t < res {
					res = t
				}
			}
			if mid != right {
				t := findMin(nums[mid+1 : right+1])
				if t < res {
					res = t
				}
			}
			return res
		} else {
			if nums[mid] <= nums[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return res
}

// @lc code=end
