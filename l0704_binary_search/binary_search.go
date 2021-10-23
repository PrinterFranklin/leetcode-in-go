package leetcode0704

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	// quit when the interval is empty, use '<=' instead of '<'
	for lo <= hi {
		// prevent integer overflow
		// better performance: mid := lo + ((hi - lo) >> 1)
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}
