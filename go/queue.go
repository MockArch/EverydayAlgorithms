// A queue is a useful data structure in programming.
// It is similar to the ticket queue outside a cinema hall,
// where the first person entering the queue is the first person who gets the ticket.

// Queue follows the First In First Out(FIFO) rule -
// the item that goes in first is the item that comes out first too.

//   Enqueue --- >> [||, ||, |] --> Dequeue

// methods

// Enqueue: Add an element to the end of the queue
// Dequeue: Remove an element from the front of the queue
// IsEmpty: Check if the queue is empty
// IsFull: Check if the queue is full
// Peek: Get the value of the front of the queue without removing it
// Display : get printed all the value in Queue
package main

import "fmt"

type myQueue []int

func (q *myQueue) Enqueue(val int) {
	*q = append(*q, val)
}

func (q *myQueue) Dequeue() int {

	if len(*q) <= 0 {
		return -1
	} else {
		r := (*q)[0]
		*q = (*q)[1:]
		return r
	}
}

func (q *myQueue) Display() {
	i := 0
	for i = 0; i != -1; {
		s := (*q).Dequeue()
		if s != -1 {
			fmt.Println(s)
		} else {
			i = -1
		}
	}
}

func (q *myQueue) IsEmpty() bool {

	if len(*q) <= 0 {
		return true
	} else {
		return false
	}

}
func main() {
	var q = &myQueue{}

	q.Enqueue(12)
	q.Enqueue(13)
	q.Enqueue(14)
	fmt.Println("Is Queue is empty", q.IsEmpty())
	q.Enqueue(15)
	q.Enqueue(16)
	q.Enqueue(17)
	q.Display()
	fmt.Println("Is Queue is empty", q.IsEmpty())
	// q.Enqueue(12)
	// fmt.Println(q.([]int))
}
