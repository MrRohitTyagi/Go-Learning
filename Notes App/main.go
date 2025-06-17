package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/notes/note"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// Read text from terminal with spaces
	fmt.Println("Enter a label")
	label, _ := reader.ReadString('\n')
	fmt.Println("Enter the content")

	content, _ := reader.ReadString('\n')

	if label == "" {
		return
	}

	note := note.New(label, content)
	note.Print()
	note.Save()

}
