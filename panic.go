package main

func main() {

	for i := 0; i < 1000; i++ {
		if i == 100 {
			panic("This will stop the program")
		}
	}
}
