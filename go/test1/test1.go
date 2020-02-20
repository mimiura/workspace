package test1

import "fmt"

type Person struct {
	Name string
}

// Person 型に対してメソッドを定義する
func (p Person) Greet(msg string) {
	fmt.Printf("%s, I'm %s.\n", msg, p.Name)
}
