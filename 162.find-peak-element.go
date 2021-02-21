/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 *
 * https://leetcode.com/problems/find-peak-element/description/
 *
 * algorithms
 * Medium (43.95%)
 * Likes:    2541
 * Dislikes: 2532
 * Total Accepted:    458.4K
 * Total Submissions: 1M
 * Testcase Example:  '[1,2,3,1]'
 *
 * A peak element is an element that is strictly greater than its neighbors.
 *
 * Given an integer array nums, find a peak element, and return its index. If
 * the array contains multiple peaks, return the index to any of the peaks.
 *
 * You may imagine that nums[-1] = nums[n] = -∞.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3,1]
 * Output: 2
 * Explanation: 3 is a peak element and your function should return the index
 * number 2.
 *
 * Example 2:
 *
 *
 * Input: nums = [1,2,1,3,5,6,4]
 * Output: 5
 * Explanation: Your function can return either index number 1 where the peak
 * element is 2, or index number 5 where the peak element is 6.
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 1000
 * -2^31 <= nums[i] <= 2^31 - 1
 * nums[i] != nums[i + 1] for all valid i.
 *
 *
 *
 * Follow up: Could you implement a solution with logarithmic complexity?
 */
package leetgo

// @lc code=start
func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		switch {
		case i == 0 && nums[i] > nums[i+1]:
			return i
		case i == len(nums)-1 && nums[i] > nums[i-1]:
			return i
		case i > 0 && i < len(nums)-1 && nums[i] > nums[i-1] && nums[i] > nums[i+1]:
			return i
		}
	}
	return -1
}

// @lc code=end
