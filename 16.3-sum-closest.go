/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */
package leetgo

import "sort"

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			t := nums[i] + nums[left] + nums[right]
			if t == target {
				return t
			}

			if abs016(res-target) > abs016(t-target) {
				res = t
			}

			if t < target {
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else {
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			}
		}
	}
	return res
}

func abs016(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end
