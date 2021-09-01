package leetcode0144

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	preorder(root, &result)
	return result
}

func preorder(node *TreeNode, result *[]int) {
	if node != nil {
		*result = append(*result, node.Val)
		preorder(node.Left, result)
		preorder(node.Right, result)
	}
}
