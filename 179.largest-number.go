/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 *
 * https://leetcode.com/problems/largest-number/description/
 *
 * algorithms
 * Medium (30.62%)
 * Likes:    2806
 * Dislikes: 298
 * Total Accepted:    237.4K
 * Total Submissions: 775.2K
 * Testcase Example:  '[10,2]'
 *
 * Given a list of non-negative integers nums, arrange them such that they form
 * the largest number.
 *
 * Note: The result may be very large, so you need to return a string instead
 * of an integer.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [10,2]
 * Output: "210"
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [3,30,34,5,9]
 * Output: "9534330"
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [1]
 * Output: "1"
 *
 *
 * Example 4:
 *
 *
 * Input: nums = [10]
 * Output: "10"
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 100
 * 0 <= nums[i] <= 10^9
 *
 *
 */
package leetgo

import (
	"sort"
	"strconv"
)

// @lc code=start
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return compare(nums[i], nums[j])
	})
	var res string
	for i, n := range nums {
		if i == 0 && n == 0 {
			return "0"
		}
		res += strconv.FormatInt(int64(n), 10)
	}
	return res
}

func compare(x, y int) bool {
	xx, yy := x, y
	t1, t2 := x, y
	for {
		xx = xx / 10
		t2 *= 10
		if xx == 0 {
			break
		}
	}
	for {
		yy = yy / 10
		t1 *= 10
		if yy == 0 {
			break
		}
	}
	t1, t2 = t1+y, t2+x
	// fmt.Println(x, y, t1, t2)
	return t1 > t2
}

// @lc code=end
