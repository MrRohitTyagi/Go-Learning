package main

import (
	"fmt"

	"example.com/investment/Interfaces/employee"
)

type Person struct {
	name string
}

// Interfaces helps us to define a contract that a type must implement
// helps us in writing more generic code
// basically make ssure 2 structs have a mehtod common to both of them
// here we have a common interface that both Person and Employee implement
// so we can pass both of them to a function that accepts this interface

type common interface {
	Print()
}

func (per Person) Print() {
	fmt.Println("per =>", per.name)
}

func doSomething(data common) {
	data.Print()
}

func main() {

	var person = Person{
		name: "John Doe",
	}
	var employee = employee.New("Jane Doe", 30)

	doSomething(person)
	doSomething(employee)
}
