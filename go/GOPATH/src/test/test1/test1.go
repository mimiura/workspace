package test1

import "fmt"

type Person struct {
	Name string
}

func (p Person) greet(msg string) {
	fmt.Printf("%s, I'm %s.\n", msg, p.Name)
}

func (p Person) Greet(msg string) {
	p.greet(msg)
}
