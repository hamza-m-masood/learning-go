package tuples

import "fmt"

// A tuple is a finite sorted list of elements.
func PowerSeries(a int) (int, int) {
	return a * a, a * a * a
}

func Tuple() {
	var square int
	var cube int
	square, cube = PowerSeries(3)
	fmt.Println("Square", square, "Cube", cube)
}
