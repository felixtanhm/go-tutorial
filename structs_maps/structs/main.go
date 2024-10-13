package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
}

var jim = person{
	firstName: "Jim",
	lastName:  "Tan",
	contact: contactInfo{
		email: "jim@gmail.com",
	},
}

func main() {
	jim.updateName("Felix")
	jim.print()
}

func (personPointer *person) updateName(newName string) {
	(*personPointer).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
