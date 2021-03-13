// 二叉树的深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.left), maxDepth(root.right)) + 1
}
func max(i, j) int {
	if i > j {
		return i
	}
	return j
}