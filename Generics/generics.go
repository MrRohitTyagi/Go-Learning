package generics

func main() {
	intVar := addArgs(1, 2)
	floatVar := addArgs(1.1, 2.2)
	stringVar := addArgs("hello", "world")

}

// similar to Typescript's generics
// it allows us to write functions that can accept different types of arguments
// in this case, we are using a type constraint to specify that the type must be int, string or float64
// the type constraint is specified using a type parameter T
// the function will accept two arguments of type T and return a value of type T
// this is a simple example of generics in Go
func addArgs[T int | string | float64](a, b T) T {
	return a + b
}
