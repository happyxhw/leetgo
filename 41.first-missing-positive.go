/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 *
 * https://leetcode.com/problems/first-missing-positive/description/
 *
 * algorithms
 * Hard (33.23%)
 * Likes:    4819
 * Dislikes: 906
 * Total Accepted:    420.1K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,2,0]'
 *
 * Given an unsorted integer array nums, find the smallest missing positive
 * integer.
 *
 * Follow up: Could you implement an algorithm that runs in O(n) time and uses
 * constant extra space.?
 *
 *
 * Example 1:
 * Input: nums = [1,2,0]
 * Output: 3
 * Example 2:
 * Input: nums = [3,4,-1,1]
 * Output: 2
 * Example 3:
 * Input: nums = [7,8,9,11,12]
 * Output: 1
 *
 *
 * Constraints:
 *
 *
 * 0 <= nums.length <= 300
 * -2^31 <= nums[i] <= 2^31 - 1
 *
 *
 */
package leetgo

// @lc code=start
func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, item := range nums {
		if i != item-1 {
			return i + 1
		}
	}
	return n + 1
}

// @lc code=end
