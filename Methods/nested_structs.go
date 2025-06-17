package main

import (
	admin "example.com/investment/Methods/Admin"
	employee "example.com/investment/Methods/Employee"
)

func main() {

	emp, _ := employee.New("Ashok", 100000)
	// emp.Print()

	admin1 := admin.New("a@a.com", "qwerty", &emp)
	admin1.Print()
}
