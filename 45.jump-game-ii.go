/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 *
 * https://leetcode.com/problems/jump-game-ii/description/
 *
 * algorithms
 * Hard (30.79%)
 * Likes:    3363
 * Dislikes: 161
 * Total Accepted:    299.1K
 * Total Submissions: 960.1K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * Given an array of non-negative integers nums, you are initially positioned
 * at the first index of the array.
 *
 * Each element in the array represents your maximum jump length at that
 * position.
 *
 * Your goal is to reach the last index in the minimum number of jumps.
 *
 * You can assume that you can always reach the last index.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [2,3,1,1,4]
 * Output: 2
 * Explanation: The minimum number of jumps to reach the last index is 2. Jump
 * 1 step from index 0 to 1, then 3 steps to the last index.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [2,3,0,1,4]
 * Output: 2
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 3 * 10^4
 * 0 <= nums[i] <= 10^5
 *
 *
 */
package leetgo

// @lc code=start
func jump(nums []int) int {
	var jumps, curEnd, curFarthest int
	for i := 0; i < len(nums)-1; i++ {
		if curFarthest < i+nums[i] {
			curFarthest = i + nums[i]
		}
		if i == curEnd {
			curEnd = curFarthest
			jumps++
			if curEnd >= len(nums)-1 {
				break
			}
		}
	}
	return jumps
}

// @lc code=end
