// Adapter is a structural design pattern that
// allows objects with incompatible interfaces to collaborate.
package main

import "fmt"

type IProcess interface {
	process()
}

type Adaptee struct {
	adaptee int
}

type Adapter struct {
	adaptee Adaptee
}

func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert method")
}

func (adapter Adapter) process() {
	fmt.Println("Adapter Process")
	adapter.adaptee.convert()
}

func main() {
	var processor IProcess = Adapter{}
	processor.process()
}
