package main

import "fmt"

func AddTask(task string, status string) string {
	if status == "" {
		status = "todo"
	}
	return fmt.Sprintf("%s: %s", status, task)
}
