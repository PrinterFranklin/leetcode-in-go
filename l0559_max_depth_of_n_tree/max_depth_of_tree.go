package leetcode0559

type TreeNode struct {
	Val      int
	Children []*TreeNode
}

// Solution 1: BFS
// Time: O(n)
// Space: O(n)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	res := 0
	for len(q) > 0 {
		length := len(q)
		for i := 0; i < length; i++ {
			node := q[0]
			q = q[1:]
			for _, child := range node.Children {
				q = append(q, child)
			}
		}
		res++
	}
	return res
}

// Solution 2: DFS
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	for _, child := range root.Children {
		res = max(res, maxDepth1(child))
	}
	return res + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
