package main

import (
	"fmt"
	"time"

	"example.com/investment/input"
)

type user struct { // defining struct (JS object) type
	first     string
	last      string
	age       int
	createdAt time.Time
}

func main() {

	first := input.SInput()
	// last := input.SInput()
	last := "Tyagi"
	age := 21

	var person = user{ // basically objects
		first:     first,
		last:      last,
		age:       age,
		createdAt: time.Now(),
	}
	person.last = "yadav" // updating value of struct field

	structPointer(&person)

	structArgs(person)

}

func structArgs(person user) {

	fmt.Println("person =>", person.first)
	fmt.Println("person =>", person.last)
	fmt.Println("person =>", person.age)
	fmt.Println("person =>", person.createdAt)
}
func structPointer(person *user) {

	// this function takes a pointer to the user struct
	// and update its key
	person.age = 43

}
