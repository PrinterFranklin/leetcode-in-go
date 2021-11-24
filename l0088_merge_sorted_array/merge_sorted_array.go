package leetcode0088

// Solution 1: Reverse Two Pointers
// Time: O(m+n)
// Space: O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j, tail := m-1, n-1, m+n-1; i >= 0 || j >= 0; tail-- {
		if i < 0 {
			nums1[tail] = nums2[j]
			j--
		} else if j < 0 {
			nums1[tail] = nums1[i]
			i--
		} else if nums1[i] > nums2[j] {
			nums1[tail] = nums1[i]
			i--
		} else {
			nums1[tail] = nums2[j]
			j--
		}
	}
}

// Solution 2: Tmp Array
// Time: O(m+n)
// Space: O(m+n)
func merge1(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	i, j := 0, 0
	for {
		// append the remaining part to tail
		if i == m {
			sorted = append(sorted, nums2[j:]...)
			break
		}
		if j == n {
			sorted = append(sorted, nums1[i:]...)
			break
		}
		// add the smallest element one by one
		if nums1[i] < nums2[j] {
			sorted = append(sorted, nums1[i])
			i++
		} else {
			sorted = append(sorted, nums2[j])
			j++
		}
	}
	copy(nums1, sorted)
}
