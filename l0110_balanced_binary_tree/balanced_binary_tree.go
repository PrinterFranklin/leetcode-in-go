package leetcode0110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Solution 1: Recursion
func isBalanced1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBalanced1(root.Left) || !isBalanced1(root.Right) {
		return false
	}
	leftHight := depth(root.Left)
	RightHight := depth(root.Right)
	return abs(leftHight-RightHight) <= 1
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}

// Solution 2: Recursion + Hash Table (Time Optimized)
func isBalanced2(root *TreeNode) bool {
	height := make(map[*TreeNode]int)
	return isBalancedHash(root, height)
}

func isBalancedHash(root *TreeNode, height map[*TreeNode]int) bool {
	if root == nil {
		height[root] = 0
		return true
	}
	if !isBalancedHash(root.Left, height) || !isBalancedHash(root.Right, height) {
		return false
	}
	if abs(height[root.Left]-height[root.Right]) > 1 {
		return false
	}
	height[root] = max(height[root.Left], height[root.Right]) + 1
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
