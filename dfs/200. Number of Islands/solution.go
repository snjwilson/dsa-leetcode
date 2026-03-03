package main

func numIslands(grid [][]byte) int {
	islands := 0

	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == '0' || grid[i][j] == 'V' {
				continue
			}
			dfs(i, j, grid)
			islands++
		}
	}
	return islands
}

func dfs(row, col int, grid [][]byte) {
	if row > len(grid)-1 || col > len(grid[0])-1 || row < 0 || col < 0 {
		return
	}

	if grid[row][col] == '0' || grid[row][col] == 'V' {
		return
	}
	grid[row][col] = byte('V')
	dfs(row, col+1, grid)
	dfs(row+1, col, grid)
	dfs(row, col-1, grid)
	dfs(row-1, col, grid)
}
