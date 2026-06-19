func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	rows := len(grid)
	cols := len(grid[0])

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i >= rows || j >= cols || grid[i][j] == 0 {
			return 0
		}

		grid[i][j] = 0

		return 1 +
			dfs(i-1, j) +
			dfs(i+1, j) +
			dfs(i, j-1) +
			dfs(i, j+1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				area := dfs(r, c)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}