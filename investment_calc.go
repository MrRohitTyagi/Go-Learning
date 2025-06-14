package main

import (
	"fmt"
	"math"
)

func main() {

	var amount float64
	interest := 5.5
	var years float64
	const inflationRate = 2.0

	// get user input from terminal and store in variables using fmt.scan()
	fmt.Print("Enter the amount you want to invest: ")
	fmt.Scan(&amount) // & is used to get the address of the variable and will be used to store the input value
	fmt.Print("Enter the number of years you want to invest for: ")
	fmt.Scan(&years)

	fmt.Println("investment calculator")

	var value = amount * math.Pow(1+interest/100, years)

	var valueAfterInflation = value - (value / 100 * inflationRate * years)

	fmt.Println(valueAfterInflation)

}
