package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Text string `json:"title"` // struct tag
}

func (t Todo) Display() {
	fmt.Printf("Your todo has the following content:\n\n%v\n", t.Text)
}

func (t Todo) Save() error {
	fileName := strings.ReplaceAll(t.Text, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(t)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(text string) (*Todo, error) {
	if text == "" {
		return &Todo{}, errors.New("invalid input")
	}

	return &Todo{
		Text: text,
	}, nil
}
