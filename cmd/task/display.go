package main

import (
	"fmt"
	"io"
	"strings"
)

func PrintAppHelp(writer io.Writer) {
	message := `A simple command line task tracker.

task management:
  add		Add tasks
  clear		Clear tasks
  delete	Delete tasks
  list		List tasks
  mark		Mark task status
  update	Update task description

application information:
  help		Application and command information
  version	Application version
`

	fmt.Fprint(writer, message)
}

func PrintAddHelp(writer io.Writer) {
	message := `Add one or more tasks to the task list.

usage: task add DESCRIPTIONS [OPTIONS]

options:
  --status	Set the task status

examples:
  task add "go to the store"
  task add "meal prep" "go on a walk"
  task add "drink coffee" --status done
  task add --status in-progress "play with task tracker"
`
	fmt.Fprint(writer, message)
}

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
