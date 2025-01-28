package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string 
	Content   string
	CreatedAt time.Time `json:"created_at"`
}

// Method to display user input
func (note Note) Display() {
	fmt.Printf("Your note titled '%v' has the following content:\n\n%v\n\n", note.Title, note.Content)
}

// Method to save the note to a file
func (note Note) Save() error {
	
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	data, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}

// Constructor function for Note
func New(title string, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title or content cannot be empty")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
