package main

import (
	"fmt"
	"io"
	"slices"
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
		switch task.Status {
		case "todo":
			message.WriteString("• ")
		case "in-progress":
			message.WriteString("> ")
		case "done":
			message.WriteString("× ")
		}
		message.WriteString(task.Description + "\n")
	}

	fmt.Fprint(writer, message)
}

func (t *Tasks) Add(status string, items ...string) error {
	if status == "" {
		status = "todo"
	}

	options := []string{"todo", "in-progress", "done"}
	if !slices.Contains(options, status) {
		return fmt.Errorf("invalid status: %q (must be todo, in-progress, or done)", status)
	}

	for _, item := range items {
		task := Task{
			ID:          len(t.items) + 1,
			Description: item,
			Status:      status,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		t.items = append(t.items, task)
	}

	return nil
}
