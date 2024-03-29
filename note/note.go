package note

import (
	"errors"
	"fmt"
	"time"
	"os"
	"strings"
	"encoding/json"
)

// using back tick to assign metadata to struct also called struct tags, in this case so that json object created with it has the right keys
type Note struct {
	Title string `json:"title"`
	Content string	`json:"content"`
	CreatedAt time.Time	`json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("You note titled %v has the following content: \n\n%v\n\n", note.Title, note.Content)
}

func (note Note) Save() error{
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)
	
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644) // returning error or nil generated by WriteFile
}

func New(title string, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}