package main

/*
输入两棵二叉树A，B，判断B是不是A的子结构。（ps：我们约定空树不是任意一个树的子结构）

输入：{8,8,#,9,#,2,#,5},{8,9,#,2}
返回值：true

*/

func HasSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	if pRoot1 == nil || pRoot2 == nil {
		return false
	}
	return isSubTree(pRoot1, pRoot2) || HasSubtree(pRoot1.Left, pRoot2) || HasSubtree(pRoot1.Right, pRoot2)

}

func isSubTree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	// write code here
	if pRoot2 == nil {
		return true
	}
	if pRoot1 == nil {
		return false
	}
	if pRoot1.Val != pRoot2.Val {
		return false
	}
	return isSubTree(pRoot1.Left, pRoot2.Left) && isSubTree(pRoot1.Right, pRoot2.Right)
}
