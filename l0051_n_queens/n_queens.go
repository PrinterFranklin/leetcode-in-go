package leetcode0051

import "strings"

// Solution 1: Backtrace + Pruning
// Time: O(N!)
// Space: O(N^2)
func solveNQueens(n int) [][]string {
	res := [][]string{}
	// init board as string slice (so that we can write to it)
	board := make([][]string, n)
	for i := range board {
		board[i] = make([]string, n)
		for j := range board[i] {
			board[i][j] = "."
		}
	}
	backtrack(&res, board, 0)
	return res
}

func backtrack(res *[][]string, board [][]string, row int) {
	n := len(board)
	// Reach terminate condition, add board to res
	if row == n {
		tmp := make([]string, n)
		for i := 0; i < n; i++ {
			tmp[i] = strings.Join(board[i], "")
		}
		*res = append(*res, tmp)
		return
	}
	for col := 0; col < n; col++ {
		// Pruning: test if this is a good move
		if !isValid(board, row, col) {
			continue
		}
		// make a choice
		board[row][col] = "Q"
		backtrack(res, board, row+1)
		// rollback the choice
		board[row][col] = "."
	}
}

func isValid(board [][]string, row, col int) bool {
	n := len(board)
	// Detect column collision
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	// Detect upper left collision
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	// Detect upper right collision
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}
