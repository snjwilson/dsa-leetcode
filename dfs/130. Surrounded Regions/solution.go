package main

func solve(board [][]byte) {
	for i := range board {
		dfs(i, 0, board)
		dfs(i, len(board[0])-1, board)
	}
	for j := range board[0] {
		dfs(0, j, board)
		dfs(len(board)-1, j, board)
	}

	for i := range board {
		for j := range board[0] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'F' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(row, col int, board [][]byte) {
	if row < 0 || col < 0 || row > len(board)-1 || col > len(board[0])-1 || board[row][col] == 'X' || board[row][col] == 'F' {
		return
	}
	board[row][col] = 'F'
	dfs(row, col-1, board)
	dfs(row-1, col, board)
	dfs(row, col+1, board)
	dfs(row+1, col, board)
}
