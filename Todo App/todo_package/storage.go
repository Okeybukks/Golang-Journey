package todo_package

import (
	"encoding/json"
	"fmt"
	"os"
)

type Storage[T any] struct {
	Filename string
}

func NewStorage[T any](filename string) Storage[T] {
	return Storage[T]{Filename: filename}
}

func (s Storage[T]) Save(data *T) {
	fileData, err := json.MarshalIndent(*data, "", "\t")
	if err != nil {
		fmt.Printf("%v", err)
	}

	if err := os.WriteFile(s.Filename, fileData, 0644); err != nil {
		fmt.Printf("%v", err)
	}
}

func (s Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.Filename)
	if err != nil {
		return err
	}

	if len(fileData) == 0 {
		return nil
	}

	if err := json.Unmarshal(fileData, data); err != nil {
		fmt.Printf("%v", err)
	}

	return nil
}
