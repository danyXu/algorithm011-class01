package main

// 也是两种实现方法，递归和迭代

type Node struct {
	Val      int
	Children []*Node
}

func dfs(root *Node, level int, res *[][]int) {
	if root == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, []int{})
	}
	(*res)[level] = append((*res)[level], root.Val)
	for _, n := range root.Children {
		dfs(n, level+1, res)
	}
}

func levelOrderWithRecursive(root *Node) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	dfs(root, 0, &res)
	return res
}

func levelOrderWithIteration(root *Node) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int

	// 做广度搜索参考队列， 做深度搜索参考栈
	queue := []*Node{root}
	for len(queue) != 0 {
		var levelResult []int
		var tmp []*Node
		for _, n := range queue {
			levelResult = append(levelResult, n.Val)
			tmp = append(tmp, n.Children...)
		}

		queue = queue[:0]
		result = append(result, levelResult)
		if len(tmp) > 0 {
			queue = tmp
		}

	}

	return result
}

func main() {

}
