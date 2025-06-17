package main

import "fmt"

type user struct {
	name string
	age  int
}

// adding (u ,structye)before the function name  will make the function attact to that struct
func (u user) print(extra string) {
	fmt.Println("u.name =>", u.name)
	fmt.Println("u.age =>", u.age)
	fmt.Println("extra =>", extra)
}

// This will mutate the struct key,
// Important, pass the pointer
func (u *user) mutateName(newName string) {
	u.name = newName

}

func main() {
	var person = user{

		name: "rohit",
		age:  21,
	}

	person.print("extra stuff") // call the attached struct
	person.mutateName("ABC")
	person.print("extra stuff") // call the attached struct

	// data mutation

}
