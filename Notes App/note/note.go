package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Label   string    `json:"label"`
	Content string    `json:"content"`
	Date    time.Time `json:"date_value"`
	// the  strings are knows as struct metadata
	// these can be anything, and will be picked by the package that is using this metadata
	// like the json package will liik into the metadata and will change the keys to the one in metadata
	// cool stuff
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
