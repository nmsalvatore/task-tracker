package main

import (
	"bytes"
	"io"
	"testing"
)

const filename = "test.json"

func TestCLI_Add(t *testing.T) {
	t.Run("single task", func(t *testing.T) {
		cli := NewCLI(filename)

		args := []string{"one"}
		cli.Add(io.Discard, args)

		got := cli.tasks.Get()
		if len(got) != 1 {
			t.Fatalf("got %d tasks, want 1", len(got))
		}

		if got[0].Description != args[0] {
			t.Errorf("got %q, want %q", got[0].Description, args)
		}
	})

	t.Run("multiple tasks", func(t *testing.T) {
		cli := NewCLI(filename)

		args := []string{"one", "two", "three"}
		err := cli.Add(io.Discard, args)
		if err != nil {
			t.Fatal(err)
		}

		got := cli.tasks.Get()
		if len(got) != 3 {
			t.Fatalf("got %d tasks, want 3", len(got))
		}

		for i := range got {
			if got[i].Description != args[i] {
				t.Errorf("got %q, want %q", got[i].Description, args[i])
			}
		}
	})

	t.Run("single task message", func(t *testing.T) {
		cli := NewCLI(filename)

		buf := bytes.Buffer{}
		err := cli.Add(&buf, []string{"one"})
		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := "Added task \"one\"\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("multiple tasks message", func(t *testing.T) {
		cli := NewCLI(filename)

		buf := bytes.Buffer{}
		err := cli.Add(&buf, []string{"one", "two"})
		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := "Added task \"one\"\nAdded task \"two\"\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestCLI_Clear(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		cli := NewCLI(filename)

		err := cli.tasks.Add("", "one", "two", "three")
		if err != nil {
			t.Fatal(err)
		}

		err = cli.Clear(io.Discard, []string{})
		if err != nil {
			t.Fatal(err)
		}

		got := cli.tasks.Get()
		if len(got) != 0 {
			t.Errorf("got %d tasks, want 0", len(got))
		}
	})

	t.Run("by status", func(t *testing.T) {
		cli := NewCLI(filename)

		err := cli.tasks.Add("", "one", "two", "three", "four")
		if err != nil {
			t.Fatal(err)
		}

		err = cli.tasks.Mark("done", 1, 2)
		if err != nil {
			t.Fatal(err)
		}

		err = cli.Clear(io.Discard, []string{"done"})
		if err != nil {
			t.Fatal(err)
		}

		got := cli.tasks.Get()
		if len(got) != 2 {
			t.Errorf("got %d tasks, want 2", len(got))
		}
	})

	t.Run("all message", func(t *testing.T) {
		cli := NewCLI(filename)

		err := cli.tasks.Add("", "one", "two")
		if err != nil {
			t.Fatal(err)
		}

		buf := bytes.Buffer{}
		err = cli.Clear(&buf, []string{})
		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := "Cleared all tasks\n"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("by status message", func(t *testing.T) {
		cli := NewCLI(filename)

		err := cli.tasks.Add("", "one", "two", "three", "four")
		if err != nil {
			t.Fatal(err)
		}

		err = cli.tasks.Mark("done", 1, 2)
		if err != nil {
			t.Fatal(err)
		}

		buf := bytes.Buffer{}
		err = cli.Clear(&buf, []string{"done"})
		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := "Cleared all tasks with status \"done\"\n"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("invalid status", func(t *testing.T) {
		cli := NewCLI(filename)

		err := cli.tasks.Add("zero", "one", "two", "three", "four")
		if err == nil {
			t.Error("want error, but didn't get one")
		}
	})
}
