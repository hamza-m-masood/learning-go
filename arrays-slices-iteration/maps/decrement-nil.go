package maps

import "fmt"

func Decrement() {
	m1 := map[string]int{"first": 0}
	m1["second"]--
	fmt.Println(m1["second"])
}
