package main

import (
	"bytes"
	"fmt"
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
		want := "added task \"one\"\n"

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
		want := "added task \"one\"\nadded task \"two\"\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("single task with status flag before tasks", func(t *testing.T) {
		cli := NewCLI(filename)

		status := "in-progress"
		args := []string{"--status", status, "drink coffee"}

		err := cli.Add(io.Discard, args)
		if err != nil {
			t.Fatal(err)
		}

		got := cli.tasks.Get()
		if got[0].Status != status {
			t.Errorf("got %q, want %q", got[0].Status, status)
		}
	})

	t.Run("single task with status flag after tasks", func(t *testing.T) {
		cli := NewCLI(filename)

		status := "in-progress"
		args := []string{"drink coffee", "--status", status}

		err := cli.Add(io.Discard, args)
		if err != nil {
			t.Fatal(err)
		}

		got := cli.tasks.Get()
		if len(got) != 1 {
			t.Fatalf("got length %d, want 1", len(got))
		}

		if got[0].Status != status {
			t.Errorf("got %q, want %q", got[0].Status, status)
		}
	})

	t.Run("status flag with invalid value", func(t *testing.T) {
		cli := NewCLI(filename)

		status := "didit"
		args := []string{"drink coffee", "--status", status}

		err := cli.Add(io.Discard, args)
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})

	t.Run("status flag as last arg, no value", func(t *testing.T) {
		cli := NewCLI(filename)

		args := []string{"drink coffee", "--status"}

		err := cli.Add(io.Discard, args)
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})

	t.Run("invalid flag", func(t *testing.T) {
		cli := NewCLI(filename)

		args := []string{"drink coffee", "--guzzle", "yess"}

		err := cli.Add(io.Discard, args)
		if err == nil {
			t.Error("wanted error, but didn't get one")
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
		want := "cleared all tasks\n"
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
		want := "cleared all tasks marked done\n"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("invalid status", func(t *testing.T) {
		cli := NewCLI(filename)

		err := cli.tasks.Add("", "one", "two", "three", "four")
		if err != nil {
			t.Fatal(err)
		}

		err = cli.Clear(io.Discard, []string{"bleh"})
		if err == nil {
			t.Error("want error, but didn't get one")
		}
	})
}

func TestCLI_Delete(t *testing.T) {
	t.Run("single task", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one", "two", "three")

		err := cli.Delete(io.Discard, []string{"2"})
		if err != nil {
			t.Fatal(err)
		}

		got := cli.tasks.Get()
		want := []string{"one", "three"}

		if len(got) != 2 {
			t.Fatalf("got length %d, want 2", len(got))
		}

		for i := range got {
			if got[i].Description != want[i] {
				t.Errorf("got %q, want %q", got[i].Description, want[i])
			}
		}
	})

	t.Run("multiple tasks", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one", "two", "three", "four")

		err := cli.Delete(io.Discard, []string{"1", "2"})
		if err != nil {
			t.Fatal(err)
		}

		got := cli.tasks.Get()
		want := []string{"three", "four"}

		if len(got) != 2 {
			t.Fatalf("got length %d, want 2", len(got))
		}

		for i := range got {
			if got[i].Description != want[i] {
				t.Errorf("got %q, want %q", got[i].Description, want[i])
			}
		}
	})

	t.Run("no tasks", func(t *testing.T) {
		cli := NewCLI(filename)
		err := cli.Delete(io.Discard, nil)
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})

	t.Run("single task message", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one", "two")

		buf := bytes.Buffer{}
		err := cli.Delete(&buf, []string{"1"})
		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := "task 1 deleted\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("multiple tasks message", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one", "two", "three", "four")

		buf := bytes.Buffer{}
		err := cli.Delete(&buf, []string{"1", "3"})
		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := "task 1 deleted\ntask 3 deleted\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestCLI_List(t *testing.T) {
	t.Run("no tasks", func(t *testing.T) {
		cli := NewCLI(filename)

		buf := bytes.Buffer{}
		cli.List(&buf, nil)

		got := buf.String()
		want := "task list empty\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("some todos", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one", "two", "three")

		buf := bytes.Buffer{}
		cli.List(&buf, nil)

		got := buf.String()
		want := "• 1: one\n• 2: two\n• 3: three\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("by status, in-progress", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one", "two")
		cli.tasks.Add("in-progress", "three", "four")
		cli.tasks.Add("done", "five", "six")

		buf := bytes.Buffer{}
		cli.List(&buf, []string{"in-progress"})

		got := buf.String()
		want := fmt.Sprintf("> 3: %s\n> 4: %s\n", Bold("three"), Bold("four"))

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("by status, done", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one", "two")
		cli.tasks.Add("in-progress", "three", "four")
		cli.tasks.Add("done", "five", "six")

		buf := bytes.Buffer{}
		cli.List(&buf, []string{"done"})

		got := buf.String()
		want := fmt.Sprintf("× 5: %s\n× 6: %s\n", Strike("five"), Strike("six"))

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("invalid status", func(t *testing.T) {
		cli := NewCLI(filename)
		err := cli.List(io.Discard, []string{"complete"})
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})

	t.Run("too many arguments", func(t *testing.T) {
		cli := NewCLI(filename)
		err := cli.List(io.Discard, []string{"done", "extra"})
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})
}

func TestCLI_Mark(t *testing.T) {
	t.Run("single task", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one")

		status := "in-progress"
		err := cli.Mark(io.Discard, []string{"1", status})
		if err != nil {
			t.Fatal(err)
		}

		got := cli.tasks.Get()
		if got[0].Status != status {
			t.Errorf("got status %q, want %q", got[0].Status, status)
		}
	})

	t.Run("multiple tasks", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one", "two", "three")

		status := "done"
		err := cli.Mark(io.Discard, []string{"1", "2", "3", status})
		if err != nil {
			t.Fatal(err)
		}

		got := cli.tasks.Get()

		for i := range got {
			if got[i].Status != status {
				t.Errorf("got status %q, want %q", got[i].Status, status)
			}
		}
	})

	t.Run("no arguments", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one")

		err := cli.Mark(io.Discard, []string{})
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})

	t.Run("no status", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one")

		err := cli.Mark(io.Discard, []string{"1"})
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})

	t.Run("no id", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one")

		err := cli.Mark(io.Discard, []string{"done"})
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})

	t.Run("invalid status", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one")

		err := cli.Mark(io.Discard, []string{"1", "doing"})
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one")

		err := cli.Mark(io.Discard, []string{"2", "done"})
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})

	t.Run("valid and invalid ids", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one")

		err := cli.Mark(io.Discard, []string{"1", "2", "done"})
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})

	t.Run("status first", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one")

		err := cli.Mark(io.Discard, []string{"done", "1"})
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})

	t.Run("message, single task", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one")

		status := "in-progress"
		buf := bytes.Buffer{}

		err := cli.Mark(&buf, []string{"1", status})
		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := fmt.Sprintf("task 1 marked %q\n", status)

		if got != want {
			t.Errorf("got message %q, want %q", got, want)
		}
	})

	t.Run("message, multiple tasks", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one", "two")

		status := "in-progress"
		buf := bytes.Buffer{}

		err := cli.Mark(&buf, []string{"1", "2", status})
		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := fmt.Sprintf("task 1 marked %q\ntask 2 marked %q\n", status, status)

		if got != want {
			t.Errorf("got message %q, want %q", got, want)
		}
	})
}

func TestCLI_Update(t *testing.T) {
	t.Run("successful", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one", "two")

		description := "eat tacos"

		err := cli.Update(io.Discard, []string{"1", description})
		if err != nil {
			t.Fatal(err)
		}

		got := cli.tasks.Get()

		if got[0].Description != description {
			t.Errorf("got message %q, want %q", got[0].Description, description)
		}
	})

	t.Run("message", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one", "two")

		buf := bytes.Buffer{}

		err := cli.Update(&buf, []string{"2", "party"})
		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := "task 2 description updated to \"party\"\n"

		if got != want {
			t.Errorf("got message %s, want %q", got, want)
		}
	})

	t.Run("more than two arguments", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one", "two")

		err := cli.Update(io.Discard, []string{"2", "1", "party"})
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})

	t.Run("less than two arguments", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one", "two")

		err := cli.Update(io.Discard, []string{"1"})
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		cli := NewCLI(filename)
		cli.tasks.Add("", "one", "two")

		err := cli.Update(io.Discard, []string{"3", "pizza party"})
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})
}
