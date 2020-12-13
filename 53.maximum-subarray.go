/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */
package leetgo

// dynamic-programming

// @lc code=start
func maxSubArray(nums []int) int {
	cur, max := 0, nums[0]
	for _, n := range nums {
		if cur+n > n {
			cur += n
		} else {
			cur = n
		}
		if cur > max {
			max = cur
		}
	}

	return max
}

// @lc code=end
