/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 *
 * https://leetcode.com/problems/word-search/description/
 *
 * algorithms
 * Medium (36.38%)
 * Likes:    4913
 * Dislikes: 215
 * Total Accepted:    580.5K
 * Total Submissions: 1.6M
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * Given an m x n board and a word, find if the word exists in the grid.
 *
 * The word can be constructed from letters of sequentially adjacent cells,
 * where "adjacent" cells are horizontally or vertically neighboring. The same
 * letter cell may not be used more than once.
 *
 *
 * Example 1:
 *
 *
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word
 * = "ABCCED"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word
 * = "SEE"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word
 * = "ABCB"
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * m == board.length
 * n = board[i].length
 * 1 <= m, n <= 200
 * 1 <= word.length <= 10^3
 * board and word consists only of lowercase and uppercase English letters.
 *
 *
 */
package leetgo

// @lc code=start
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	m, n := len(board)-1, len(board[0])-1
	return helper079(board, word, 0, 0, m, n, 0)
}

func helper079(board [][]byte, word string, row, col, m, n, index int) bool {
	if row < 0 || row > m || col < 0 || col > n {
		return false
	}
	for i := row; i < m; i++ {
		for j := col; j < n; j++ {
			if board[i][j] == word[index] {

			}
		}
	}
	return false
}

// @lc code=end
