package main

//  AUTHORS  Mockarach
import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type linkedList struct {
	head *Node
	tail *Node
}

func (l *linkedList) insertNode(val int) {
	newNode := &Node{
		data: val,
	}
	if l.head == nil {
		l.head = newNode
		l.head = newNode
	} else {

		currentNode := l.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		l.tail = newNode

	}

}

func (l *linkedList) displayNode() {

	headNode := l.head
	// sizeOf := l.sizeOfLinkedList()
	var formattedAllNodes string
	for i := 0; headNode != nil; i++ {
		// fmt.Println("value = %v and prev = %v and next= %v\n", headNode.data, headNode.prev, headNode.next)
		formattedAllNodes += fmt.Sprintf("| %d | <---> ", headNode.data)
		headNode = headNode.next
	}
	fmt.Println(formattedAllNodes)
}

func (l *linkedList) sizeOfLinkedList() int {
	counter := 0
	headNode := l.head

	for headNode != nil {
		headNode = headNode.next
		counter++
	}
	return counter

}

func main() {
	fmt.Println("Welcome to the doubly linked list in the Go.")
	mainLinkedList := linkedList{}
	mainLinkedList.insertNode(12)
	mainLinkedList.insertNode(3)
	mainLinkedList.insertNode(1)
	mainLinkedList.insertNode(52)
	mainLinkedList.insertNode(62)
	mainLinkedList.insertNode(92)
	mainLinkedList.insertNode(42)
	mainLinkedList.insertNode(82)
	mainLinkedList.displayNode()
}

// output
// Welcome to the doubly linked list in the Go.
// | 12 | <---> | 3 | <---> | 1 | <---> | 52 | <---> | 62 | <---> | 92 | <---> | 42 | <---> | 82 | <--->
