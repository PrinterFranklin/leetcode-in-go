package leetcode0435

import "sort"

// Solution 1: Greedy (recommended)
// Time: O(nlogn)
// Space: O(logn)
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	// sort the arrays by the right limit
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	res, right := 1, intervals[0][1]
	for _, p := range intervals[1:] {
		if p[0] >= right {
			res++
			right = p[1]
		}
	}
	return len(intervals) - res
}
