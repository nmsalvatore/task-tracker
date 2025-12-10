package main

import (
	"fmt"
	"io"
	"strings"
)

func PrintTasks(writer io.Writer, tasks []Task) {
	// TODO: pass "by status, done" test,
	// 		 write tests got Get to filter by status, then put status in Get(status)
	// 		 then update all Get calls to take an empty string or a status
	//
	// TODO: check for invalid status

	if len(tasks) == 0 {
		fmt.Fprintln(writer, "no tasks")
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
