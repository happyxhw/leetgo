/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */
package leetgo

// @lc code=start
func search_0(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[start] {
			if target <= nums[mid] && target >= nums[start] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}

// public int findFirstOrLast(int[] nums, int target,boolean firstOrLast){
// 	int resIndex = -1;
// 	int left = 0;
// 	int right = nums.length-1;
// 	int mid;
// 	while (left<=right){
// 		mid = left + (right - left)/2;
// 		if (firstOrLast)
// 			if (nums[mid]>=target) right = mid -1;
// 			else left = mid+1;
// 		else
// 			if (nums[mid]<=target) left = mid+1;
// 			else right = mid-1;
// 		if (nums[mid] == target) resIndex = mid;
// 	}
// 	return resIndex;
// }
// @lc code=end
