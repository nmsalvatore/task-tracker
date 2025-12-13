package main

import (
	"fmt"
	"io"
	"strings"
)

func PrintTasks(writer io.Writer, tasks []Task) {
	if len(tasks) == 0 {
		fmt.Fprintln(writer, "Task list is empty")
		return
	}

	message := &strings.Builder{}
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
