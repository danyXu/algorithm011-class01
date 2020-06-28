### 课后作业

---

&nbsp;&nbsp;&nbsp;&nbsp;Go语言中没有直接的Queue和PriorityQueue的容器来直接使用，基础的SDK中提供的是Slice（可理解为动态数组），list（支持任意类型的双向链表）和heap（堆），其中Queue可以使用Slice或list来实现，priorityQueue可以使用heap来实现。
<br>

下面就具体来分析下这3个结构的核心方法：<br>
- Slice是一个Struct,包含一个指向底层数组的指针，一个本身的长度，一个本身的容量。其中主要看下添加元素和删除元素的方法：
```go
    type slice struct {
        array unsafe.Pointer // 通用指针类型
        len   int
        cap   int
    }
    
    // 当添加新的元素时，如果超出底层数组的cap，则会扩展size，具体的方法如下：
    func growslice(et *_type, old slice, cap int) slice {
        // ----- 忽略前面的，cap可理解为现在需要的新容量，肯定是大于现在slice的容器
        newcap := old.cap
        doublecap := newcap + newcap
        if cap > doublecap {
            // 如果新的cap大于原来2倍，直接就扩展到2倍
            newcap = cap
        } else {
            if old.len < 1024 {
                // 如果新cap没有原来2倍，且老的cap也没有1024，则直接将老的cap翻倍，在整理数量小时，扩展快些
                newcap = doublecap
            } else {
                // Check 0 < newcap to detect overflow
                // and prevent an infinite loop.
                for 0 < newcap && newcap < cap {
                    // 如果老的cap超过1024，则以1.25倍扩张
                    newcap += newcap / 4
                }
                // Set newcap to the requested cap when
                // the newcap calculation overflowed.
                if newcap <= 0 {
                    newcap = cap
                }
            }
        }
        // ---
        // roundupsize函数会做个内存对齐的相关操作，所以最后的size会有一定的变化
        capmem = roundupsize(uintptr(newcap))
        overflow = uintptr(newcap) > maxAlloc
        newcap = int(capmem)
        // ---
    }
    
    // slice的删除存在一些语法糖，底层应该还是数组的copy
    
    // 删除第i个元素，保持顺序
    a = append(a[:i], a[i+1:]...) 或 a = a[:i+copy(a[i:], a[i+1:])]
    
    // 删除第i个元素，不保持顺序
    a[i] = a[len(a)-1]    a = a[:len(a)-1]
    
    // 访问元素则直接按照下标随机访问就可以
```

- List是一个循环双向链表，有很多方法可以提供给上层来使用，如：
    1. Front(),Back():获取第一个和最后一个
    2. PushFront(v),PushBack(v),InsertBefore(v,e),InsertAfter(v,e):插入到前面或后面，插入到某个元素的前面或后面
    3. PushFrontList(l),PushBackList(l):在前面或后面插入一个list
    4. MoveToFront(e),MoveToBack(e),MoveBefore(e,mark),MoveAfter(e,mark)：移动元素...

```go
    // Element is an element of a linked list.
    type Element struct {
        // Next and previous pointers in the doubly-linked list of elements.
        // To simplify the implementation, internally a list l is implemented
        // as a Ring, such that &l.root is both the next element of the last
        // list element (l.Back()) and the previous element of the first list
        // element (l.Front()).
        next, prev *Element
    
        // The list to which this element belongs.
        list *List
    
        // The value stored with this element.
        Value interface{}
    }
    
    // Next returns the next list element or nil.
    func (e *Element) Next() *Element {
        // 因为底层会构成环，所以会多判断next是否是root本身
        if p := e.next; e.list != nil && p != &e.list.root {
            return p
        }
        return nil
    }
    
    // Prev returns the previous list element or nil.
    func (e *Element) Prev() *Element {
        // 因为底层会构成环，所以会多判断prev是否是root本身
        if p := e.prev; e.list != nil && p != &e.list.root {
            return p
        }
        return nil
    }
    
    // List represents a doubly linked list.
    // The zero value for List is an empty list ready to use.
    type List struct {
        root Element // sentinel list element, only &root, root.prev, and root.next are used
        len  int     // current list length excluding (this) sentinel element
    }
    
    // Init initializes or clears list l.
    func (l *List) Init() *List {
        l.root.next = &l.root
        l.root.prev = &l.root
        l.len = 0
        return l
    }
    
    // 下面就示例下insert和remove
    // insert inserts e after at, increments l.len, and returns e.
    func (l *List) insert(e, at *Element) *Element {
        // 正常的指针交换
        e.prev = at
        e.next = at.next
        e.prev.next = e
        e.next.prev = e
        e.list = l
        l.len++
        return e
    }
    
    // remove removes e from its list, decrements l.len, and returns e.
    func (l *List) remove(e *Element) *Element {
        // prev的next指向e的next
        e.prev.next = e.next
        // e的next的prev指向e的prev
        e.next.prev = e.prev
        e.next = nil // avoid memory leaks
        e.prev = nil // avoid memory leaks
        e.list = nil
        l.len--
        return e
    }
```

- Heap在Go语言包中，是一个interface，具体实现由开发者自己定义，因为堆是一个完全二叉树，所以底层的存储容器可以直接是一个slice，当然也可以是上面的list，并且堆的实现是最小堆，下面核心看下堆里的down和up函数，down代表下沉，up代码上浮，这是堆处理的两种常见方式，代码示例如下：
```go

package heap

import "sort"

// The Interface type describes the requirements
// for a type using the routines in this package.
// Any type that implements it may be used as a
// min-heap with the following invariants (established after
// Init has been called or if the data is empty or sorted):
//
//	!h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()
//
// Note that Push and Pop in this interface are for package heap's
// implementation to call. To add and remove things from the heap,
// use heap.Push and heap.Pop.
type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}

// Init establishes the heap invariants required by the other routines in this package.
// Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// The complexity is O(n) where n = h.Len().
func Init(h Interface) {
	// heapify
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		// 下沉是从长度为中间节点开始的，直到最上面
		down(h, i, n)
	}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func Push(h Interface, x interface{}) {
	h.Push(x)
	// 因为元素是直接放在最后，所以这里适合使用上浮逻辑
	up(h, h.Len()-1)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func Pop(h Interface) interface{} {
	n := h.Len() - 1
	// 都是与最后一个元素交换
	h.Swap(0, n)
	// 这里将最后一个元素放在了第一的位置上，适合下沉方法
	down(h, 0, n)
	return h.Pop()
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func Remove(h Interface, i int) interface{} {
	n := h.Len() - 1
	if n != i {
		// 都是与最后一个元素交换
		h.Swap(i, n)
		// 先下沉，如果下沉没有任何操作，说明后续子节点已经有序，此时执行上浮，如果下沉有操作，无需执行上浮
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.Pop()
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func Fix(h Interface, i int) {
	// 同上
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

func up(h Interface, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			// i==j只在j=i=0时即第一个节点时才可能，有序也就不用上浮
			break
		}
		h.Swap(i, j)
		// 交换更新变量j为父节点
		j = i
	}
}

func down(h Interface, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
		    // 已经到达最后的叶子节点
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			// 交换为右节点，如果右节点比较小
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	// 如果一开始就满足要求，则无需任何操作
	return i > i0
}


```
