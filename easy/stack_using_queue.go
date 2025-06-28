/*
Implement a last-in-first-out (LIFO) stack using only two queues. The implemented stack should support all the functions of a normal stack (push, top, pop, and empty).

Implement the MyStack class:

void push(int x) Pushes element x to the top of the stack.
int pop() Removes the element on the top of the stack and returns it.
int top() Returns the element on the top of the stack.
boolean empty() Returns true if the stack is empty, false otherwise.
Notes:

You must use only standard operations of a queue, which means that only push to back, peek/pop from front, size and is empty operations are valid.
Depending on your language, the queue may not be supported natively. You may simulate a queue using a list or deque (double-ended queue) as long as you use only a queue's standard operations
*/
package main

type MyStackNode struct {
	Val  int
	Next *MyStackNode
}

func (node *MyStackNode) Push(x int) {
	node.Next = &MyStackNode{
		Val: x,
	}
}

type MyStack struct {
	OriginQueue  *MyStackNode
	ReverseQueue *MyStackNode
	Size         int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	// init
	originNode := &MyStackNode{
		Val: x,
	}
	if this.Size == 0 {
		this.OriginQueue = originNode
		this.ReverseQueue = originNode
		return
	}

	this.Size++
}

func (this *MyStack) Pop() int {
	if this.Size != 0 {
		this.Size--
	}
	return 0
}

func (this *MyStack) Top() int {

}

func (this *MyStack) Empty() bool {
	if this.Size == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {

}
