// 判断一颗二叉树是不是对称的  先思考对称树的特性
// 遍历两棵树，想想为什么分两颗树遍历，而不是一颗？
type TreeNode struct {
	val         int
	left, right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return Symmetric(root.root)
}
func Symmetric(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1!=nil ||t2！=nil {
		return false
	}
	return t1.val==t2.val && Symmetric(t1.left,t2.right) && Symmetric(t1.right,t2.left)
}