package leetcode0022

// Solution 1: Backtrace + Pruning
// Time: O(4^n)
// Space: O(n)
func generateParenthesis(n int) []string {
	res := []string{}
	backtrack(n, n, "", &res)
	return res
}

// @l: the remaining number of left parentheses
// @r: the remaining number of right parentheses
// @str: current parentheses string
// @res: the final result slice
func backtrack(l, r int, str string, res *[]string) {
	if l == 0 && r == 0 {
		*res = append(*res, str)
		return
	}
	// we can always try to add more left parentheses
	if l > 0 {
		backtrack(l-1, r, str+"(", res)
	}
	// we can only add right parentheses when we have more left ones
	if r > 0 && l < r {
		backtrack(l, r-1, str+")", res)
	}
}
