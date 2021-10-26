package leetcode0055

// Solution 1: DP + Backtrace
// Time: O(n^2)
// Space: O(n)
func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && (nums[j] >= i-j) {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(nums)-1]
}

// Solution 2: Greedy
// Time: O(n)
// Space: O(1)
func canJump1(nums []int) bool {
	maxDistance := nums[0]
	for i := 1; i < len(nums); i++ {
		if maxDistance >= len(nums)-1 {
			return true
		} else if i > maxDistance {
			return false
		} else if i+nums[i] > maxDistance {
			maxDistance = i + nums[i]
		}
	}
	return true
}
