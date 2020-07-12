package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 只有当根节点时才返回
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	// 看左右节点的存在情况
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 左右都存在
	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}

	return left
}

func main() {

}
