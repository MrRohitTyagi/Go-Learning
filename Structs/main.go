package main

import (
	"fmt"
	"time"

	"example.com/investment/input"
)

type user struct {
	first     string
	last      string
	age       int
	createdAt time.Time
}

func main() {

	first := input.SInput()
	last := input.SInput()
	age := input.NInput()

	var person = user{

		first:     first,
		last:      last,
		age:       age,
		createdAt: time.Now(),
	}
	structArgs(person)
}

func structArgs(person user) {

	fmt.Println("person =>", person.first)
	fmt.Println("person =>", person.last)
	fmt.Println("person =>", person.age)
	fmt.Println("person =>", person.createdAt)
}
