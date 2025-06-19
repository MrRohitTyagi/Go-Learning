package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	var personArray = [5]person{
		person{name: "rohit", age: 21},
		person{name: "mohit", age: 22},
		person{name: "sohit", age: 25},
		person{name: "fohit", age: 11},
		person{name: "lohit", age: 51},
	}
	fmt.Println("personArray =>", personArray)

	fmt.Println(" =>", len(personArray), cap(personArray))

}
