package leetcode0046

// Solution 1: Backtrack
// Time: O(n*n!)
// Space: O(n)
func permute(nums []int) [][]int {
	res := [][]int{}
	track := make([]int, len(nums))
	cur := []int{}
	backtrack(&res, nums, track, cur)
	return res
}

func backtrack(res *[][]int, nums []int, track []int, cur []int) {
	n := len(nums)
	if len(cur) == n {
		// must copy the cur array
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i, v := range nums {
		if track[i] == 1 {
			continue
		}
		cur = append(cur, v)
		track[i] = 1
		backtrack(res, nums, track, cur)
		cur = cur[:len(cur)-1]
		track[i] = 0
	}
}
