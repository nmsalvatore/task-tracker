package main

import (
	"fmt"
	"io"
	"time"
)

type Task struct {
	id          int
	description string
	status      string
	createdAt   time.Time
	updatedAt   time.Time
}

type Tasks struct {
	items []Task
}

func (t Tasks) List(writer io.Writer) {
	tasks := t.Get()

	if len(tasks) == 0 {
		fmt.Fprintln(writer, "no tasks")
		return
	}

	fmt.Fprintln(writer, "so many tasks")
}

func (t Tasks) Get() []Task {
	return t.items
}
