package leetcode0300

// Solution 1: DP
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	res := 1
	for i := 1; i < len(nums); i++ {
		max := 0
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				continue
			}
			if dp[j] > max {
				max = dp[j]
			}
		}
		dp[i] = max + 1
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

// Solution 2: DP + Greedy + Binary Search
func lengthOfLIS1(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	dp := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] > dp[len(dp)-1] {
			dp = append(dp, nums[i])
		} else if nums[i] < dp[len(dp)-1] {
			// binary search viriation:
			// find the smallest element which is bigger than the given element
			l, r := 0, len(dp)-1
			for l < r {
				m := l + (r-l)/2
				if dp[m] >= nums[i] {
					r = m
				} else {
					l = m+1
				}
			}
			dp[l] = nums[i]
		}
	}
	return len(dp)
}