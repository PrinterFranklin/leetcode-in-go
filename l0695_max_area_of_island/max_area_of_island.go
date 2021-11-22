package leetcode0695

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, dfs(grid, i, j))
			}
		}
	}
	return maxArea
}

func dfs(grid [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	area := 1
	grid[i][j] = 0
	area += dfs(grid, i-1, j)
	area += dfs(grid, i, j-1)
	area += dfs(grid, i+1, j)
	area += dfs(grid, i, j+1)
	return area
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}