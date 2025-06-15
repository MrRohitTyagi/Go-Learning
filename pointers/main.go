package main

import "fmt"

func main() {

	var a = 10
	var aPointer *int // * sign indicates that aPointer is a pointer to an int
	aPointer = &a     // aPointer now holds the address of a

	fmt.Println("\n aPointer =>", aPointer)   // this will print the address of the variables
	fmt.Println("\n *aPointer =>", *aPointer) // this will print the value of the variable it points to

	a = 23 // changing the value of a

	fmt.Println("a =>", a)
	fmt.Println("*aPointer =>", *aPointer) // this will still print the value of a, which is 23
}
