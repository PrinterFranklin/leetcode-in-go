package leetcode0543

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	diameter := 0
	depth(root, &diameter)
	return diameter
}

func depth(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left, diameter)
	right := depth(root.Right, diameter)
	*diameter = max(left+right, *diameter)
	return max(left, right) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
