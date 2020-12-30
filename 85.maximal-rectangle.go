/*
 * @lc app=leetcode id=85 lang=golang
 *
 * [85] Maximal Rectangle
 *
 * https://leetcode.com/problems/maximal-rectangle/description/
 *
 * algorithms
 * Hard (39.02%)
 * Likes:    3652
 * Dislikes: 79
 * Total Accepted:    208.1K
 * Total Submissions: 533.2K
 * Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
 *
 * Given a rows x cols binary matrix filled with 0's and 1's, find the largest
 * rectangle containing only 1's and return its area.
 *
 *
 * Example 1:
 *
 *
 * Input: matrix =
 * [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
 * Output: 6
 * Explanation: The maximal rectangle is shown in the above picture.
 *
 *
 * Example 2:
 *
 *
 * Input: matrix = []
 * Output: 0
 *
 *
 * Example 3:
 *
 *
 * Input: matrix = [["0"]]
 * Output: 0
 *
 *
 * Example 4:
 *
 *
 * Input: matrix = [["1"]]
 * Output: 1
 *
 *
 * Example 5:
 *
 *
 * Input: matrix = [["0","0"]]
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * rows == matrix.length
 * cols == matrix.length
 * 0 <= row, cols <= 200
 * matrix[i][j] is '0' or '1'.
 *
 *
 */
package leetgo

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	var res int
	m, n := len(matrix), len(matrix[0])
	heights := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				heights[j] = 0
			} else {
				heights[j] += 1
			}
		}
		t := helper085(heights)
		if t > res {
			res = t
		}
	}

	return res
}

func helper085(heights []int) int {
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
