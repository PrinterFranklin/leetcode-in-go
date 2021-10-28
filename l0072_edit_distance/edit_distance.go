package leetcode0072

// Solution 1: DP
// Time: O(mn)
// Space: O(mn)
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	// init dp matrix
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// init base case
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	// fill the whole dp matrix
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

func min(args ...int) int {
	minValue := args[0]
	for _, arg := range args {
		if arg < minValue {
			minValue = arg
		}
	}
	return minValue
}

// Solution 2: Recursion (too slow)
func minDistance1(word1 string, word2 string) int {
	return dp(word1, word2, len(word1), len(word2))
}

func dp(word1 string, word2 string, i, j int) int {
	if i == 0 {
		return j
	} else if j == 0 {
		return i
	} else if word1[i-1] == word2[j-1] {
		return dp(word1, word2, i-1, j-1)
	} else {
		m1 := dp(word1, word2, i-1, j)
		m2 := dp(word1, word2, i, j-1)
		m3 := dp(word1, word2, i-1, j-1)
		return min(m1, m2, m3) + 1
	}
}
