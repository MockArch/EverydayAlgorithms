package main

type Node struct {
	value int
	next  *Node
}

type linkedList struct {
	head *Node
	len  int
}

func (l *linkedList) addNode(val int) {
	newNode := &Node{value: val}

	if l.len == 0 {
		l.head = newNode
		l.len++
		return
	}
	tmpPtr := l.head
	for i := 0; i < l.len; i++ {
		if tmpPtr.next == nil {
			tmpPtr.next = newNode
			l.len++
			return

		}
		tmpPtr = tmpPtr.next
	}

}
