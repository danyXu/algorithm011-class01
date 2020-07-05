package main

// 二叉树中序遍历，递归方法：使用栈来实现 时间复杂度为：o(n)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归实现
func inorderTraversalWithRecursive(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	data := make([]int, 0)
	data = append(data, inorderTraversal(root.Left)...)
	data = append(data, root.Val)
	data = append(data, inorderTraversal(root.Right)...)
	return data
}

// 循环实现
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var stack []*TreeNode
	var result []int

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		result = append(result, root.Val)
		stack = stack[:len(stack)-1]
		root = root.Right
	}

	return result
}

func main() {

}
