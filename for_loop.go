package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("traditional LOOP", i)
	}

	for {
		fmt.Println("This is an infinite loop")
		break
	}

	isGraterThan10 := true

	var i = 0
	for isGraterThan10 {

		fmt.Println("This is a while loop type", i)
		i++
		if i > 10 {
			isGraterThan10 = false
		}
	}
}
