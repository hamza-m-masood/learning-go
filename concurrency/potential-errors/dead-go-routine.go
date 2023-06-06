package main

import "fmt"

func main() {
	fmt.Println("dead")
	go func() {
	}()
}

// empty go routine will hang around until process exits
// will not be garbage collected
