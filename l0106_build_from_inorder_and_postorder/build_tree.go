package leetcode0106

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{postorder[len(postorder)-1], nil, nil}
	pivot := 0
	for ; pivot < len(inorder); pivot++ {
		if inorder[pivot] == postorder[len(postorder)-1] {
			break
		}
	}
	root.Left = buildTree(inorder[:pivot], postorder[:pivot])
	root.Right = buildTree(inorder[pivot+1:], postorder[pivot:len(postorder)-1])

	return root
}
