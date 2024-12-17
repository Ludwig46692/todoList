package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	title       string
	completed   bool
	createdAt   time.Time
	completedAt *time.Time
}

type todos []Todo

func (todos *todos) addTodo(title string) {
	newTodo := Todo{
		title:       title,
		completed:   false,
		createdAt:   time.Now(),
		completedAt: nil,
	}
	*todos = append(*todos, newTodo)
}

func (todos *todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (todos *todos) deleteTodo(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *todos) toggleTodo(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	isCompleted := t[index].completed
	if !isCompleted {
		completionTime := time.Now()
		t[index].completedAt = &completionTime
	}
	t[index].completed = !isCompleted
	return nil
}

func (todos *todos) editTodo(index int, title string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return nil
	}
	t[index].title = title
	return nil
}

func (todos *todos) printTodos() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.completed {
			completed = "✅"
			if t.completedAt != nil {
				completedAt = t.completedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(
			strconv.Itoa(index),
			t.title,
			completed,
			t.createdAt.Format(time.RFC1123),
			completedAt,
		)
	}
	table.Render()
}
