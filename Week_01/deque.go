package main

import "fmt"

// 基本上底层是使用循环单链表来实现
type Element struct {
	Pre, Next *Element
	Value     int
}

//leetcode submit region begin(Prohibit modification and deletion)
type MyCircularDeque struct {
	Root *Element
	Len  int
	Cap  int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	root := Element{
		Value: -1,
	}

	root.Pre = &root
	root.Next = &root
	return MyCircularDeque{
		Root: &root,
		Len:  0,
		Cap:  k,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.Len == this.Cap {
		return false
	}

	// 交换
	after := this.Root.Next
	currentEle := Element{
		Pre:   this.Root,
		Next:  after,
		Value: value,
	}

	this.Root.Next = &currentEle
	after.Pre = &currentEle

	this.Len++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.Len == this.Cap {
		return false
	}

	// 交换
	before := this.Root.Pre
	currentEle := Element{
		Pre:   before,
		Next:  this.Root,
		Value: value,
	}

	this.Root.Pre = &currentEle
	before.Next = &currentEle

	this.Len++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.Len == 0 {
		return false
	}

	this.Root.Next.Next.Pre = this.Root
	this.Root.Next = this.Root.Next.Next
	this.Len--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.Len == 0 {
		return false
	}

	this.Root.Pre.Pre.Next = this.Root
	this.Root.Pre = this.Root.Pre.Pre
	this.Len--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.Len == 0 {
		return -1
	}

	return this.Root.Next.Value
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.Len == 0 {
		return -1
	}
	return this.Root.Pre.Value

}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.Len == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.Len == this.Cap
}

func main() {
	circularDeque := Constructor(3)      // 设置容量大小为3
	circularDeque.InsertFront(3)         // 返回 true
	circularDeque.InsertLast(9)          // 返回 true
	fmt.Println(circularDeque.GetRear()) // 返回 9
}
