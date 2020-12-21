/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */
package leetgo

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	iMin, iMax, halfLen := 0, m, (m+n+1)/2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < m && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			maxLeft := nums1[i-1]
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				if nums1[i-1] < nums2[j-1] {
					maxLeft = nums2[j-1]
				}
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			minRight := nums1[i]
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				if nums1[i] > nums2[j] {
					minRight = nums2[j]
				}
			}
			return (float64(maxLeft) + float64(minRight)) / 2.0
		}
	}
	return 0.0
}

// @lc code=end
