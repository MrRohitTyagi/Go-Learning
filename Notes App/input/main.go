package input

import "fmt"

func SInput() string {
	var variable string
	fmt.Scan(&variable)
	return variable
}
func NInput() int {
	var variable int
	fmt.Scan(&variable)
	return variable
}
