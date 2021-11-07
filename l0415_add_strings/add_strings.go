package leetcode0415

import "strconv"

// Solution 1: Simulate
// Time: O(max(m, n))
// Space: O(1)
func addStrings(num1 string, num2 string) string {
	res := ""
	carry := 0

	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		sum := x + y + carry
		res = strconv.Itoa(sum%10) + res
		carry = sum / 10
	}

	return res
}
