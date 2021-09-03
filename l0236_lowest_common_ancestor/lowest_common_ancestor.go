package leetcode0236

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor1(root.Left, p, q)
	right := lowestCommonAncestor1(root.Right, p, q)
	if left != nil {
		if right != nil {
			return root
		}
		return left
	}
	return right
}
