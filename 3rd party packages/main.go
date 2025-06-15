package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func main() {
	res := randomdata.Address()

	fmt.Println("res", res)
}
