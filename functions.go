package main

import "fmt"

func main() {

	var output = add(3, 5)

	fmt.Print(output)

}

func add(a int, b int) string {
	// Sprint is used to format the string and return it
	return fmt.Sprint("The addition is: ", a+b)
}
