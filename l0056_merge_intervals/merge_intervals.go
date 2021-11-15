package leetcode0056

import "sort"

// Solution 1: Sort + Greedy
// Time: O(nlogn)
// Space: O(nlogn) from sort
func merge(intervals [][]int) [][]int {
	// Step 1: sort by first element
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	prev := intervals[0]
	// Step 2: merge one by one
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] {
			res = append(res, prev)
			prev = cur
		} else {
			prev[1] = max(prev[1], cur[1])
		}
	}
	res = append(res, prev)
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
