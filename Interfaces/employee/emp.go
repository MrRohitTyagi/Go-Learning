package employee

import "fmt"

type Emp struct {
	name string
	age  int
}

func New(name string, age int) Emp {

	return Emp{
		name: name,
		age:  age,
	}
}

func (e Emp) Print() {

	fmt.Println("per =>", e.name)

}
