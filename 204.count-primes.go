/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 *
 * https://leetcode.com/problems/count-primes/description/
 *
 * algorithms
 * Easy (32.22%)
 * Likes:    2742
 * Dislikes: 730
 * Total Accepted:    439.8K
 * Total Submissions: 1.4M
 * Testcase Example:  '10'
 *
 * Count the number of prime numbers less than a non-negative number, n.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 10
 * Output: 4
 * Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
 *
 *
 * Example 2:
 *
 *
 * Input: n = 0
 * Output: 0
 *
 *
 * Example 3:
 *
 *
 * Input: n = 1
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= n <= 5 * 10^6
 *
 *
 */
package leetgo

import "math"

// @lc code=start
func countPrimes(n int) int {
	if n < 2 {
		return 0
	}

	primes := make([]int, n)
	p := int(math.Sqrt(float64(n)))
	res := 0

	for i := 2; i <= p; i++ {
		if primes[i] == 1 {
			continue
		}

		for j := i * i; j < n; j += i {
			if primes[j] == 0 {
				primes[j] = 1
				res++
			}
		}
	}

	return n - res - 2
}

// @lc code=end
