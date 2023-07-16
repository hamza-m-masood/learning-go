// simple queue implementation FIFO
package main

import "fmt"

type queue []int

// add to the start
func (q *queue) Enqueue(v int) {
	*q = append(queue{v}, *q...)
}

// delete from the start
func (q *queue) Dequeue() {
	*q = append((*q)[1:], (*q)[2:]...)
}
func main() {
	q := make(queue, 0)
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
}
