package leetcode0045

// Solution 1: DP
// Time: O(n^2)
// Space: O(n)
func jump(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 0
	for i := 1; i < len(nums); i++ {
		minTimes := 10000
		for j := 0; j < i; j++ {
			if j + nums[j] >= i && dp[j] + 1 < minTimes {
				minTimes = dp[j] + 1
			}
		}
		dp[i] = minTimes
	}
	return dp[len(nums)-1]
}

// Solution 2: Greedy
// Time: O(n)
// Space: O(1)
func jump1(nums []int) int {
	// end is the current right limit
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < len(nums) - 1; i++ {
		if i + nums[i] > maxPosition {
			maxPosition = i + nums[i]
		}
		// every time we refresh maxPosition, we add to steps
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}