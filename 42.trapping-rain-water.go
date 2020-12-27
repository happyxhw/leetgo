/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 *
 * https://leetcode.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (50.40%)
 * Likes:    9285
 * Dislikes: 140
 * Total Accepted:    634K
 * Total Submissions: 1.3M
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * Given n non-negative integers representing an elevation map where the width
 * of each bar is 1, compute how much water it can trap after raining.
 *
 *
 * Example 1:
 *
 *
 * Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * Output: 6
 * Explanation: The above elevation map (black section) is represented by array
 * [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue
 * section) are being trapped.
 *
 *
 * Example 2:
 *
 *
 * Input: height = [4,2,0,3,2,5]
 * Output: 9
 *
 *
 *
 * Constraints:
 *
 *
 * n == height.length
 * 0 <= n <= 3 * 10^4
 * 0 <= height[i] <= 10^5
 *
 *
 */
package leetgo

// @lc code=start
func trap(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}
	var stack []int
	var area, i int
	for i < n {
		if len(stack) == 0 || height[i] <= height[stack[len(stack)-1]] {
			stack = append(stack, i)
			i++
		} else {
			pre := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				minHeight := height[i]
				if minHeight > height[top] {
					minHeight = height[top]
				}
				area += (minHeight - height[pre]) * (i - top - 1)
			}
		}
	}
	return area
}

// @lc code=end
