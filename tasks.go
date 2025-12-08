package main

import (
	"errors"
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
		item := fmt.Sprintf("%d: %s\n", task.ID, task.Description)
		message.WriteString(item)
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
		now := time.Now()
		task := Task{
			ID:          len(t.items) + 1,
			Description: item,
			Status:      status,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		t.items = append(t.items, task)
	}

	return nil
}

func (t *Tasks) Mark(id int, status string) error {
	if status == "" {
		return errors.New("mark status empty")
	}

	options := []string{"todo", "in-progress", "done"}
	if !slices.Contains(options, status) {
		return fmt.Errorf("invalid status: %q (must be todo, in-progress, or done)", status)
	}

	for i := range t.items {
		if t.items[i].ID == id {
			t.items[i].Status = status
			t.items[i].UpdatedAt = time.Now()
			return nil
		}
	}

	// test this
	return errors.New("task not found")
}

func (t *Tasks) Delete(ids ...int) error {
	startLength := len(t.items)

	for _, id := range ids {
		for i := range t.items {
			if t.items[i].ID == id {
				t.items = slices.Delete(t.items, i, i+1)
				break
			}
		}
	}

	if len(t.items) == startLength {
		return errors.New("task not found")
	}

	return nil
}
