package employee

import (
	"errors"
	"fmt"
)

type Emp struct {
	name   string
	salary float64
}

func (emp Emp) Print() {
	message := fmt.Sprintln(emp.name + " earns " + fmt.Sprint(emp.salary))
	fmt.Print(message)
}

func New(name string, salary float64) (Emp, error) {

	isRich := isRich(salary)

	if isRich == true {
		return Emp{
			name:   name,
			salary: salary,
		}, nil

	} else {

		return Emp{}, errors.New(fmt.Sprintln(name + " IS POOR AF"))
	}
}

func isRich(salary float64) bool {

	if salary < 10000 {
		return false
	} else {
		return true
	}

}
