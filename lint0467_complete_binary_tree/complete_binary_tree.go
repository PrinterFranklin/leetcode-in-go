package lintcode0467

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
}

// Solution 1: level order traversal
// Time: O(n)
// Space: O(n)
func isCompleteTree(root *TreeNode) bool {
	q := []*TreeNode{root}
	for i := 0; i < len(q); i++ {
		node := q[len(q)-1]
		if node == nil {
			continue
		}
		q = append(q, node.Left)
		q = append(q, node.Right)
	}
	// remove nil nodes in tail
	for len(q) > 0 && q[len(q)-1] == nil {
		q = q[:len(q)-1]
	}
	if len(q) == 0 {
		return true
	}
	for i := 0; i < len(q); i++ {
		if q[i] == nil {
			return false
		}
	}
	return true
}
