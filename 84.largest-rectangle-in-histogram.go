/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 *
 * https://leetcode.com/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (36.27%)
 * Likes:    4814
 * Dislikes: 100
 * Total Accepted:    312.1K
 * Total Submissions: 860.5K
 * Testcase Example:  '[2,1,5,6,2,3]'
 *
 * Given n non-negative integers representing the histogram's bar height where
 * the width of each bar is 1, find the area of largest rectangle in the
 * histogram.
 *
 *
 *
 *
 * Above is a histogram where width of each bar is 1, given height =
 * [2,1,5,6,2,3].
 *
 *
 *
 *
 * The largest rectangle is shown in the shaded area, which has area = 10
 * unit.
 *
 *
 *
 * Example:
 *
 *
 * Input: [2,1,5,6,2,3]
 * Output: 10
 *
 *
 */
package leetgo

// @lc code=start
// 单调栈
func largestRectangleArea(heights []int) int {
	var res int
	var stack []int
	// 保证可以正常结束
	heights = append(heights, 0)

	for i := 0; i < len(heights); {
		// 保持单调
		if len(stack) == 0 || heights[stack[len(stack)-1]] < heights[i] {
			stack = append(stack, i)
			i++
		} else {
			pre := stack[len(stack)-1]
			// 不断出栈，直到 i 满足单调性，将 i 推入栈
			stack = stack[:len(stack)-1]

			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			t := heights[pre] * width
			if t > res {
				res = t
			}

		}
	}
	return res
}

// @lc code=end
