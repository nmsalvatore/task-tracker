package main

import (
	"fmt"
	"io"
	"strings"
)

func PrintTasks(writer io.Writer, tasks []Task) {
	message := &strings.Builder{}

	if len(tasks) == 0 {
		message.WriteString("task list empty\n")
	}

	for _, task := range tasks {
		var description string

		switch task.Status {
		case "todo":
			message.WriteString("• ")
			description = task.Description
		case "in-progress":
			message.WriteString("> ")
			description = Bold(task.Description)
		case "done":
			message.WriteString("× ")
			description = Strike(task.Description)
		}

		item := fmt.Sprintf("%d: %s\n", task.ID, description)
		message.WriteString(item)
	}

	fmt.Fprint(writer, message)
}

func Bold(s string) string {
	return "\033[1m" + s + "\033[0m"
}

func Strike(s string) string {
	return "\033[9m" + s + "\033[0m"
}
