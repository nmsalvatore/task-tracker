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
		switch task.Status {
		case "todo":
			message.WriteString("• ")
		case "in-progress":
			message.WriteString("> ")
		case "done":
			message.WriteString("× ")
		}
		item := fmt.Sprintf("%d: %s\n", task.ID, task.Description)
		message.WriteString(item)
	}

	fmt.Fprint(writer, message)
}
