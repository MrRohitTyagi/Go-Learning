package utils

func Sum(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}
func Mul(a, b int) int {
	return a * b
}
func Div(a, b int) int {
	if b == 0 {
		return 0 // Avoid division by zero
	}
	return a / b
}
