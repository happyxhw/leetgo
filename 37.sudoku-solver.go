/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 */
package leetgo

// @lc code=start
func solveSudoku(board [][]byte) {
	solve(board, 0, 0)
}

func solve(board [][]byte, x, y int) bool {
	for i := x; i < 9; i++ {
		for j := y; j < 9; j++ {
			if board[i][j] != '.' {
				continue
			}
			var s byte
			for s = '1'; s <= '9'; s++ {
				if isValid(board, i, j, s) {
					board[i][j] = s
					if solve(board, i, j+1) {
						return true
					}
					board[i][j] = '.'
				}
			}
			return false
		}
		y = 0
	}
	return true
}

func isValid(board [][]byte, i, j int, s byte) bool {
	x := 3 * (i / 3)
	y := 3 * (j / 3)
	for k := 0; k < 9; k++ {
		if board[k][j] == s || board[i][k] == s || board[x+k/3][y+k%3] == s {
			return false
		}
	}
	return true
}

// @lc code=end
