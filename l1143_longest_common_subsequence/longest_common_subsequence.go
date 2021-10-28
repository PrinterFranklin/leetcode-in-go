package leetcode1143

// Solution 1: DP
// Time: O(mn)
// Space: O(mn)
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	// init dp matrix as (m+1) * (n+1)
	// dp[i][0] and dp[0][j] (0<=i<=m, 0<=j<=n) will be 0
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// DP formula
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[m][n]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
