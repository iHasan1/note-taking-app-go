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

type displayer interface {
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Displayer()
// }

// instead of above create an embedded interface
type outputtable interface {
	saver
	displayer
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

	err = outputData(todo)
	if err != nil {
		return
	}

	outputData(userNote) // no need to handle err as program will end here anyways
}

// can also use "any" as argument in place of interface{}
func printSomething(value interface{}) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("float:", floatVal)
		return
	}

	strVal, ok := value.(string)
	if ok {
		fmt.Println(strVal)
		return
	}
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("float:", value)
	// case string:
	// 	fmt.Println(value)
	// default:
	// 	//...
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
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