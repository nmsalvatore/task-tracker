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
		want := "• drink coffee\n"

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
		want := "• do a little dance\n• make a little love\n"

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
	})

	t.Run("multiple tasks", func(t *testing.T) {
		tasks := Tasks{}
		items := []string{"do a little dance", "make a little love", "get down tonight"}
		tasks.Add(items...)

		got := tasks.Get()

		if len(got) != 3 {
			t.Errorf("got %d tasks, want 3", len(got))
		}
	})

	t.Run("id", func(t *testing.T) {
		tasks := Tasks{}
		item := "fly to the moon"
		tasks.Add(item)

		got := tasks.Get()

		if got[0].ID != 1 {
			t.Errorf("got %d, want %d", got[0].ID, 1)
		}
	})

	t.Run("description", func(t *testing.T) {
		tasks := Tasks{}
		item := "drink coffee"
		tasks.Add(item)

		got := tasks.Get()

		if got[0].Description != item {
			t.Errorf("got %q, want %q", got[0].Description, item)
		}
	})

	t.Run("status", func(t *testing.T) {
		tasks := Tasks{}
		item := "go on a walk"
		tasks.Add(item)

		got := tasks.Get()

		if got[0].Status != "todo" {
			t.Errorf("got %q, want %q", got[0].Status, "todo")
		}
	})
}
