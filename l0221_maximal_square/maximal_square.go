package leetcode0221

// Solution 1: DP
// Time: O(mn)
// Space: O(mn)
func maximalSquare(matrix [][]byte) int {
	// dp: max side value of the rectangle
	dp := make([][]int, len(matrix))
	maxSide := 0
	// calculating max side for every rectangle
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				}
			}
			if dp[i][j] > maxSide {
				maxSide = dp[i][j]
			}
		}
	}
	return maxSide * maxSide
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
