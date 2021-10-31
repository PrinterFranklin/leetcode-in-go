package leetcode0718

// Solution 1: DP
// Time: O(mn)
// Space: O(mn)
func findLength1(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	res := 0
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}

// Solution 2: DP (optimized)
// Time: O(mn)
// Space: O(n)
func findLength2(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	res := 0
	dp := make([]int, n+1)
	for i := 1; i <= m; i++ {
		for j := n; j >= 1; j-- {
			if nums1[i-1] == nums2[j-1] {
				dp[j] = dp[j-1] + 1
				// in the optimized solution, we must set zero here
			} else {
				dp[j] = 0
			}
			if dp[j] > res {
				res = dp[j]
			}
		}
	}
	return res
}

// Solution 3: Sliding Window
// Time: O((m+n)*min(m,n))
// Space: O(1)
func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	res := 0
	// i is the right moving distance of nums1
	for i := 0; i < m; i++ {
		// len is the length of current subarray
		len := min(m, n-i)
		maxLen := maxLength(nums1, nums2, 0, i, len)
		res = max(res, maxLen)
	}
	// j is the right moving distance of nums2
	for j := 0; j < n; j++ {
		len := min(n, m-j)
		maxLen := maxLength(nums1, nums2, j, 0, len)
		res = max(res, maxLen)
	}
	return res
}

func maxLength(nums1, nums2 []int, offset1, offset2, len int) int {
	res, k := 0, 0
	for i := 0; i < len; i++ {
		if nums1[offset1+i] == nums2[offset2+i] {
			k++
		} else {
			k = 0
		}
		res = max(res, k)
	}
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
