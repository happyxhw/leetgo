/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */
package leetgo

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			c := board[i][j]
			if c == '.' {
				continue
			}
			for k := 0; k < 9; k++ {
				if k != i && board[k][j] != '.' && board[k][j] == c {
					return false
				}
				if k != j && board[i][k] != '.' && board[i][k] == c {
					return false
				}
				t := 3*(i/3) + k/3
				s := 3*(j/3) + k%3
				if t != i && s != j && board[t][s] != '.' && board[t][s] == c {
					return false
				}
			}
		}
	}
	return true
}

// @lc code=end
