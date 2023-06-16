package slices

import "fmt"

func CopyArray() {
	var d [4]string
	d[0] = "hello"
	d[1] = "why"
	d[2] = "thank"
	d[3] = "you"
	fmt.Println(d)
	s := make([]string, 0)
	s = d[:2]
	fmt.Println(s)
	fmt.Println(len(s), cap(s))
	// this will error out
	// fmt.Println(s[3])
	s = append(s, "test")
	fmt.Println(s)
	fmt.Println(len(s), cap(s))
}
