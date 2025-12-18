package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestDisplay_PrintTasks(t *testing.T) {
	t.Run("no tasks message", func(t *testing.T) {
		tasks := Tasks{}

		buffer := bytes.Buffer{}
		PrintTasks(&buffer, tasks.Get())

		got := buffer.String()
		want := "task list empty\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("single task", func(t *testing.T) {
		tasks := Tasks{}

		err := tasks.Add("", "drink coffee")
		if err != nil {
			t.Fatal(err)
		}

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

		err := tasks.Add("", descriptions...)
		if err != nil {
			t.Fatal(err)
		}

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

		err := tasks.Add("in-progress", "build task tracker")
		if err != nil {
			t.Fatal(err)
		}

		buffer := bytes.Buffer{}
		PrintTasks(&buffer, tasks.Get())

		got := buffer.String()
		want := fmt.Sprintf("> 1: %s\n", Bold("build task tracker"))

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("single task done", func(t *testing.T) {
		tasks := Tasks{}

		err := tasks.Add("done", "drink coffee")
		if err != nil {
			t.Fatal(err)
		}

		buffer := bytes.Buffer{}
		PrintTasks(&buffer, tasks.Get())

		got := buffer.String()
		want := fmt.Sprintf("× 1: %s\n", Strike("drink coffee"))

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("by status", func(t *testing.T) {
		tasks := Tasks{}

		err := tasks.Add("", "first", "second", "third")
		if err != nil {
			t.Fatal(err)
		}

		err = tasks.Mark("done", 2)
		if err != nil {
			t.Fatal(err)
		}

		buffer := bytes.Buffer{}
		items, _ := tasks.GetByStatus("done")
		PrintTasks(&buffer, items)

		got := buffer.String()
		want := fmt.Sprintf("× 2: %s\n", Strike("second"))

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
