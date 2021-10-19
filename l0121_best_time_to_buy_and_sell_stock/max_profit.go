package leetcode0121

// Solution 1: Greedy
func maxProfit(prices []int) int {
	min, max := int(^uint(0)>>1), 0
	for _, v := range prices {
		// min is the minimum value in left region
		if v < min {
			min = v
		}
		// max is the maximum profit
		if v-min > max {
			max = v - min
		}
	}
	return max
}
