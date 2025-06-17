package main

import "fmt"

type str string // this will make  out own type str
// that is just a extension of string type
// advantage ? we can assign custom methods to it now

func (text str) log() {
	fmt.Println(text)
}

func main() {

	var name str = "Rohit"

	name.log() // attached custom method
}
