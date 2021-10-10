package leetcode0034

func searchRange(nums []int, target int) []int {
	lo, hi := 0, len(nums)-1
	// range indexes
	l, r := -1, -1
	for lo <= hi {
		mid := (hi-lo)/2 + lo
		if nums[mid] == target {
			l, r = mid, mid
			break
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	// search around mid
	if l != -1 {
		for i := l; i > -1 && nums[i] == target; i-- {
			l = i
		}
		for i := r; i < len(nums) && nums[i] == target; i++ {
			r = i
		}
	}
	return []int{l, r}
}
