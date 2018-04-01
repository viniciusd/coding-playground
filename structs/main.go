package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	foo := person{
		firstName: "Foo",
		lastName:  "Bar",
		contactInfo: contactInfo{
			email:   "foo@bar.com",
			zipCode: 42000,
		},
	}

	foo.updateName("Hey")
	foo.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
