### 第一周总结

---

&nbsp;&nbsp;&nbsp;&nbsp;第一周课程内容整理来说比较简单，数组、队列、栈、链表都是比较基础的线性结构，但经过一些习题的练习，发现一些基本的思路和技巧，现整理如下：
> 数组：<br>
&nbsp;&nbsp;&nbsp;&nbsp;数组的查找和修改的时间复杂度是o(1)，删除和添加（指定位置）都是o(n)，因为数组要维护数据的顺序排列，需要移动部分数据。<br>
&nbsp;&nbsp;&nbsp;&nbsp;在解决具体的数组习题时，发现了一些比较不错的思路或代码片段整理如下：<br>
&nbsp;&nbsp;- 双指针法（快慢或前后）和数组翻转
```go
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	i, j := 0, 1 // i,j 两个指针
	for ; j <= len(nums)-1; j++ {
		if nums[j] != nums[j-1] {
			if j-i > 1 {
				// 如果非相邻则复制
				nums[i+1] = nums[j]
			}
			i++
		}
	}

	return i + 1
}

// https://leetcode-cn.com/problems/rotate-array/
func reverse(nums []int, start, end int) {
	// 翻转数组
	for start < end {
		nums[end], nums[start] = nums[start], nums[end]
		start++
		end--
	}
}

```
> 链表：<br>
&nbsp;&nbsp;&nbsp;&nbsp;链表的添加和删除的时间复杂度是o(1)，查找的（指定位置）都是o(n)，因为链表的数据存储是非连续的，得一个个的找。为了加快链表的查找速度，产生了一个新的数据结构跳表，核心思想是建立索引加快查找，也是一个常用的空间换时间的方式。跳表的查找时间复杂度为o(logn)<br>
>> // 跳表的时间复杂度分析:
>> 建设我们有n个节点，每K个节点取一个索引，因为最上层就k个节点，因为第一层是n个节点，第二层n/k，第三次是n/k^2,第四层是n/k^3,那么第h层就是n/k^(h-1),所以就有n/k^(h-1)=k,即h=log(k)n,又因为每层至多遍历k次,所以跳表的时间复杂度就为k*log(k)n=o(logn)


> &nbsp;&nbsp;&nbsp;&nbsp;在解决具体的链表习题时，发现了一些比较不错的思路或代码片段整理如下：<br>
&nbsp;&nbsp;- 递归和迭代是常用的链表处理模式，递归很大程度上的空间复杂度会高一些，掌握熟练的递归写法，经常可以写出一些比较精简高效看起来舒服的代码
```go
// https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 递归写法
    	switch {
	case l1 == nil:
		return l2
	case l2 == nil:
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next,l2)
		return l1
	}
	
	l2.Next = mergeTwoLists(l1,l2.Next)
	return l2
}

// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	// 判断边界
	if head == nil || head.Next == nil {
		return head
	}

	result := head.Next
	p := new(ListNode)
	// 循环交换，迭代处理
	for head != nil && head.Next != nil {
		next := head.Next.Next
		p.Next = head.Next
		head.Next.Next = head
		head = next
		p = p.Next.Next
	}
	
	p.Next = head
	return result
}

```    
> 栈或队列：<br>
&nbsp;&nbsp;&nbsp;&nbsp;栈或队列是一个按某个规则进行插入和删除的，如先进后出，先进先出，栈考察的应该会多些，从做题的感觉来看，第一某些抽象的题目比较难发现最近相关性而忽略掉使用栈，第二单独使用一个栈的方法应该比较少，更多是使用双栈或双队列来解决问题，这部分练习的不多，暂且无特别的发现<br>