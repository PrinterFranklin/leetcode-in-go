package leetcode0674

// Solution 1: Greedy
// Time: O(n)
// Space: O(1)
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	max, start := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			start = i
		}
		if i-start+1 > max {
			max = i - start + 1
		}
	}
	return max
}
