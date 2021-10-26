package leetcode0198

// Solution 1: DP
// Time: O(n)
// Space: O(1)
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	first := nums[0]
	second := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
