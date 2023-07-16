package main

import "fmt"

type element struct {
	value interface{}
	next  *element
}

type List struct {
	FirstElement *element
	LastElement  *element
	Size         int
}

// new (initilize the linked list)
func New(values ...interface{}) *List {
	list := &List{}
	for _, value := range values {
		list.Append(value)
	}
	return list
}

// append element
func (list *List) Append(values ...interface{}) {
	for _, v := range values {
		newElement := &element{value: v}
		if list.Size == 0 {
			list.FirstElement = newElement
			list.LastElement = newElement
		} else {
			list.LastElement.next = newElement
			list.LastElement = newElement
		}
		list.Size++
	}
}

// prepend elment
func (list *List) Prepend(values ...interface{}) {
	for _, v := range values {
		newElement := &element{value: v}
		if list.Size == 0 {
			list.FirstElement = newElement
			list.LastElement = newElement
		} else {
			newElement.next = list.FirstElement
			list.FirstElement = newElement
		}
		list.Size++
	}
}

//TODO: remove elment

//TODO: empty (returns if list does not contain enaything)

//TODO: clear (removes all elements from the list)

func main() {
	list := New(2, 3, 3)
	list.Prepend(5, 6, 8)
	fmt.Println(list.FirstElement.value)

}
