package leetcode0003

// Solution 1: Sliding Window (hash table)
// Time: O(n)
// Space: O(N)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	res, left, right := 0, 0, 0
	// init hash map (char -> index)
	set := map[byte]int{}
	for right < len(s) {
		if _, ok := set[s[right]]; ok {
			// if char is in the set, move left index
			if set[s[right]] >= left {
				left = set[s[right]] + 1
			}
		}
		// add or refresh the index
		set[s[right]] = right
		// record the max length
		res = max(right-left+1, res)
		right++
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// Solution 2: Sliding Window (bit array, AC)
// Time: O(n)
// Space: O(N)
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	res, left, right := 0, 0, 0
	// there are 256 ASCII chars in total
	var bitSet [256]uint8
	for left < len(s) {
		if right < len(s) && bitSet[s[right]] == 0 {
			bitSet[s[right]] = 1
			right++
		} else {
			bitSet[s[left]] = 0
			left++
		}
		res = max(res, right-left)
	}
	return res
}
