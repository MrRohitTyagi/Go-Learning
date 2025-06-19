package main

import "fmt"

func main() {

	// Arrays are fixed-size sequences of elements of the same type.
	//  var prices = [length of array]type{v1,v2,v3,...}
	var prices = [3]int{100, 200, 300}
	var fruits = [3]string{"apple", "banana", "cherry"}
	fmt.Println("prices =>", prices)
	fmt.Println("fruits =>", fruits)

}
