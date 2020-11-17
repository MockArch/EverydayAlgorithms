package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type linkedList struct {
	head *Node
	len  int64
}

func (l *linkedList) insert(v int) {
	newNode := &Node{value: v}
	if l.head == nil && l.len == 0 {
		l.head = newNode
		l.len += 1
	} else {
		ptr := l.head
		for ptr.next != nil {
			ptr = ptr.next
		}
		ptr.next = newNode
	}
}

func (l *linkedList) printNodes() {

	if l.head == nil {
		fmt.Println("Node is empyt")
	} else {
		ptr := l.head
		for ptr.next != nil {
			fmt.Println(ptr.value)
			ptr = ptr.next
		}
	}
}

func (l *linkedList) insertAt(position, v int) {

	if position >= 0 {
		newNode := &Node{value: v}
		tmpPtr := l.head
		posCounter := 0
		for tmpPtr != nil {
			if posCounter == position {
				nextNode := tmpPtr.next
				tmpPtr.next = newNode
				newNode.next = nextNode
				break
			}
			tmpPtr = tmpPtr.next
			posCounter += 1
		}
	}
}

func main() {
	fmt.Println("Welcome to single single linked List")
	newLinkedList := linkedList{}
	newLinkedList.insert(12)
	newLinkedList.insert(14)
	newLinkedList.insert(33)
	newLinkedList.insert(43)
	newLinkedList.insert(54)
	newLinkedList.insert(16)
	newLinkedList.insert(17)
	newLinkedList.insert(73)
	newLinkedList.insert(78)
	newLinkedList.insert(100)

	newLinkedList.insertAt(6, 1000)
	newLinkedList.printNodes()

}
