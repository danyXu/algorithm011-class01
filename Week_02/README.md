### 第二周总结

---

&nbsp;&nbsp;&nbsp;&nbsp;本周学习了哈希表，二叉树，堆。
* 其中哈希表的本质是使用数据的随机访问特性，在底层实现上是通过一个映射函数来定位元素下标，从而达到从key到value的快速映射。hash表的实现主要关注点有：哈希函数、装载因子、冲突解决方法、扩容方法。
    *   其中哈希函数要求计算简单、结果分布均匀、冲突概率低的特性
    *   装载因子不能太大，太大会导致后续大量的冲突，降低了查询性能，太小又比较浪费内存
    *   冲突解决方法：有开放寻址和链表法。开放寻址适合数据量小的场景，链表法适合大数据量，且链表法可以在链表长度超过一定阈值时换成红黑树或跳表等性能更优的结构
    *   扩容可以是一次性，也可以是动态的。一次性扩容遇到数据量大时，会影响性能，动态扩容更合适，可以在多次操作中慢慢迁移老的数据
    *   哈希结构的代码，感觉难点在于构造key，如何简单的构造key是一个考察点，如简单排序（nlogn），26个字母顺序构造key等等
* 二叉树和堆，都是树结构，其中树的遍历是常考的，主要有2种写法：递归和迭代。因为树本身就是递归定义的，好说，迭代则根据不同场景可以采用不同的方式，如深度遍历一般都采用栈结果，如果是广度遍历一般都采用队列
    *   递归示例：
    ```go
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
    ```
    *   迭代示例（栈）
    ```go
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
    ```
    *   迭代示例2（队列）
    ```go
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
    ```