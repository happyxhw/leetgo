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
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	m, n := len(board)-1, len(board[0])-1

	visited := make([][]bool, m+1)
	for i := range visited {
		visited[i] = make([]bool, n+1)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if helper079(board, word, i, j, m, n, 0, visited) {
				return true
			}
		}
	}

	return false
}

func helper079(board [][]byte, word string, i, j, m, n, idx int, visited [][]bool) bool {
	if idx == len(word) {
		return true
	}
	if i < 0 || i > m || j < 0 || j > n || visited[i][j] || board[i][j] != word[idx] {
		return false
	}
	visited[i][j] = true
	res := helper079(board, word, i-1, j, m, n, idx+1, visited) ||
		helper079(board, word, i+1, j, m, n, idx+1, visited) ||
		helper079(board, word, i, j-1, m, n, idx+1, visited) ||
		helper079(board, word, i, j+1, m, n, idx+1, visited)
	visited[i][j] = false

	return res
}

// @lc code=end
