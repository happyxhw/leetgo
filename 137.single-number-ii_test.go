/*
 * @lc app=leetcode id=137 lang=golang
 *
 * [137] Single Number II
 *
 * https://leetcode.com/problems/single-number-ii/description/
 *
 * algorithms
 * Medium (53.60%)
 * Likes:    2312
 * Dislikes: 392
 * Total Accepted:    273.4K
 * Total Submissions: 510K
 * Testcase Example:  '[2,2,3,2]'
 *
 * Given an integer array nums where every element appears three times except
 * for one, which appears exactly once. Find the single element and return
 * it.
 *
 *
 * Example 1:
 * Input: nums = [2,2,3,2]
 * Output: 3
 * Example 2:
 * Input: nums = [0,1,0,1,0,1,99]
 * Output: 99
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 3 * 10^4
 * -2^31 <= nums[i] <= 2^31 - 1
 * Each element in nums appears exactly three times except for one element
 * which appears once.
 *
 *
 *
 * Follow up: Your algorithm should have a linear runtime complexity. Could you
 * implement it without using extra memory?
 *
 */
package leetgo

import (
	"fmt"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	x := []int{2, 2, 3, 2}
	fmt.Println(singleNumber(x))
}
