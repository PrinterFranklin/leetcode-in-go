package leetcode0062

func uniquePaths(m int, n int) int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				res[i][j] = 1
			} else {
				res[i][j] = res[i-1][j] + res[i][j-1]
			}
		}
	}
	return res[m-1][n-1]
}
