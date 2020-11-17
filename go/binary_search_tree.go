package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Val         int
	Left, Right *Node
}

func (l *Node) insertNode(val int) error {

	newNode := &Node{Val: val}

	if l == nil {
		return errors.New("tree is nil")
	}

	if l.Val == val {
		return errors.New("The Tree already has the value")
	}

	if val < l.Val {
		if l.Left == nil {
			l.Left = newNode
			return nil
		}
		l.Left.insertNode(val)

	} else if val > l.Val {
		if l.Right == nil {
			l.Right = newNode
			return nil
		}
		l.Left.insertNode(val)
	}
	// fmt.Println(l)
	return nil
}

func (l *Node) displayAllNodes() {
	// pass
}

func (l *Node) fetchMin() int {
	if l.Left == nil {
		return l.Val
	}
	return l.Left.fetchMin()

}

func (l *Node) fetchMax() int {
	if l.Right == nil {
		return l.Val
	}
	return l.Right.fetchMax()
}

func main() {
	fmt.Println("Welcome to to the Binary search Tree")
	rootNode := &Node{Val: 55}
	rootNode.insertNode(13)
	rootNode.insertNode(12)
	rootNode.insertNode(10)
	rootNode.insertNode(18)
	rootNode.insertNode(20)
	rootNode.insertNode(21)
	rootNode.insertNode(6)

	minMumVal := rootNode.fetchMin()
	maxMumVal := rootNode.fetchMax()

	fmt.Println("Minimum value is BST is ", minMumVal)
	fmt.Println("Maximum value in BST is ", maxMumVal)
}
