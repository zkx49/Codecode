// 二叉树的镜像
/* type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
} */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return
	}
	root.left, root.right = root.right, root.left
	invertTree(root.left)
	invertTree(root.right)
	return root
}