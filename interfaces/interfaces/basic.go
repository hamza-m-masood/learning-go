package interfaces

import "fmt"

type sentient interface {
	hello()
}
type Person struct {
	name string
}

type Animal struct {
	name string
}

func (p Person) hello() {
	fmt.Println("hello, my name is", p.name)
}

func (a Animal) hello() {
	fmt.Println("hello, my name is", a.name)
}

func sayHello(s sentient) {
	s.hello()
}

func RunBasic() {
	a := Animal{"walrus"}
	p := Person{
		name: "hermes",
	}
	sayHello(a)
	sayHello(p)

}
