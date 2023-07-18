package main

type List{
	elements interface{}
	size int
}

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

// TODO: instantiate new list

// TODO: append value(s) at the end of list

// TODO: return the element at specific index

// TODO: return all elements in list

// TODO: remove element at given index

//rTODO: eturn index of given value in list

// TODO: (internal) grow the list by the growthfactor

// TODO: (internal) srhink the list by the shrinkfactor
