package leetcode0235

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// Solution 1: Iteration
func lowestCommonAncestorBST1(root, p, q *TreeNode) *TreeNode {
	for (root.Val - p.Val) * (root.Val - q.Val) > 0 {
		if p.Val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

// Solution 2: Recursion
func lowestCommonAncestorBST2(root, p, q *TreeNode) *TreeNode {
	if (root.Val - p.Val) * (root.Val - q.Val) <= 0 {
		return root
	}
	if (p.Val < root.Val) {
		return lowestCommonAncestorBST2(root.Left, p, q)
	} else {
		return lowestCommonAncestorBST2(root.Right, p, q)
	}
}