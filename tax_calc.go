package main

import (
	"fmt"
)

func main() {

	var salary float64
	const tax = 12.5

	fmt.Print("Enter your annual salary: ")
	fmt.Scan(&salary)
	fmt.Println("Tax Calculator")
	var taxAmount = salary * tax / 100
	fmt.Printf("Your tax amount is: %.2f\n", taxAmount)
	var netSalary = salary - taxAmount
	fmt.Printf("Your net salary after tax is: %.2f\n", netSalary)

}
