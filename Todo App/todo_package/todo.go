package todo_package

import (
	"errors"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string) error {
	if err := todos.ValidateTitle(title); err != nil {
		return err
	}

	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)

	return nil

}

func (todos Todos) ValidateIndex(index int) error {
	if index < 0 || index >= len(todos) {
		err := errors.New("error: invalid task index parsed")
		fmt.Printf("%v\n", err)
		return err
	}

	return nil
}

func (todos Todos) ValidateTitle(title string) error {
	t := todos
	if title == "" {
		err := errors.New("error: empty text cannot be a task title")
		fmt.Printf("%v\n", err)
		return err
	}
	for _, todo := range t {
		if todo.Title == title && !todo.Completed {
			err := errors.New("error: task already exist and not completed")
			fmt.Printf("%v\n", err)
			return err
		}
	}

	return nil
}

func (todos *Todos) Toggle(index int) error {
	if err := todos.ValidateIndex(index); err != nil {
		return err
	}
	t := *todos

	isCompleted := t[index].Completed
	if !isCompleted {
		now := time.Now()
		t[index].CompletedAt = &now
	}

	t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) Edit(index int, title string) error {
	if err := todos.ValidateIndex(index); err != nil {
		return err
	}
	if err := todos.ValidateTitle(title); err != nil {
		return err
	}
	t := *todos
	t[index].Title = title

	return nil
}

func (todos *Todos) Delete(index int) error {
	if err := todos.ValidateIndex(index); err != nil {
		return err
	}
	t := *todos
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) Print() error {
	myTable := table.NewWriter()
	myTable.SetOutputMirror(os.Stdout)
	myTable.AppendHeader(table.Row{"#", "Task", "Completed", "Created At", "Completed At"})

	t := *todos

	for index, todo := range t {
		completed := "❌"
		timeCompleted := ""
		if todo.Completed {
			completed = "✅"
			now := time.Now().Format(time.RFC1123)
			timeCompleted = now
		}
		myTable.AppendRows([]table.Row{
			{index, todo.Title, completed, todo.CreatedAt.Format(time.RFC1123), timeCompleted},
		})
		myTable.AppendSeparator()
	}
	myTable.SetStyle(table.StyleLight)
	myTable.Render()

	return nil
}
