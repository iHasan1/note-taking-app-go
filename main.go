package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.com/note-taking-app/note"
	"example.com/note-taking-app/todo"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return
	}

	fmt.Println("Saving the note succeeded!")
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title:")

	content := getUserInput("Note Content:")

	return title, content
}

func getUserInput(prompt string) (string) {
	fmt.Printf("%v ",prompt)

	/* scanning like this wont work with longer inputs
	// var value string
	// fmt.Scanln(&value) // scanln also allows to register a enter key event
	*/

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')	// we need single quotes here, not double to specify single characters

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")	// we need to remove the line break from string because readString does not remove it
	text = strings.TrimSuffix(text, "\r")	// in windows sometimes line break is defined by \r

	return text
}