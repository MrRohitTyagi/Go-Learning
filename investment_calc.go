package main

import (
	"fmt"
	"math"
)

func main() {

	var amount float64 = 1000
	interest := 5.5
	var years float64 = 10

	fmt.Println("investment calculator")

	var value = amount * math.Pow(1+interest/100, years)
	const inflationRate = 2.0

	var valueAfterInflation = value - (value / 100 * inflationRate * years)

	fmt.Println(valueAfterInflation)

}
