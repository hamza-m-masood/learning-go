package generics

import "fmt"

func RunBasic() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}
	fmt.Printf("ints: %v\nfloats: %v", SumIntsOrFloats[string, int64](ints), SumIntsOrFloats[string, float64](floats))
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
