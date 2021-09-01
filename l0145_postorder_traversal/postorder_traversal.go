package leetcode0145

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	postorder(root, &result)
	return result
}

func postorder(node *TreeNode, result *[]int) {
	if node != nil {
		postorder(node.Left, result)
		postorder(node.Right, result)
		*result = append(*result, node.Val)
	}
}
