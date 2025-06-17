package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Label   string
	Content string
	Date    time.Time
}

func New(label, content string) Note {

	return Note{
		Label:   label,
		Content: content,
		Date:    time.Now(),
	}
}
func (note Note) Print() {
	fmt.Println(note.Label + " " + note.Content)
}
func (note Note) Save() {
	// converting a struct into a json
	// only keys with capital Letters will be added into json
	json, _ := json.Marshal(note)

	fileName := strings.TrimSuffix(note.Label, "\n")
	fileName = strings.TrimSuffix(note.Label, "\r")
	fileName = strings.ReplaceAll(fileName, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	os.WriteFile(fileName, json, 0644)
}
