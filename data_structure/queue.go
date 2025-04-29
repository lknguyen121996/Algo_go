package main

import "fmt"

func main() {
	queue := []int{1, 2, 3, 4, 5}
	queue = append(queue, 6)
	fmt.Println(queue)
}

type Queue struct {
	curr LinkedNode
	next LinkedNode
}

type LinkedNode struct {
	Val  int
	Next *LinkedNode
}

func (q *Queue) enqueue(val int) {
	q.next.Val = val
	q.next = q.next.Next
}

func (q *Queue) dequeue() int {
	val := q.curr.Val
	q.curr = q.curr.Next
	return val
}
