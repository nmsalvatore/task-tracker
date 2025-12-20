package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestDisplay_HelpMessages(t *testing.T) {
	t.Run("app help message", func(t *testing.T) {
		buf := bytes.Buffer{}
		PrintAppHelp(&buf)

		got := buf.String()
		want := "A simple command line task tracker"

		if !strings.Contains(got, want) {
			t.Errorf("%q not found in %q", want, got)
		}
	})

	t.Run("add help message", func(t *testing.T) {
		buf := bytes.Buffer{}
		PrintAddHelp(&buf)

		got := buf.String()
		want := []string{
			"usage: task add DESCRIPTIONS",
			`task add "go to the store"`,
			`task add "meal prep" "go on a walk"`,
		}

		for i := range want {
			if !strings.Contains(got, want[i]) {
				t.Errorf("%q not found in %q", want[i], got)
			}
		}
	})

	t.Run("clear help message", func(t *testing.T) {
		buf := bytes.Buffer{}
		PrintClearHelp(&buf)

		got := buf.String()
		want := []string{
			"usage: task clear [STATUS]",
			"task clear",
			"task clear done",
			"task clear in-progress",
		}

		for i := range want {
			if !strings.Contains(got, want[i]) {
				t.Errorf("%q not found in %q", want[i], got)
			}
		}
	})

	t.Run("delete help message", func(t *testing.T) {
		buf := bytes.Buffer{}
		PrintDeleteHelp(&buf)

		got := buf.String()
		want := []string{
			"usage: task delete IDS",
			"task delete 1",
			"task delete 1 2 3",
		}

		for i := range want {
			if !strings.Contains(got, want[i]) {
				t.Errorf("%q not found in %q", want[i], got)
			}
		}
	})

	t.Run("help help message", func(t *testing.T) {
		buf := bytes.Buffer{}
		PrintHelpHelp(&buf)

		got := buf.String()
		want := []string{
			"usage: task help [COMMAND]",
			"task help add",
		}

		for i := range want {
			if !strings.Contains(got, want[i]) {
				t.Errorf("%q not found in %q", want[i], got)
			}
		}
	})

	t.Run("list help message", func(t *testing.T) {
		buf := bytes.Buffer{}
		PrintListHelp(&buf)

		got := buf.String()
		want := []string{
			"usage: task list [STATUS]",
			"task list",
			"task list in-progress",
		}

		for i := range want {
			if !strings.Contains(got, want[i]) {
				t.Errorf("%q not found in %q", want[i], got)
			}
		}
	})

	t.Run("mark help message", func(t *testing.T) {
		buf := bytes.Buffer{}
		PrintMarkHelp(&buf)

		got := buf.String()
		want := []string{
			"usage: task mark IDS STATUS",
			"task mark 1 done",
			"task mark 1 2 3 in-progress",
			"task mark 4 todo",
		}

		for i := range want {
			if !strings.Contains(got, want[i]) {
				t.Errorf("%q not found in %q", want[i], got)
			}
		}
	})

	t.Run("update help message", func(t *testing.T) {
		buf := bytes.Buffer{}
		PrintUpdateHelp(&buf)

		got := buf.String()
		want := []string{
			"usage: task update ID DESCRIPTION",
			`task update 1 "new task description"`,
		}

		for i := range want {
			if !strings.Contains(got, want[i]) {
				t.Errorf("%q not found in %q", want[i], got)
			}
		}
	})

	t.Run("version help message", func(t *testing.T) {
		buf := bytes.Buffer{}
		PrintVersionHelp(&buf)

		got := buf.String()
		want := "usage: task version"

		if !strings.Contains(got, want) {
			t.Errorf("%q not found in %q", want, got)
		}
	})
}

func TestDisplay_PrintTasks(t *testing.T) {
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
