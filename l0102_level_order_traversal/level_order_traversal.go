package leetcode0102

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	// a queue to store TreeNode
	q := []*TreeNode{root}
	for len(q) > 0 {
		level := []int{}
		length := len(q)
		for i := 0; i < length; i++ {
			// pop node out of queue head
			node := q[0]
			q = q[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
