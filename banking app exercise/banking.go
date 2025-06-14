package main

import "fmt"

func main() {

	var total float64 = 100000.0

	// for i := 0; i < 2; i++ { // loop will run
	// 	banking(total)
	// }

	fmt.Println("\n\n\nWelcome to my banking app\n\n\n")

	for { // infinite loop
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
		fmt.Print("Your current balance is: ", total)
	} else if choice == 2 {
		pf("Enter amount to deposit: ")
		var deposit float64
		fmt.Scan(&deposit)
		total = total + deposit
		pf("updated balance: ", total)
	} else if choice == 3 {

		pf("Enter amount to withdraw")

		var amount float64
		fmt.Scan(&amount)

		total = total - amount

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
