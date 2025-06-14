package main

import (
	"fmt"
	"os"
)

func main() {
	var inp string
	fmt.Println("Select Option")
	fmt.Println("1. Read File")
	fmt.Println("2. Write File")

	fmt.Scan(&inp)

	switch inp {

	case "1":
		read()
	case "2":

		fmt.Println("Enter the text: ")
		var data string
		fmt.Scan(&data)
		write(data)

	}

}

func read() {

	data, _ := os.ReadFile("temp.txt")
	dataInString := string(data) // Convert byte slice to string
	fmt.Print(dataInString)
}

func write(data string) {

	var byteData = []byte(data) // Convert string to byte slice

	os.WriteFile("temp.txt", byteData, 0644)

}
