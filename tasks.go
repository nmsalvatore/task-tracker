package main

import (
	"fmt"
	"io"
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

func (t Tasks) Get() []Task {
	return t.items
}

func (t Tasks) List(writer io.Writer) {
	tasks := t.Get()
	if len(tasks) == 0 {
		fmt.Fprintln(writer, "no tasks")
		return
	}
	fmt.Fprintln(writer, "so many tasks")
}

func (t *Tasks) Add(items ...string) {
	for _, item := range items {
		task := Task{
			Description: item,
		}
		t.items = append(t.items, task)
	}
}
