package leetcode0070

// Solution 0: recursion (too slow)
func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// Solution 1: DP
// Time: O(N)
// Space: O(1)
func climbStairs1(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	p, q, r := 0, 1, 2
	for i := 0; i < n-2; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
