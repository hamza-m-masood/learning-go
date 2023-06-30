package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	m := make(map[string]int64)
	fmt.Println(SumNumbers[string, int64](m))
}
