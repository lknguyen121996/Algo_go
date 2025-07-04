package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func toList(head *Node) []int {
	var result []int
	for head != nil {
		result = append(result, head.data)
		head = head.next
	}
	return result
}

func main() {
	head := &Node{data: 1}
	head.next = &Node{data: 2}
	head.next.next = &Node{data: 3}
	head.next.next.next = &Node{data: 4}
	head.next.next.next.next = &Node{data: 5}
	head.next.next.next.next.next = &Node{data: 6}
	head.next.next.next.next.next.next = &Node{data: 7}
	head.next.next.next.next.next.next.next = &Node{data: 8}
	head.next.next.next.next.next.next.next.next = &Node{data: 9}
	head.next.next.next.next.next.next.next.next.next = &Node{data: 10}
	fmt.Println(toList(head))
	fmt.Println(toList(reverseLinkList2(head)))
}

// bad solution
func reverseLinkList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	var (
		origin = head
		left   = head
		tmp    = head
		right  = head.next
	)
	for right.next != nil {
		left = right
		right = right.next
		left.next = tmp
		tmp = left
	}
	right.next = left
	origin.next = nil
	return right
}

// good solution use recursion
func reverseLinkList2(head *Node) *Node {
	if head == nil{
		return nil // now head is the tail
	}
	newHead := head
	//  |
	//  1 -> 2 -> 3 -> 4 -> 5 -> nil
	if head.next != nil{
		newHead = reverseLinkList2(head.next)
		// newhead = 5 -> nil
		//                 |
		//  1 -> 2 -> 3 -> 4 -> 5 -> nil
		head.next.next = head // reverse the link
		// 4 -> 5 -> nil => 4 -> 5 -> 4
	}
	// 5 -> 4 -> nil
	head.next = nil // set the tail's next to nil
	return newHead // 4
}
