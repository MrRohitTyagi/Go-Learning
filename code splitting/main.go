package main

import (
	"fmt"

	"example.com/splitting/utils" // for custom packages the folder name should be same as the package name
	// and the folder should be in the same directory as the main.go file
	// "example.com/splitting/utils" // for custom packages the folder name should be same as the package name
)

func main() {

	re := utils.Sum(1, 2)

	fmt.Println("return", re)
}
