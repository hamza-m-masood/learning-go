package main

type List struct {
	elements []interface{}
	size     int
}

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

// TODO: instantiate new list
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Append(values...)
	}
	return list
}

// TODO: append value(s) at the end of list
func (list *List) Append(values ...interface{}) {
	list.growBy(len(values))
	for _, v := range values {
		list.elements = append(list.elements, v)
	}
}

// TODO: return the element at specific index

// TODO: return all elements in list

// TODO: remove element at given index

// TODO: eturn index of given value in list

// TODO: (internal) grow the list by the growthfactor
func (list *List) growBy(newElementsLength int) {
	cap := len(list.elements)
	if list.size+newElementsLength >= cap {
		newCapacity := int(growthFactor * float32(cap+newElementsLength))
	}
}

// TODO: (internal) srhink the list by the shrinkfactor
