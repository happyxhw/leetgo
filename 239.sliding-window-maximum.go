/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 *
 * https://leetcode.com/problems/sliding-window-maximum/description/
 *
 * algorithms
 * Hard (44.42%)
 * Likes:    5065
 * Dislikes: 211
 * Total Accepted:    361.5K
 * Total Submissions: 811.1K
 * Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
 *
 * You are given an array of integers nums, there is a sliding window of size k
 * which is moving from the very left of the array to the very right. You can
 * only see the k numbers in the window. Each time the sliding window moves
 * right by one position.
 *
 * Return the max sliding window.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
 * Output: [3,3,5,5,6,7]
 * Explanation:
 * Window position                Max
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7       3
 * ⁠1 [3  -1  -3] 5  3  6  7       3
 * ⁠1  3 [-1  -3  5] 3  6  7       5
 * ⁠1  3  -1 [-3  5  3] 6  7       5
 * ⁠1  3  -1  -3 [5  3  6] 7       6
 * ⁠1  3  -1  -3  5 [3  6  7]      7
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1], k = 1
 * Output: [1]
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [1,-1], k = 1
 * Output: [1,-1]
 *
 *
 * Example 4:
 *
 *
 * Input: nums = [9,11], k = 2
 * Output: [11]
 *
 *
 * Example 5:
 *
 *
 * Input: nums = [4,-2], k = 2
 * Output: [4]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 * 1 <= k <= nums.length
 *
 *
 */
package leetgo

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0, len(nums))
	var h Heap
	for i, v := range nums {
		for h.Size() != 0 && h.Top().Index <= i-k {
			h.Pop()
		}
		h.Push(&pair{Val: v, Index: i})
		if i >= k-1 {
			res = append(res, h.Top().Val)
		}
	}
	return res
}

type pair struct {
	Val   int
	Index int
}

type Heap []*pair

func (h Heap) Size() int {
	return len(h)
}

func (h Heap) Top() *pair {
	return h[0]
}

func (h *Heap) Push(item *pair) {
	*h = append(*h, item)
	k := len(*h) - 1
	for k > 0 {
		if item.Val > (*h)[(k-1)/2].Val {
			(*h)[k] = (*h)[(k-1)/2]
		} else {
			break
		}
		k = (k - 1) / 2
	}
	(*h)[k] = item
}

func (h *Heap) Pop() {
	if len(*h) == 0 {
		return
	}
	(*h)[0] = (*h)[len((*h))-1]
	*h = (*h)[:len((*h))-1]
	if len(*h) == 0 {
		return
	}
	item := (*h)[0]
	k, n := 0, len(*h)
	for 2*k+1 < n {
		maxIndex := 2*k + 1
		if 2*k+2 < n && (*h)[2*k+2].Val > (*h)[2*k+1].Val {
			maxIndex++
		}
		if item.Val < (*h)[maxIndex].Val {
			(*h)[k] = (*h)[maxIndex]
		} else {
			break
		}
		k = maxIndex
	}
	(*h)[k] = item
}

// @lc code=end
