package main

import "fmt"

type Person struct {
	Name string
}

// Person 型に対してメソッドを定義する
func (p Person) Greet(msg string) {
	fmt.Printf("%s, I'm %s.\n", msg, p.Name)
}

func main() {
	a := Person{Name: "Taro"} // 値型の変数を用意する
	a.Greet("Hi")             //=> Hi, I'm Taro.
}
