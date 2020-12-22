/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */
package leetgo

// @lc code=start
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	var maxArea int
	for left < right {
		t := right - left
		var area int
		if height[left] > height[right] {
			area = t * height[right]
			right--
		} else {
			area = t * height[left]
			left++
		}
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

// @lc code=end
