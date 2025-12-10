package main

import (
	"bytes"
	"testing"
)

func TestPrintTasks(t *testing.T) {
	t.Run("no tasks message", func(t *testing.T) {
		tasks := Tasks{}
		buffer := bytes.Buffer{}

		PrintTasks(&buffer, tasks.Get())

		got := buffer.String()
		want := "no tasks\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("single task", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "drink coffee")
		buffer := bytes.Buffer{}

		PrintTasks(&buffer, tasks.Get())

		got := buffer.String()
		want := "• 1: drink coffee\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("multiple tasks", func(t *testing.T) {
		tasks := Tasks{}
		descriptions := []string{"do a little dance", "make a little love"}
		tasks.Add("", descriptions...)
		buffer := bytes.Buffer{}

		PrintTasks(&buffer, tasks.Get())

		got := buffer.String()
		want := "• 1: do a little dance\n• 2: make a little love\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("single task in progress", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("in-progress", "build task tracker")
		buffer := bytes.Buffer{}

		PrintTasks(&buffer, tasks.Get())

		got := buffer.String()
		want := "> 1: build task tracker\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("single task done", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("done", "drink coffee")
		buffer := bytes.Buffer{}

		PrintTasks(&buffer, tasks.Get())

		got := buffer.String()
		want := "× 1: drink coffee\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
