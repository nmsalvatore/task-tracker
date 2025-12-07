package main

import "testing"

func TestAddTask(t *testing.T) {
	t.Run("success message", func(t *testing.T) {
		got := AddTask("drink coffee", "")
		want := "todo: drink coffee"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
