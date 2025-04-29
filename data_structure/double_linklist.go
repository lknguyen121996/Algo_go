package main

import "fmt"

type MyLinkedList struct {
	Head *MyLinkedListNode
	Tail *MyLinkedListNode
}

type MyLinkedListNode struct {
	Value int
	Next  *MyLinkedListNode
	Prev  *MyLinkedListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	count := 0
	for node := this.Head; node != nil; node = node.Next {
		if count == index {
			return node.Value
		}
		count++
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	if this.Head == nil && this.Tail == nil {
		initNode := &MyLinkedListNode{val, nil, nil}
		this.Head = initNode
		this.Tail = initNode
		return
	}
	newNode := &MyLinkedListNode{Value: val, Prev: nil, Next: this.Head}
	this.Head.Prev = newNode
	this.Head = newNode
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.Head == nil && this.Tail == nil {
		initNode := &MyLinkedListNode{val, nil, nil}
		this.Head = initNode
		this.Tail = initNode
		return
	}
	newNode := &MyLinkedListNode{Value: val, Prev: this.Tail, Next: nil}
	this.Tail.Next = newNode
	this.Tail = newNode
}

// cover outOfIndex
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	newNode := &MyLinkedListNode{Value: val, Prev: nil, Next: nil}
	count := 0
	for node := this.Head; node != nil; node = node.Next {
		// move to left
		if count == index {
			newNode.Prev = node.Prev
			node.Prev.Next = newNode
			newNode.Next = node
			node.Prev = newNode
			return
		}
		count++
	}
	return
}

// cover outOfIndex
func (this *MyLinkedList) DeleteAtIndex(index int) {
	count := 0
	for node := this.Head; node != nil; node = node.Next {
		// move to left
		if count == index {
			// head tail 0
			if this.Head.Next == nil && this.Tail.Prev == nil {
				this.Head = nil
				this.Tail = nil
				return
			}
			if count == 0 {
				this.Head.Next.Prev = nil
				this.Head = this.Head.Next
			}
			if node.Next == nil {
				this.Tail.Prev.Next = nil
				this.Tail = this.Tail.Prev
			}
			node.Prev.Next = node.Next
			node.Next.Prev = node.Prev
			node = nil
			return
		}
		count++
	}
}

func (this *MyLinkedList) PrintList() []int {
	array := make([]int, 0)
	for node := this.Head; node != nil; node = node.Next {
		array = append(array, node.Value)
	}
	return array
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
func main() {
	obj := Constructor()

	//node := obj.Get(0)
	//fmt.Println(node)
	//
	//obj.AddAtHead(1)
	//obj.AddAtHead(0)
	//
	//obj.AddAtTail(2)
	//obj.AddAtTail(3)
	//fmt.Println(obj.PrintList())
	//
	//obj.AddAtIndex(1, 4)
	//fmt.Println(obj.PrintList())
	//
	//obj.DeleteAtIndex(1)
	//fmt.Println(obj.PrintList())
	obj.AddAtHead(1)
	obj.DeleteAtIndex(0)
	fmt.Println(obj)

}
