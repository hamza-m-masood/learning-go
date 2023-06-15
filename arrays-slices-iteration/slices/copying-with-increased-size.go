package slices

import "fmt"

func Copy() {
	s1 := []string{"abra", "kadabra"}
	fmt.Println(s1)
	var s2 []string
	s2 = append(s2, "hello")
	fmt.Println(s2)
}
