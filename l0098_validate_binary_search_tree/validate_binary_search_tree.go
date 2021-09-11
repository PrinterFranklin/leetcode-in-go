package leetcode0098

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Solution 1: recursion
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

// Solution 2: inorder traversal
func isValidBST2(root *TreeNode) bool {
	pre := math.MinInt64
	var helper func(root *TreeNode) bool
	helper = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if !helper(root.Left) {
			return false
		}
		if root.Val <= pre {
			return false
		}
		pre = root.Val
		return helper(root.Right)
	}
	return helper(root)
}
