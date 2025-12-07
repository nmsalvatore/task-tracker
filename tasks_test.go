package main

import (
	"bytes"
	"testing"
)

func TestListTasks(t *testing.T) {
	t.Run("no tasks message", func(t *testing.T) {
		buffer := bytes.Buffer{}
		ListTasks(&buffer)

		got := buffer.String()
		want := "no tasks\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
