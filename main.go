package main

import "fmt"
import "example.com/note-taking-app/note"

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

}

func getNoteData() (string, string) {
	title := getUserInput("Note Title:")

	content := getUserInput("Note Content:")

	return title, content
}

func getUserInput(prompt string) (string) {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value) // scanln also allows to register a enter key event

	return value
}