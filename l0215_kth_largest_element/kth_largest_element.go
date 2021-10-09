package leetcode0215

import "math/rand"

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, lo, hi, index int) int {
	p := randomPartition(nums, lo, hi)
	if p == index {
		return nums[p]
	} else if p < index {
		return quickSelect(nums, p+1, hi, index)
	} else {
		return quickSelect(nums, lo, p-1, index)
	}
}

func randomPartition(nums []int, lo, hi int) int {
	r := rand.Intn(hi-lo+1) + lo
	nums[r], nums[hi] = nums[hi], nums[r]
	return partition(nums, lo, hi)
}

func partition(nums []int, lo, hi int) int {
	pivot := nums[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if nums[j] < pivot {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	nums[hi], nums[i] = nums[i], nums[hi]
	return i
}
