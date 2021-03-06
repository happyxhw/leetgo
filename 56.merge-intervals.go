/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 *
 * https://leetcode.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (40.38%)
 * Likes:    5967
 * Dislikes: 345
 * Total Accepted:    757.5K
 * Total Submissions: 1.9M
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * Given an array of intervals where intervals[i] = [starti, endi], merge all
 * overlapping intervals, and return an array of the non-overlapping intervals
 * that cover all the intervals in the input.
 *
 *
 * Example 1:
 *
 *
 * Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
 * Output: [[1,6],[8,10],[15,18]]
 * Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into
 * [1,6].
 *
 *
 * Example 2:
 *
 *
 * Input: intervals = [[1,4],[4,5]]
 * Output: [[1,5]]
 * Explanation: Intervals [1,4] and [4,5] are considered overlapping.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= intervals.length <= 10^4
 * intervals[i].length == 2
 * 0 <= starti <= endi <= 10^4
 *
 *
 */
package leetgo

import (
	"sort"
)

// @lc code=start

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	var res [][]int
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	start, end := intervals[0][0], intervals[0][1]
	for _, item := range intervals {
		if end >= item[0] {
			if end < item[1] {
				end = item[1]
			}
		} else {
			res = append(res, []int{start, end})
			start, end = item[0], item[1]
		}
	}
	res = append(res, []int{start, end})
	return res
}

// @lc code=end
