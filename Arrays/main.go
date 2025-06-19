package main

import "fmt"

type arr [3]string

func main() {

	// Arrays are fixed-size sequences of elements of the same type.
	//  var prices = [length of array]type{v1,v2,v3,...}
	var prices = [3]int{100, 200, 300}
	var fruits = arr{"apple", "banana", "cherry"}
	fmt.Println("prices =>", prices)
	fmt.Println("fruits =>", fruits)

	arrayLoops(fruits)
}

func arrayLoops(fruits arr) {
	for i := 0; i < len(fruits); i++ {

		f := fruits[i]

		println(f)

	}
}
