package leetcode0042

// Best Solution: Two Pointers
// Time: O(N)
// Space: O(1)
func trap(height []int) int {
	res, left, right, lmax, rmax := 0, 0, len(height) - 1, 0, 0
	for left <= right {
		// key point: make sure we are on the lower side
		if height[left] <= height[right] {
			if height[left] > lmax {
				lmax = height[left]
			} else {
				res += lmax - height[left]
			}
			left++
		} else {
			if height[right] > rmax {
				rmax = height[right]
			} else {
				res += rmax - height[right]
			}
			right--
		}
	}
	return res
}

// Solution 2: Brute Force
// Time: O(N^2)
// Space: O(1)
func trap1(height []int) int {
	res := 0
	for i, v := range height {
		lmax, rmax := 0, 0
		for _, w := range height[i:] {
			rmax = max(rmax, w)
		}
		for _, w := range height[:i+1] {
			lmax = max(lmax, w)
		}
		res += min(lmax, rmax) - v
	}
	return res
}

// Solution 3: DP
// Time: O(N)
// Space: O(N)
func trap2(height []int) int {
	res, n := 0, len(height)

	lmax := make([]int, n)
	lmax[0] = height[0]
	for i := 1; i < n; i++ {
		lmax[i] = max(lmax[i-1], height[i])
	}

	rmax := make([]int, n)
	rmax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rmax[i] = max(rmax[i+1], height[i])
	}

	for i, v := range height {
		res += min(lmax[i], rmax[i]) - v
	}

	return res
}

// Solution 4: Single Stack
// Time: O(N)
// Space: O(N)
func trap3(height []int) int {
	res, stack := 0, []int{}
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			// pop out from stack
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curWidth := i - left - 1
			curHeight := min(height[left], h) - height[top]
			res += curWidth * curHeight
		}
		stack = append(stack, i)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}