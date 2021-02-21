/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 *
 * https://leetcode.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (60.01%)
 * Likes:    4550
 * Dislikes: 254
 * Total Accepted:    791.2K
 * Total Submissions: 1.3M
 * Testcase Example:  '[3,2,3]'
 *
 * Given an array nums of size n, return the majority element.
 *
 * The majority element is the element that appears more than ⌊n / 2⌋ times.
 * You may assume that the majority element always exists in the array.
 *
 *
 * Example 1:
 * Input: nums = [3,2,3]
 * Output: 3
 * Example 2:
 * Input: nums = [2,2,1,1,1,2,2]
 * Output: 2
 *
 *
 * Constraints:
 *
 *
 * n == nums.length
 * 1 <= n <= 5 * 10^4
 * -2^31 <= nums[i] <= 2^31 - 1
 *
 *
 *
 * Follow-up: Could you solve the problem in linear time and in O(1) space?
 */
package leetgo

import (
	"sort"
)

// @lc code=start
func majorityElement(nums []int) int {
	sort.Ints(nums)
	var t int
	for i := range nums {
		if i == 0 || nums[i] == nums[i-1] {
			t++
			continue
		}
		if t > len(nums)/2 {
			return nums[i-1]
		}
		t = 1
	}
	if t > len(nums)/2 {
		return nums[len(nums)-1]
	}
	return -1
}

// @lc code=end
