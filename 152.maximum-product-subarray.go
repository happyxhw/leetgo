/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 *
 * https://leetcode.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (32.73%)
 * Likes:    6230
 * Dislikes: 205
 * Total Accepted:    446.4K
 * Total Submissions: 1.4M
 * Testcase Example:  '[2,3,-2,4]'
 *
 * Given an integer array nums, find the contiguous subarray within an array
 * (containing at least one number) which has the largest product.
 *
 * Example 1:
 *
 *
 * Input: [2,3,-2,4]
 * Output: 6
 * Explanation: [2,3] has the largest product 6.
 *
 *
 * Example 2:
 *
 *
 * Input: [-2,0,-1]
 * Output: 0
 * Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 *
 */
package leetgo

// @lc code=start
func maxProduct(nums []int) int {
	res := nums[0]

	curMax, curMin := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			curMax, curMin = curMin, curMax
		}

		if nums[i] < curMax*nums[i] {
			curMax = curMax * nums[i]
		} else {
			curMax = nums[i]
		}
		if nums[i] > curMin*nums[i] {
			curMin = curMin * nums[i]
		} else {
			curMin = nums[i]
		}

		if res < curMax {
			res = curMax
		}
	}
	return res
}

// @lc code=end
