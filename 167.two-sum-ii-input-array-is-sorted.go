/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input array is sorted
 *
 * https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
 *
 * algorithms
 * Easy (55.61%)
 * Likes:    2365
 * Dislikes: 688
 * Total Accepted:    526.4K
 * Total Submissions: 946.7K
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers numbers that is already sorted in ascending
 * order, find two numbers such that they add up to a specific target number.
 *
 * Return the indices of the two numbers (1-indexed) as an integer array answer
 * of size 2, where 1 <= answer[0] < answer[1] <= numbers.length.
 *
 * You may assume that each input would have exactly one solution and you may
 * not use the same element twice.
 *
 *
 * Example 1:
 *
 *
 * Input: numbers = [2,7,11,15], target = 9
 * Output: [1,2]
 * Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
 *
 *
 * Example 2:
 *
 *
 * Input: numbers = [2,3,4], target = 6
 * Output: [1,3]
 *
 *
 * Example 3:
 *
 *
 * Input: numbers = [-1,0], target = -1
 * Output: [1,2]
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= numbers.length <= 3 * 10^4
 * -1000 <= numbers[i] <= 1000
 * numbers is sorted in increasing order.
 * -1000 <= target <= 1000
 * Only one valid answer exists.
 *
 *
 */
package leetgo

// @lc code=start
func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		t := nums[left] + nums[right]
		if t == target {
			return []int{left + 1, right + 1}
		} else if t > target {
			for right > left && nums[right] == nums[right-1] {
				right--
			}
			right--
		} else {
			for right > left && nums[left] == nums[left+1] {
				left++
			}
			left++
		}
	}
	return []int{-1, -1}
}

// @lc code=end
