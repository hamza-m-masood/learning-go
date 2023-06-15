package maps

import "fmt"

func Append() {
	m1 := make(map[string]string)
	m1["mark"] = "john"
	test := map[string]string{"hello": "hello"}
	keys := make([]string, 0, len(m1))
	for k := range m1 {
		keys = append(keys, k)
	}
	fmt.Println(test)
	fmt.Println(keys)

}
