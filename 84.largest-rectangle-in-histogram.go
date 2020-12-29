/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 *
 * https://leetcode.com/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (36.27%)
 * Likes:    4814
 * Dislikes: 100
 * Total Accepted:    312.1K
 * Total Submissions: 860.5K
 * Testcase Example:  '[2,1,5,6,2,3]'
 *
 * Given n non-negative integers representing the histogram's bar height where
 * the width of each bar is 1, find the area of largest rectangle in the
 * histogram.
 *
 *
 *
 *
 * Above is a histogram where width of each bar is 1, given height =
 * [2,1,5,6,2,3].
 *
 *
 *
 *
 * The largest rectangle is shown in the shaded area, which has area = 10
 * unit.
 *
 *
 *
 * Example:
 *
 *
 * Input: [2,1,5,6,2,3]
 * Output: 10
 *
 *
 */
package leetgo

// @lc code=start
func largestRectangleArea(heights []int) int {

}

// @lc code=end

// class Solution {
// 	public:
// 		int largestRectangleArea(vector<int> &height) {
// 			int res = 0;
// 			stack<int> st;
// 			height.push_back(0);
// 			for (int i = 0; i < height.size(); ++i) {
// 				if (st.empty() || height[st.top()] < height[i]) {
// 					st.push(i);
// 				} else {
// 					int cur = st.top(); st.pop();
// 					res = max(res, height[cur] * (st.empty() ? i : (i - st.top() - 1)));
// 					--i;
// 				}
// 			}
// 			return res;
// 		}
// 	};
