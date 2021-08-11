package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var tree *TreeNode
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			tree = &TreeNode{
				Val:   inorder[i],
				Left:  buildTree(preorder[1:i+1], inorder[0:i]),
				Right: buildTree(preorder[i+1:], inorder[i+1:]),
			}
		}
	}
	return tree

}
