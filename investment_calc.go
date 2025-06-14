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

	var final = amount * math.Pow(1+interest/100, years)

	fmt.Println(final)

}
