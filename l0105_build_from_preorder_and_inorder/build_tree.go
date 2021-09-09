package leetcode0105

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	pivot := 0
	for ; pivot < len(inorder); pivot++ {
		if inorder[pivot] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:1+pivot], inorder[:pivot])
	root.Right = buildTree(preorder[1+pivot:], inorder[pivot+1:])

	return root
}