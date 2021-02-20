/*
 * @lc app=leetcode id=398 lang=golang
 *
 * [398] Random Pick Index
 *
 * https://leetcode.com/problems/random-pick-index/description/
 *
 * algorithms
 * Medium (57.64%)
 * Likes:    604
 * Dislikes: 780
 * Total Accepted:    110.4K
 * Total Submissions: 191.5K
 * Testcase Example:  '["Solution","pick"]\n[[[1,2,3,3,3]],[3]]'
 *
 * Given an array of integers with possible duplicates, randomly output the
 * index of a given target number. You can assume that the given target number
 * must exist in the array.
 *
 * Note:
 * The array size can be very large. Solution that uses too much extra space
 * will not pass the judge.
 *
 * Example:
 *
 *
 * int[] nums = new int[] {1,2,3,3,3};
 * Solution solution = new Solution(nums);
 *
 * // pick(3) should return either index 2, 3, or 4 randomly. Each index should
 * have equal probability of returning.
 * solution.pick(3);
 *
 * // pick(1) should return 0. Since in the array only nums[0] is equal to 1.
 * solution.pick(1);
 *
 *
 */
package leetgo

import (
	"math/rand"
	"time"
)

// @lc code=start
type Solution struct {
	pool []int
}

func Constructor_0(nums []int) Solution {
	return Solution{pool: nums}
}

func (this *Solution) Pick(target int) int {
	cnt, res := 0, -1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(this.pool); i++ {
		if this.pool[i] != target {
			continue
		}
		cnt++
		if rand.Int()%cnt == 0 {
			res = i
		}
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
// @lc code=end
