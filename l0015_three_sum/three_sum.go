package leetcode0015

import "sort"

// Solution 1: Two Pointers
// Time: O(n^2)
// Space: O(1)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		// if the smallest element is bigger than 0, sum can't be 0
		if nums[i] > 0 {
			break
		}
		// ignore duplicate cases
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		// increase left and decrease right
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			} else if sum > 0 {
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else {
				tmp := []int{nums[i], nums[left], nums[right]}
				res = append(res, tmp)
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}
	return res
}
