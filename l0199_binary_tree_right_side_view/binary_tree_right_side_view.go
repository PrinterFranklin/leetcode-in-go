package leetcode0199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		length := len(q)
		res = append(res, q[len(q)-1].Val)
		for i := 0; i < length; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return res
}
