package leetcode0053

// Solution 1: DP
// Time: O(n)
// Space: O(1)
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// Solution 2: Greedy
// Time: O(n)
// Space: O(1)
func maxSubArray1(nums []int) int {
	max, count := -10000, 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > max {
			max = count
		}
		if count < 0 {
			count = 0
		}
	}
	return max
}
