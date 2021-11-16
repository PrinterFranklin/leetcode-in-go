package leetcode0054

// Solution 1: Simulation
// Time: O(m*n)
// Space: O(1)
func spiralOrder(matrix [][]int) []int {
	res := []int{}

	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for top < bottom && left < right {
		for i := left; i < right; i++ {
			res = append(res, matrix[top][i])
		}
		for i := top; i < bottom; i++ {
			res = append(res, matrix[i][right])
		}
		for i := right; i > left; i-- {
			res = append(res, matrix[bottom][i])
		}
		for i := bottom; i > top; i-- {
			res = append(res, matrix[i][left])
		}
		top++
		bottom--
		left++
		right--
	}
	if top == bottom {
		// Note: use '<=' instead of '<'
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
	} else if left == right {
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][left])
		}
	}

	return res
}
