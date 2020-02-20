package main

import (
	"fmt"
	"test/test1"
	"test/test2"
)

func main() {
	fmt.Printf("Hello World\n")

	a := test1.Person{Name: "Taro"}
	a.Greet("Hi")

	b := new(test1.Person)
	b.Name = "Hanako"
	b.Greet("Hi")

	c := test2.NewPerson("Jiro")
	c.Greet("Hi")

	d := new(test2.Person)
	d.SetName("Saburo")
	d.Greet("Hi")
}
