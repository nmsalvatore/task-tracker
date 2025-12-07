package main

import (
	"bytes"
	"testing"
)

func TestList(t *testing.T) {
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

	t.Run("single task", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("drink coffee")

		buffer := bytes.Buffer{}
		tasks.List(&buffer)

		got := buffer.String()
		want := "todo:\tdrink coffee\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("multiple tasks", func(t *testing.T) {
		tasks := Tasks{}
		items := []string{"do a little dance", "make a little love"}
		tasks.Add(items...)

		buffer := bytes.Buffer{}
		tasks.List(&buffer)

		got := buffer.String()
		want := "todo:\tdo a little dance\ntodo:\tmake a little love\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("single task", func(t *testing.T) {
		tasks := Tasks{}
		item := "drink coffee"
		tasks.Add(item)

		got := tasks.Get()

		if len(got) != 1 {
			t.Errorf("got %d tasks, want 1", len(got))
		}

		if got[0].Description != item {
			t.Errorf("got %q, want %q", got[0].Description, item)
		}
	})

	t.Run("multiple tasks", func(t *testing.T) {
		tasks := Tasks{}
		items := []string{"do a little dance", "make a little love"}
		tasks.Add(items...)

		got := tasks.Get()

		if len(got) != 2 {
			t.Errorf("got %d tasks, want 2", len(got))
		}

		for i := range 2 {
			if got[i].Description != items[i] {
				t.Errorf("got %q, want %q", got[i].Description, items[i])
			}
		}
	})
}
