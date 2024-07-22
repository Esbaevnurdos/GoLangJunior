package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/main/note"
	"example.com/main/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content, err := getNoteData()
    

	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

		if err != nil {
		fmt.Println(err)
		return
	}


	userNote, err := note.New(title, content)



    err = outputData(userNote)

		if err != nil {
		return
	}

	err = outputData(todo)

		if err != nil {
		return
	}

}

func outputData(data outputtable) error {
    data.Display()
	return saveData(data)
}

func saveData(data saver) error {
      err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed!")
		return err
	}

	fmt.Println("\n","Saving the note succeeded!")
	return nil
}

func getNoteData() (string, string, error) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content, nil
}
 
func getUserInput(prompt string) (string) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

     return text
	
}


// for anything inputs types
func printSomething(value interface{}) { 
// switch value.(type) {
// case int:
// }
// value.Name
}