package main

import "fmt"

func main() {

	const value = 1

	switch value {

	case 1:
		{
			fmt.Print("case 1")
		}
	case 2:
		fmt.Print("case 2")
	case 3:
		fmt.Print("case 3")
	default:
		fmt.Print("default case")

	}
}
