### 第四周总结

---

&nbsp;&nbsp;&nbsp;&nbsp;本周学习了广度优先搜索和深度优先搜索、贪心和二分查找算法。
* 其中广度优先和深度优先算法使用较多，很多问题既可以用广度也可以深度，但存在一个最合适的，具体来说深度是递归来实现，通过系统的调用栈来实现回溯，而广度优先搜索则是通过队列来实现
* 贪心算法是一种特殊的最优算法，通过局部最优就能推导出全局最优，如哈夫曼编码、最短路径以及最小生成树等
* 二分查找，有3个条件：
*       单调递增
*       通过索引访问
*       有界


&nbsp;&nbsp;&nbsp;&nbsp;下面示例下经典的广度和深度解决一个问题：
*   ```go
    // 429 N叉树的层序遍历：深度搜索
    func dfs(root *Node, level int, res *[][]int){
        if root == nil{
            return
        }
        if len(*res) == level{
            *res = append(*res, []int{})
        }
        (*res)[level] = append((*res)[level], root.Val)
        for _, n := range root.Children{
            dfs(n, level+1, res)
        }
    }
    
    func levelOrder(root *Node) [][]int {
       if root == nil {
           return nil
       }
    
       res := make([][]int,0)
       dfs(root,0,&res)
       return res
    }

    // 广度搜索
    func levelOrder(root *Node) [][]int {
        if root == nil {
    		return nil
    	}
    
    	var result [][]int
    
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
    
    ```