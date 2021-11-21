package leetcode0008

import "math"

func myAtoi(s string) int {
	abs, sign, i, n := 0, 1, 0, len(s)
	// ignore the beginning empty char
	for i < n && s[i] == ' ' {
		i++
	}
	// looking for positive and negative char
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}
	// calculate the sum
	for i < n && s[i] >= '0' && s[i] <= '9' {
		abs = 10*abs + int(s[i]-'0')
		if sign*abs < math.MinInt32 {
			return math.MinInt32
		} else if sign*abs > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return sign * abs
}
