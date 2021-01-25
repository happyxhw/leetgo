/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 *
 * https://leetcode.com/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Hard (46.03%)
 * Likes:    4544
 * Dislikes: 219
 * Total Accepted:    366.2K
 * Total Submissions: 794.2K
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * Given an unsorted array of integers nums, return the length of the longest
 * consecutive elements sequence.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [100,4,200,1,3,2]
 * Output: 4
 * Explanation: The longest consecutive elements sequence is [1, 2, 3, 4].
 * Therefore its length is 4.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [0,3,7,2,5,8,4,6,0,1]
 * Output: 9
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= nums.length <= 10^4
 * -10^9 <= nums[i] <= 10^9
 *
 *
 *
 * Follow up: Could you implement the O(n) solution?
 */
package leetgo

// @lc code=start
func longestConsecutive(nums []int) int {
	numsMap := make(map[int]bool, len(nums))
	for _, n := range nums {
		numsMap[n] = true
	}
	var maxN int
	for _, n := range nums {
		if !numsMap[n-1] {
			t := n + 1
			for numsMap[t] {
				t++
			}
			if t-n > maxN {
				maxN = t - n
			}
		}
	}
	return maxN
}

// @lc code=end
