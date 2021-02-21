/*
 * @lc app=leetcode id=164 lang=golang
 *
 * [164] Maximum Gap
 *
 * https://leetcode.com/problems/maximum-gap/description/
 *
 * algorithms
 * Hard (36.79%)
 * Likes:    1079
 * Dislikes: 206
 * Total Accepted:    98.9K
 * Total Submissions: 268.8K
 * Testcase Example:  '[3,6,9,1]'
 *
 * Given an unsorted array, find the maximum difference between the successive
 * elements in its sorted form.
 *
 * Return 0 if the array contains less than 2 elements.
 *
 * Example 1:
 *
 *
 * Input: [3,6,9,1]
 * Output: 3
 * Explanation: The sorted form of the array is [1,3,6,9], either
 * (3,6) or (6,9) has the maximum difference 3.
 *
 * Example 2:
 *
 *
 * Input: [10]
 * Output: 0
 * Explanation: The array contains less than 2 elements, therefore return 0.
 *
 * Note:
 *
 *
 * You may assume all elements in the array are non-negative integers and fit
 * in the 32-bit signed integer range.
 * Try to solve it in linear time/space.
 *
 *
 */
package leetgo

// @lc code=start
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	maxVal, minVal := nums[0], nums[0]
	for _, n := range nums {
		if n > maxVal {
			maxVal = n
		}
		if n < minVal {
			minVal = n
		}
	}

	size := (maxVal-minVal)/len(nums) + 1
	cnt := (maxVal-minVal)/size + 1

	minList, maxList := make([]int, cnt), make([]int, cnt)
	for i := range minList {
		minList[i] = -1
	}
	for j := range maxList {
		maxList[j] = -1
	}

	for _, n := range nums {
		idx := (n - minVal) / size
		if minList[idx] == -1 || n < minList[idx] {
			minList[idx] = n
		}
		if maxList[idx] == -1 || n > maxList[idx] {
			maxList[idx] = n
		}
	}
	var pre, res int
	for i := 1; i < cnt; i++ {
		if minList[i] == -1 || maxList[i] == -1 {
			continue
		}
		if maxList[pre] != -1 {
			t := minList[i] - maxList[pre]
			if t > res {
				res = t
			}
		}
		pre = i
	}

	return res
}

// @lc code=end
