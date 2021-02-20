/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 *
 * https://leetcode.com/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (43.74%)
 * Likes:    6342
 * Dislikes: 149
 * Total Accepted:    487.2K
 * Total Submissions: 1.1M
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * Given an integer array nums, return the length of the longest strictly
 * increasing subsequence.
 *
 * A subsequence is a sequence that can be derived from an array by deleting
 * some or no elements without changing the order of the remaining elements.
 * For example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [10,9,2,5,3,7,101,18]
 * Output: 4
 * Explanation: The longest increasing subsequence is [2,3,7,101], therefore
 * the length is 4.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [0,1,0,3,2,3]
 * Output: 4
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [7,7,7,7,7,7,7]
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 2500
 * -10^4 <= nums[i] <= 10^4
 *
 *
 *
 * Follow up:
 *
 *
 * Could you come up with the O(n^2) solution?
 * Could you improve it to O(n log(n)) time complexity?
 *
 *
 */
package leetgo

// 这并不是最优的方法，这是暴力解的，复杂度是 o(n*n)
// 更好的方法是二分查找，但是我没想明白，复杂度是 o(nlogn)

// @lc code=start
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	var res int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}

// @lc code=end
