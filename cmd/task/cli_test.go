package main

import (
	"bytes"
	"io"
	"testing"
)

const filename = "test.json"

func TestCLI_Add(t *testing.T) {
	t.Run("single task", func(t *testing.T) {
		want := []string{"one"}

		cli := NewCLI(filename)
		cli.Add(io.Discard, want)

		got := cli.tasks.Get()
		if len(got) != 1 {
			t.Fatalf("got %d tasks, wanted 1", len(got))
		}

		if got[0].Description != want[0] {
			t.Errorf("got %q, wanted %q", got[0].Description, want)
		}
	})

	t.Run("multiple tasks", func(t *testing.T) {
		want := []string{"one", "two", "three"}

		cli := NewCLI(filename)
		cli.Add(io.Discard, want)

		got := cli.tasks.Get()
		if len(got) != 3 {
			t.Fatalf("got %d tasks, wanted 3", len(got))
		}

		for i := range got {
			if want[i] != got[i].Description {
				t.Errorf("got %q, wanted %q", got[i].Description, want[i])
			}
		}
	})

	t.Run("single task message", func(t *testing.T) {
		cli := NewCLI(filename)
		buf := bytes.Buffer{}
		cli.Add(&buf, []string{"one"})

		got := buf.String()
		want := "Added task \"one\"\n"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("multiple tasks message", func(t *testing.T) {
		cli := NewCLI(filename)
		buf := bytes.Buffer{}
		cli.Add(&buf, []string{"one", "two"})

		got := buf.String()
		want := "Added task \"one\"\nAdded task \"two\"\n"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}
