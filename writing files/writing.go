package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	for {

		var inp string
		fmt.Println("Select Option")
		fmt.Println("1. Read File")
		fmt.Println("2. Write File")

		fmt.Scan(&inp)

		switch inp {

		case "1":
			read("defaultValue")
		case "2":

			fmt.Println("Enter the text: ")
			var data string
			fmt.Scan(&data)
			write(data)

		}

	}
}

func read(defaultValue string) (string, error) {

	data, err := os.ReadFile("temp.txt")

	//Error handeling

	if err != nil {

		fmt.Println("Has error")
		return defaultValue, errors.New(err.Error())
	}

	fmt.Println("err", err)
	dataInString := string(data) // Convert byte slice to string
	fmt.Print(dataInString)

	return dataInString, nil
}

func write(data string) {

	var byteData = []byte(data) // Convert string to byte slice

	os.WriteFile("temp.txt", byteData, 0644)

}
