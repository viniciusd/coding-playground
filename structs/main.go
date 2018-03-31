package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	foo := person{firstName: "Foo", lastName: "Bar"}
	fmt.Println(foo)
}
