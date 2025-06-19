package main

import "fmt"

func main() {
	printSomething(1)
	printSomething(1.1)
	printSomething("Hello, World!")
}

// it will accept any type of value
// the type is known as empty interface
func printSomething(value any) { // or you can use interface{} instead of any
	fmt.Println(value)
}
