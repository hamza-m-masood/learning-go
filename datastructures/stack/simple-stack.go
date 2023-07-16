package main

import (
	"fmt"
)

// simple stack creation: LIFO
type stack []string

var s stack

// push
func (s *stack) push(v string) {
	*s = append(*s, v)
}

// pop
func (s *stack) pop() {
	//check if slice is empty
	if len(*s) == 0 {
		fmt.Println("empty slice")
	} else {
		//get the top most index from the slice
		i := len(*s) - 1
		//remove the top most element using index
		*s = (*s)[:i]
	}
}
func main() {
	s = make([]string, 0)
	s.push("hello")
	s.push("why")
	s.pop()
	fmt.Println(s)

}
