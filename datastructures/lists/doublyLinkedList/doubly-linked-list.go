package main

import "fmt"

type element struct {
	value interface{}
	next  *element
	prev  *element
}

type List struct {
	FirstElement *element
	LastElement  *element
	Size         int
}

// new
func New(values ...interface{}) *List {
	list := &List{}
	for _, value := range values {
		list.Append(value)
	}
	return list
}

// append
func (list *List) Append(values ...interface{}) {
	for _, v := range values {
		newElement := &element{value: v}
		if list.Size == 0 {
			list.FirstElement = newElement
			list.LastElement = newElement
		} else {
			list.LastElement.next = newElement
			list.LastElement.next.prev = list.LastElement
			list.LastElement = newElement
		}
		list.Size++
	}
}

// prepend
func (list *List) Prepend(values ...interface{}) {
	for _, v := range values {
		newElement := &element{value: v}
		if list.Size == 0 {
			list.FirstElement = newElement
			list.LastElement = newElement
		} else {
			newElement.next = list.FirstElement
			list.FirstElement.prev = newElement
			list.FirstElement = newElement
		}
		list.Size++
	}
}

// remove elment
func (list *List) Remove(id int) {
	element := list.FirstElement
	if id == 0 {
		list.FirstElement = element.next
		list.FirstElement.prev = nil
		list.Size--
	}
	count := 0
	for i := 0; i < list.Size; i++ {
		if count == id-1 {
			//element that is behind the element we need to do delete
			prevDelElement := element
			//element that is in front of the element we need to ddelete
			postDeleteElement := element.next.next
			//element we need to do delete
			delElement := element.next
			if delElement == list.LastElement {
				list.LastElement = prevDelElement
				list.LastElement.prev = prevDelElement.prev
				list.LastElement.next = nil
				list.Size--
				break
			}
			delElement = prevDelElement
			delElement.next = postDeleteElement
			postDeleteElement.prev = delElement
			list.Size--
			break
		}
		element = element.next
		count++
	}
}

// print all values of list
func (list *List) allValues() []interface{} {
	allValues := make([]interface{}, 0)
	element := list.FirstElement
	for i := 0; i < list.Size; i++ {
		allValues = append(allValues, element.value)
		element = element.next
	}
	return allValues
}

func main() {
	list := New(1, 4, 6, 8)
	list.Remove(2)
	fmt.Println(list.allValues()...)
}
