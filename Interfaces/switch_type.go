package main

func main() {

	var value any = "thi is a string"
	// value can be of any type
	// we can use type switch to check the type of value
	// type switch is used to determine the type of an interface value
	// it is similar to a regular switch statement, but it checks the type of the value instead of its value
	// it is useful when we want to perform different actions based on the type of the value

	actualValue, ok := value.(string)
	println(actualValue, ok) // type assertion

	switch value.(type) {

	// this only // works with the type of value, not the value itself
	// and when the type of value is any, interface{}

	case int:
		{
			println("value is int")
		}
	case float64:
		{
			println("value is float64")
		}
	case string:
		{
			println("value is string")
		}
	}

}
