package main

import (
	"bytes"
	"testing"
)

func TestListTasks(t *testing.T) {
	t.Run("no tasks message", func(t *testing.T) {
		tasks := Tasks{}
		buffer := bytes.Buffer{}
		tasks.List(&buffer)

		got := buffer.String()
		want := "no tasks\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestAddTasks(t *testing.T) {
	t.Run("single task", func(t *testing.T) {
		task := "drink coffee"

		tasks := Tasks{}
		tasks.Add(task)

		got := tasks.Get()

		if len(got) != 1 {
			t.Errorf("got %d task, want 1", len(got))
		}

		if got[0].Description != task {
			t.Errorf("got %s, want %s", got[0].Description, task)
		}
	})
}
