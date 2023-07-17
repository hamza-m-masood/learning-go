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

// TODO: remove elment
func (list *List) Remove(id int) {
	element := list.FirstElement
	if id == 0 {
		list.FirstElement = element.next
		list.Size--

	}
	count := 0
	for i := 0; i < list.Size; i++ {
		if count == id-1 {
			if element.next == list.LastElement {
				list.LastElement = element
			}
			element.next = element.next.next
			list.Size--
			break
		}
		element = element.next
		count++
	}
}

// TODO: print all values of list
func (list *List) allValues() []interface{} {
	allValues := make([]interface{}, 0)
	element := list.FirstElement
	for i := 0; i < list.Size; i++ {
		allValues = append(allValues, element.value)
		element = element.next
	}
	return allValues
}

//TODO: empty (returns if list does not contain enaything)

//TODO: clear (removes all elements from the list)

func main() {
	list := New(2, 3, 5)
	list.Remove(1)
	fmt.Println(list.allValues())
	fmt.Println(list.LastElement.value)
}
