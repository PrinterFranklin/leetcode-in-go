package leetcode0470

import "math/rand"

func rand7() int {
	return rand.Intn(7) + 1
}

// Solution 1: Reject Sampling
func rand10() int {
	for {
		r1, r2 := rand7(), rand7()
		// idx = rand49()
		idx := (r1-1)*7 + r2
		// reject sampling to get rand40()
		if idx <= 40 {
			//
			return idx%10 + 1
		}
	}
}

// Solution 2: Reject Sampling (optimized)
func random10() int {
	for {
		r1, r2 := rand7(), rand7()
		idx := (r1-1)*7 + r2
		if idx <= 40 {
			return idx%10 + 1
		}
		r1, r2 = idx-40, rand7()
		idx = (r1-1)*7 + r2
		if idx <= 60 {
			return idx%10 + 1
		}
		r1, r2 = idx-60, rand7()
		idx = (r1-1)*7 + r2
		if idx <= 20 {
			return idx%10 + 1
		}
	}
}
