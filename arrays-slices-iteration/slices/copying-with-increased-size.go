package slices

import "fmt"

func Copy() {
	s1 := []string{"abra", "kadabra", "kazaam", "wow"}
	s2 := make([]string, 1, 9)
	copy(s2, s1)
	fmt.Println(s2)
}
