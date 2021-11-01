package leetcode0322

import "math"

// Solution 1: DP
// Time: O(Sn)
// Space: O(S)
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		/* can also use amount+1 */
		dp[i] = math.MaxInt32
		for _, v := range coins {
			if v <= i {
				dp[i] = min(dp[i], dp[i-v]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
