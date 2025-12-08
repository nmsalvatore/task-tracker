package main

import (
	"bytes"
	"slices"
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
		tasks.Add("", "drink coffee")

		buffer := bytes.Buffer{}
		tasks.List(&buffer)

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
		tasks.List(&buffer)

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
		tasks.List(&buffer)

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
		tasks.List(&buffer)

		got := buffer.String()
		want := "× 1: drink coffee\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("single task", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "drink coffee")

		got := tasks.Get()

		if len(got) != 1 {
			t.Errorf("got %d tasks, want 1", len(got))
		}
	})

	t.Run("multiple tasks", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "do a little dance", "make a little love", "get down tonight")

		got := tasks.Get()

		if len(got) != 3 {
			t.Errorf("got %d tasks, want 3", len(got))
		}
	})

	t.Run("id", func(t *testing.T) {
		tasks := Tasks{}
		description := "fly to the moon"
		tasks.Add("", description)

		got := tasks.Get()

		if got[0].ID != 1 {
			t.Errorf("got %d, want %d", got[0].ID, 1)
		}
	})

	t.Run("id after delete", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "one", "two", "three", "four")
		tasks.Delete(2)
		tasks.Add("", "five")

		got := tasks.Get()

		if got[3].ID != 5 {
			t.Errorf("got %d, want %d", got[3].ID, 5)
		}
	})

	t.Run("description", func(t *testing.T) {
		tasks := Tasks{}
		description := "drink coffee"
		tasks.Add("", description)

		got := tasks.Get()

		if got[0].Description != description {
			t.Errorf("got %q, want %q", got[0].Description, description)
		}
	})

	t.Run("status default", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "go on a walk")

		got := tasks.Get()

		if got[0].Status != "todo" {
			t.Errorf("got %q, want %q", got[0].Status, "todo")
		}
	})

	t.Run("status specified", func(t *testing.T) {
		tasks := Tasks{}
		status := "in-progress"
		tasks.Add(status, "go on a walk")

		got := tasks.Get()

		if got[0].Status != status {
			t.Errorf("got %q, want %q", got[0].Status, "in-progress")
		}
	})

	t.Run("status rejected", func(t *testing.T) {
		tasks := Tasks{}
		err := tasks.Add("dude", "go on a walk")
		if err == nil {
			t.Error("wanted error but didn't get one")
		}
	})

	t.Run("created at", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first")
		tasks.Add("", "second")

		got := tasks.Get()

		if !got[0].CreatedAt.Before(got[1].CreatedAt) {
			t.Errorf("first not created before second")
		}
	})
}

func TestMark(t *testing.T) {
	t.Run("status update", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first", "second")
		tasks.Add("done", "third")

		tasks.Mark(1, "in-progress")
		tasks.Mark(2, "done")
		tasks.Mark(3, "todo")

		got := tasks.Get()

		for i, status := range []string{"in-progress", "done", "todo"} {
			if got[i].Status != status {
				t.Errorf("got %q, want %q", got[i].Status, status)
			}
		}
	})

	t.Run("empty status", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first")

		err := tasks.Mark(1, "")
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})

	t.Run("invalid status", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first")

		err := tasks.Mark(1, "gettin' it")
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})

	t.Run("update time", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first")
		tasks.Mark(1, "in-progress")

		got := tasks.Get()

		if got[0].UpdatedAt.Equal(got[0].CreatedAt) {
			t.Error("time not updated")
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "sip tea")

		err := tasks.Mark(2, "in-progress")
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}

	})
}

func TestDelete(t *testing.T) {
	t.Run("single task", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first", "second", "third")
		tasks.Delete(2)

		got := tasks.Get()

		for _, task := range got {
			if task.Description == "second" {
				t.Error("task not deleted")
			}
		}

		if len(got) == 3 {
			t.Error("list length unchanged")
		}
	})

	t.Run("multiple tasks", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first", "second", "third")
		tasks.Delete(2, 3)

		got := tasks.Get()

		deleted := []string{"second", "third"}
		for _, task := range got {
			if slices.Contains(deleted, task.Description) {
				t.Errorf("%q not deleted", task.Description)
			}
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		tasks := Tasks{}
		err := tasks.Delete(1)
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("single task", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first", "third")
		tasks.Update(2, "second")

		got := tasks.Get()

		if got[1].Description != "second" {
			t.Errorf("got %q, want %q", got, "second")
		}
	})
}
