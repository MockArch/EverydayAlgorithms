// Prority Queue using the linked List
// A typical priority queue supports following operations.
// insert(item, priority): Inserts an item with given priority.
// getHighestPriority(): Returns the highest priority item.
// deleteHighestPriority(): Removes the highest priority item.
package main

import (
	"fmt"
)

type Node struct {
	Data     int
	Priority int
	Next     *Node
}

type linkedList struct {
	head *Node
}

func (l *linkedList) Insert(data, priority int) {

	newNode := &Node{Data: data, Priority: priority}
	if l.head == nil {
		fmt.Println("New node is EMpty")
		l.head = newNode
	} else {

		headNode := l.head
		for headNode.Next != nil {
			headNode = headNode.Next
		}
		headNode.Next = newNode
	}
}

func (l *linkedList) isEmpty() bool {
	return l.head == nil
}
func (l *linkedList) Display() {
	headNode := l.head
	fmt.Println(headNode)
	for headNode.Next != nil {
		fmt.Println(headNode.Data, headNode.Priority)
		headNode = headNode.Next

	}
}

func (l *linkedList) Pop() int {
	var highPriority int
	var priorityData int
	_ = highPriority
	_ = priorityData
	if l.isEmpty() {
		return -1
	} else {

		tmpPtr := l.head
		highPriority = -1
		for tmpPtr.Next != nil {

			if highPriority > tmpPtr.Priority {
				highPriority = tmpPtr.Priority
				priorityData = tmpPtr.Data
			} else if highPriority == -1 {
				highPriority = tmpPtr.Priority
				priorityData = tmpPtr.Data
			}

			tmpPtr = tmpPtr.Next
		}
		return priorityData
	}
}
func main() {
	fmt.Println("Welcome to Proirity Queue.")
	linkList := &linkedList{}
	fmt.Println("Is Priority Queue is Empty ", linkList.isEmpty())
	linkList.Insert(12, 10)
	linkList.Insert(13, 1)
	linkList.Insert(15, 5)
	linkList.Insert(18, 7)
	fmt.Println("Is Priority Queue is Empty ", linkList.isEmpty())
	linkList.Display()

	i := true

	for i {
		popData := linkList.Pop()
		fmt.Println("Is Priority Queue is Poping Out ", popData)

	}

}
