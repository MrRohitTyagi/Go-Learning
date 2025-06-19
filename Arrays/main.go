package main

import "fmt"

type arr [3]string

func main() {

	// Arrays are fixed-size sequences of elements of the same type.
	//  var prices = [length of array]type{v1,v2,v3,...}
	var prices = [5]int{100, 200, 300, 400, 500}
	var fruits = arr{"apple", "banana", "cherry"}
	fmt.Println("prices =>", prices)
	fmt.Println("fruits =>", fruits)

	prices[4] = 40000 // change the value of the third element
	fmt.Println("prices =>", prices)

	// Slices

	fmt.Println("index 1 to 3(excluded) =>", prices[1:3])
	fmt.Println("index 1 to last  =>", prices[1:])
	fmt.Println(" index 0 to 3rd(excluded) =>", prices[:3])

	arrayLoops(fruits)
}

func arrayLoops(fruits arr) {
	for i := 0; i < len(fruits); i++ {

		f := fruits[i]

		println(f)

	}
}
