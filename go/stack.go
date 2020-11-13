package main

import "fmt"

// Stack is a linear data structure which follows a particular order in which the operations are performed.
// The order may be LIFO(Last In First Out) or FILO(First In Last Out).

type stackDataType struct {
	stack []int
}

func (s *stackDataType) push(val int) {
	s.stack = append(s.stack, val)
}

func (s *stackDataType) pop() int {
	popOfIndex := len(*&s.stack) - 1
	if popOfIndex <= 0 {
		return -1
	} else {
		r := s.stack[popOfIndex]
		s.stack = (*&s.stack)[:popOfIndex]
		return r
	}
}

func main() {
	s := stackDataType{}
	s.push(12)
	s.push(13)
	s.push(14)
	fmt.Println(s.pop())
	fmt.Println(s.pop())
}
