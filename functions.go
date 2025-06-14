package main

import "fmt"

func main() {

	var output = add(3, 5)
	var sum, sub, multi, divi = withMultipleReturns(3, 5)

	fmt.Printf(output, sum, sub, multi, divi)

}

func add(a int, b int) string {
	// Sprint is used to format the string and return it
	return fmt.Sprint("The addition is: ", a+b)
}

func withMultipleReturns(a int, b int) (int, int, int, int) {

	sum := a + b
	sub := a - b
	mult := a * b
	divi := a / b
	return sum, sub, mult, divi
}
