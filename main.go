package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.com/note-taking-app/note"
	"example.com/note-taking-app/todo"
)

// No need to explicitly link todo and note structs as Go does it automatically by verifying the signature of the method present in the structs therefore saveData fn works
type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo Text:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = saveData(userNote)
	if err != nil {
		return
	}

	userNote.Display()
	err = saveData(userNote)
	if err != nil {
		return
	}
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}
	fmt.Println("Saving the note succeeded!")
	return nil
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