package main

import "fmt"

func main() {

	a := 10
	add10(&a)

	fmt.Println("a =>", a)
}

func add10(a *int) { // this  syntam means, this  function will receive a poiter
	*a += 10 // using * will get the value og the pointer referign to
	// this will change the value of a directly throufht reference
}
