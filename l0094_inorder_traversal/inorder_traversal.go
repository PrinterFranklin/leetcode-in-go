package leetcode0094

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	inorder(root, &result)
	return result
}

func inorder(node *TreeNode, result *[]int) {
	if node != nil {
		inorder(node.Left, result)
		*result = append(*result, node.Val)
		inorder(node.Right, result)
	}
}
