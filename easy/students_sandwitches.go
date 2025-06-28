// Number of Students Unable to Eat Lunch
// The school cafeteria offers circular and square sandwiches at lunch break, referred to by numbers 0 and 1 respectively. All students stand in a queue. Each student either prefers square or circular sandwiches.

// The number of sandwiches in the cafeteria is equal to the number of students. The sandwiches are placed in a stack. At each step:

// If the student at the front of the queue prefers the sandwich on the top of the stack, they will take it and leave the queue.
// Otherwise, they will leave it and go to the queue's end.
// This continues until none of the queue students want to take the top sandwich and are thus unable to eat.

// You are given two integer arrays students and sandwiches where sandwiches[i] is the type of the i​​​​​​th sandwich in the stack (i = 0 is the top of the stack) and students[j] is the preference of the j​​​​​​th student in the initial queue (j = 0 is the front of the queue). Return the number of students that are unable to eat.
package main

import "fmt"

func main() {
	fmt.Println(countStudents([]int{1, 1, 0, 0}, []int{0, 1, 0, 1}))
	fmt.Println(countStudents([]int{1,1,1,0,0,1}, []int{1,0,0,0,1,1}))
}

type Queue struct {
	TopNode     *NodeQueue
	CurrentNode *NodeQueue
}

type NodeQueue struct {
	Value int
	Next  *NodeQueue
}

func (q *Queue) Push(val int) {
	if q.TopNode == nil {
		q.TopNode = &NodeQueue{val, nil}
		q.CurrentNode = q.TopNode
	} else {
		q.CurrentNode.Next = &NodeQueue{val, nil}
		q.CurrentNode = q.CurrentNode.Next
	}
}

func (q *Queue) Dequeue() int {
	val := q.TopNode.Value
	q.TopNode = q.TopNode.Next
	return val
}

func (q *Queue) Peek() int {
	return q.TopNode.Value
}

func (q *Queue) Size() int {
	size := 0
	currentNode := q.TopNode
	for currentNode != nil {
		size++
		currentNode = currentNode.Next
	}
	return size
}

func ArrayToQueue(arr []int) *Queue {
	queue := &Queue{}
	for _, val := range arr {
		queue.Push(val)
	}
	return queue
}

func countStudents(students []int, sandwiches []int) int {
	studentsQueue := ArrayToQueue(students)
	sandwitchesQueue := ArrayToQueue(sandwiches)

	count := 0

	for count < len(students) {
		if studentsQueue.TopNode == nil || sandwitchesQueue.TopNode == nil {
			break
		}
		if studentsQueue.Peek() == sandwitchesQueue.Peek() {
			studentsQueue.Dequeue()
			sandwitchesQueue.Dequeue()
			count = 0
		} else {
			studentsQueue.Push(studentsQueue.Dequeue())
			count++
		}
	}
	return studentsQueue.Size()
}
