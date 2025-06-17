package admin

import (
	"fmt"

	employee "example.com/investment/Methods/Employee"
)

type Admin struct {
	// if we want to expose these struct keys, it must start with capital letter
	password string
	email    string
	employee *employee.Emp
}

func New(email string, password string, emp *employee.Emp) Admin {

	return Admin{
		password: password,
		email:    email,
		employee: emp,
	}
}
func (admin Admin) Print() {
	fmt.Println(admin.email + " " + admin.password)
	printEmp(*admin.employee)
}
func printEmp(emp employee.Emp) {
	emp.Print()
}
