package leetcode0005

// Solution 1: Expand around center
// Time: O(n^2)
// Space: O(1)
func longestPalindrome(s string) string {
	// to return a substring, we must get the start and end indexes
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		// Case 1: single center
		l1, r1 := expand(s, i, i)
		// Case 2: double center
		l2, r2 := expand(s, i, i+1)
		if r1-l1 > end-start {
			end = r1
			start = l1
		}
		if r2-l2 > end-start {
			end = r2
			start = l2
		}
	}
	return s[start : end+1]
}

func expand(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left, right = left-1, right+1
	}
	return left + 1, right - 1
}
