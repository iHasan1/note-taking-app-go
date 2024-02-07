package main

import "fmt"
import "errors"

func main() {
	title, content, err := getNoteData()

	if(err != nil) {
		fmt.Println(err)
		return
	}

	
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note Title:")

	if err != nil {
		return "", "", err
	}

	content, err := getUserInput("Note Content:")
	
	if err != nil {
		return "", "", err
	}

	return title, content, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value) // scanln also allows to register a enter key event

	if value == "" {
		return "", errors.New("Invalid input")
	}

	return value, nil
}