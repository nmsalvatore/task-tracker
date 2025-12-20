package main

import (
	"fmt"
	"io"
	"strings"
)

func PrintHelp(writer io.Writer, args []string) {
	if len(args) == 0 {
		PrintAppHelp(writer)
		return
	}

	cmd := args[0]
	switch cmd {
	case "add":
		PrintAddHelp(writer)
	case "clear":
		PrintClearHelp(writer)
	case "delete":
		PrintDeleteHelp(writer)
	case "help":
		PrintHelpHelp(writer)
	case "list":
		PrintListHelp(writer)
	case "mark":
		PrintMarkHelp(writer)
	case "update":
		PrintUpdateHelp(writer)
	case "version":
		PrintVersionHelp(writer)
	default:
		fmt.Fprintf(writer, "no command '%s'\n", cmd)
	}
}

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
  help		Application usage information
  version	Application version

Use 'task help COMMAND' for specific usage information about the above commands.
`

	fmt.Fprint(writer, message)
}

func PrintAddHelp(writer io.Writer) {
	message := `Add one or more tasks to the task list.

usage: task add DESCRIPTIONS [OPTIONS]

options:
  --status	Set the task status (todo, in-progress, done)

examples:
  task add "go to the store"
  task add "meal prep" "go on a walk"
  task add "drink coffee" --status done
  task add --status in-progress "play with task tracker"
`
	fmt.Fprint(writer, message)
}

func PrintClearHelp(writer io.Writer) {
	message := `Clear tasks from the task list, either entirely or by status.

usage: task clear [STATUS]

examples:
  task clear
  task clear done
  task clear in-progress
  task clear todo
`
	fmt.Fprint(writer, message)
}

func PrintDeleteHelp(writer io.Writer) {
	message := `Delete individual tasks from the task list.

usage: task delete IDS

examples:
  task delete 1
  task delete 1 2 3
`
	fmt.Fprint(writer, message)
}

func PrintHelpHelp(writer io.Writer) {
	message := `Usage information about the application or a specific command.

usage: task help [COMMAND]

examples:
  task help
  task help add
`
	fmt.Fprint(writer, message)
}

func PrintListHelp(writer io.Writer) {
	message := `List tasks, either entirely or by status.

usage: task list [STATUS]

examples:
  task list
  task list todo
  task list in-progress
  task list done
`
	fmt.Fprint(writer, message)
}

func PrintMarkHelp(writer io.Writer) {
	message := `Mark task status. Status must be todo, in-progress, or done.

usage: task mark IDS STATUS

examples:
  task mark 1 done
  task mark 1 2 3 in-progress
  task mark 4 todo
`
	fmt.Fprint(writer, message)
}

func PrintUpdateHelp(writer io.Writer) {
	message := `Update task description.

usage: task update ID DESCRIPTION

examples:
  task update 1 "new task description"
`
	fmt.Fprint(writer, message)
}

func PrintVersionHelp(writer io.Writer) {
	message := `Display application version.

usage: task version

examples:
  task version
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

func PrintVersion() {
	fmt.Printf("%s, version %s\n", appName, version)
}

func Bold(s string) string {
	return "\033[1m" + s + "\033[0m"
}

func Strike(s string) string {
	return "\033[9m" + s + "\033[0m"
}
