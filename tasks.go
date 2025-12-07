package main

import (
	"fmt"
	"io"
	"strings"
	"time"
)

type Task struct {
	ID          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Tasks struct {
	items []Task
}

func (t *Tasks) Get() []Task {
	return t.items
}

func (t *Tasks) List(writer io.Writer) {
	tasks := t.Get()
	if len(tasks) == 0 {
		fmt.Fprintln(writer, "no tasks")
		return
	}

	message := &strings.Builder{}
	for _, task := range tasks {
		item := fmt.Sprintf("todo:\t%s\n", task.Description)
		message.WriteString(item)
	}

	fmt.Fprint(writer, message)
}

func (t *Tasks) Add(items ...string) {
	for _, item := range items {
		task := Task{
			Description: item,
		}
		t.items = append(t.items, task)
	}
}
