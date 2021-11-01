package leetcode0575

// Solution 1: Greedy
// Time: O(n)
// Space: O(n)
func distributeCandies(candyType []int) int {
	// use map as set, struct{}{} doesn't cost space
	set := map[int]struct{}{}
	for _, v := range candyType {
		set[v] = struct{}{}
	}
	ans := len(candyType) / 2
	if len(set) < ans {
		ans = len(set)
	}
	return ans
}
