/*
 * @lc app=leetcode id=135 lang=golang
 *
 * [135] Candy
 *
 * https://leetcode.com/problems/candy/description/
 *
 * algorithms
 * Hard (32.72%)
 * Likes:    1292
 * Dislikes: 184
 * Total Accepted:    147.3K
 * Total Submissions: 448.9K
 * Testcase Example:  '[1,0,2]'
 *
 * There are N children standing in a line. Each child is assigned a rating
 * value.
 *
 * You are giving candies to these children subjected to the following
 * requirements:
 *
 *
 * Each child must have at least one candy.
 * Children with a higher rating get more candies than their neighbors.
 *
 *
 * What is the minimum candies you must give?
 *
 * Example 1:
 *
 *
 * Input: [1,0,2]
 * Output: 5
 * Explanation: You can allocate to the first, second and third child with 2,
 * 1, 2 candies respectively.
 *
 *
 * Example 2:
 *
 *
 * Input: [1,2,2]
 * Output: 4
 * Explanation: You can allocate to the first, second and third child with 1,
 * 2, 1 candies respectively.
 * ⁠            The third child gets 1 candy because it satisfies the above two
 * conditions.
 *
 *
 */
package leetgo

// @lc code=start
func candy(ratings []int) int {
	res := make([]int, len(ratings))
	for i := range res {
		res[i] = 1
	}
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i+1] > ratings[i] {
			res[i+1] = res[i] + 1
		}
	}
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			if res[i] >= res[i-1] {
				res[i-1] = res[i] + 1
			}
		}
	}
	t := 0
	for _, item := range res {
		t += item
	}
	return t
}

// @lc code=end
