package main

import (
	"errors"
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

	// cleaner way to create a struct
	var person2, _ = newPerson("rohit", "tyagi", 21)

	fmt.Println("person2 =>", person2)

	var person3, err = newPerson("abc", "xyz", 52)

	fmt.Println("person3 =>", person3)
	fmt.Println("err =>", err)

}

func newPerson(fName string, Lname string, age int) (*user, error) {

	// can also add some valiations here
	if age > 50 {
		return nil, errors.New("age is too high")
	}

	return &user{
		first: fName,
		last:  Lname,
		age:   age,
	}, nil
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
