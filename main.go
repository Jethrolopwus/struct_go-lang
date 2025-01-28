package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	//  displaying the user input on the terminal
	userNote.Display()

	// Save the note and check for errors

	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed:", err)
		return
	}

	fmt.Println("Saving the note succeeded.")
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	text = strings.TrimSpace(text) 
	return text
}
