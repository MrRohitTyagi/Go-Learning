package main

import (
	"fmt"

	"example.com/splitting/constants/apikeys"
	"example.com/splitting/utils" // for custom packages the folder name should be same as the package name

	// and the folder should be in the same directory as the main.go file
	// "example.com/splitting/utils" // for custom packages the folder name should be same as the package name
	"example.com/splitting/constants"
)

func main() {

	key := apikeys.MongoKey

	re := utils.Sum(constants.A, constants.B)

	fmt.Println("return", key, re)
}
