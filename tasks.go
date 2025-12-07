package main

import (
	"fmt"
	"io"
)

func ListTasks(writer io.Writer) {
	tasks := []string{}

	if len(tasks) == 0 {
		fmt.Fprintln(writer, "no tasks")
		return
	}

	fmt.Fprintln(writer, "so many tasks")
}
