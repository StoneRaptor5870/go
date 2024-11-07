package main

import (
	"bufio"
	"fmt"
	"interfaces/note"
	"interfaces/todo"
	"os"
	"strings"
)

type saver interface { // convention to add -er (saver) if only one method is there
	Save() error
}

// type displayer interface {
// 	Display()
// }

// type outputtable interface {
// 	Save() error
// 	Display()
// }

type outputtable interface { // embedded interface
	saver
	Display()
}

func main() {
	printSomething("Notes & Todo App.")

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(userNote)
}

func printSomething(value interface{}) { // interface {} is for any value
	switch value.(type) { // type switches
	case int:
		fmt.Println("Interger:", value)
	case string:
		fmt.Println(value)
	}

	/* typedVal, ok := value.(int)         type info from values
	   if ok {
	       fmt.Println("Interger", typedVal)
		   return
	   }
	*/
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("saving the note failed")
		return err
	}

	fmt.Println("saving the note succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
