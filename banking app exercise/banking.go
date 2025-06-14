package main

import (
	"fmt"
	"os"
	"strconv"
)

const db = "db.txt"

func write(number float64) {

	var data = []byte(fmt.Sprint(number))

	os.WriteFile(db, data, 0064)
}

func read() float64 {
	var data, _ = os.ReadFile(db)

	dataInString := string(data)
	balance, _ := strconv.ParseFloat(dataInString, 64)
	return balance

}

func main() {

	fmt.Println("\n\n\nWelcome to my banking app\n\n\n")
	for { // infinite loop
		total := read()
		banking(total)
	}
}

func banking(total float64) {

	fmt.Println("Please select an option:")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	// store the choice
	var choice int
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your current balance is: ", total)
	} else if choice == 2 {
		pf("Enter amount to deposit: ")
		var deposit float64
		fmt.Scan(&deposit)
		total = total + deposit
		write(total)
		pf("updated balance: ", total)
	} else if choice == 3 {

		pf("Enter amount to withdraw")

		var amount float64
		fmt.Scan(&amount)

		total = total - amount

		write(total)

		pf("updated balance:", total)
	} else if choice == 4 {
		pf("Thank you for using the banking app. Goodbye!")
		return
	} else {
		pf("Invalid choice. Please try again.")

	}

}

func pf(str string, args ...interface{}) {
	fmt.Println(str, args)
}
