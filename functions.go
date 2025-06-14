package main

import "fmt"

func main() {

	var output = add(3, 5)

	fmt.Print(output)

}

func add(a int, b int) int {
	return a + b
}
