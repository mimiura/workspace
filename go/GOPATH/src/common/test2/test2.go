package test2

import "fmt"

type Person struct {
	name string
}

func NewPerson(name string) *Person {
	p := new(Person)
	p.name = name
	return p
}

func (p *Person) greet(msg string) {
	fmt.Printf("%s, I'm %s.\n", msg, p.name)
}

func (p *Person) Greet(msg string) {
	p.greet(msg)
}

func (p *Person) SetName(aaa string) {
	p.name = aaa
}
