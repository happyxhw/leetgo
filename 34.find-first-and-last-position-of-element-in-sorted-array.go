/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */
package leetgo

// @lc code=start
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	start, end := 0, len(nums)-1
	for start <= end {
		if nums[end] < target {
			break
		}
		mid := (start + end) / 2
		if target <= nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
		if nums[mid] == target {
			res[0] = mid
		}
	}
	if res[0] == -1 {
		return []int{-1, -1}
	}

	start, end = 0, len(nums)-1
	for start <= end {
		if nums[start] > target {
			break
		}
		mid := (start + end) / 2
		if target >= nums[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
		if nums[mid] == target {
			res[1] = mid
		}
	}

	return res
}

// @lc code=end
