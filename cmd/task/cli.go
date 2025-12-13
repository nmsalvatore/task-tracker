package main

import (
	"errors"
	"fmt"
	"io"
)

type CLI struct {
	tasks    *Tasks
	filename string
}

func NewCLI(filename string) *CLI {
	return &CLI{
		tasks:    &Tasks{},
		filename: filename,
	}
}

func (c *CLI) Add(writer io.Writer, args []string) error {
	var descriptions []string
	var status string

	for i := 0; i < len(args); i++ {
		if args[i] == "--status" {
			if i+1 >= len(args) {
				return errors.New("--status flag requires a value")
			}
			status = args[i+1]
			i++
		} else {
			descriptions = append(descriptions, args[i])
		}
	}

	err := c.tasks.Add(status, descriptions...)
	if err != nil {
		return err
	}

	for _, task := range descriptions {
		fmt.Fprintf(writer, "Added task %q\n", task)
	}

	return nil
}

func (c *CLI) Clear(writer io.Writer, args []string) error {
	if len(args) == 0 {
		c.tasks.Clear()
		fmt.Fprintln(writer, "Cleared all tasks")
		return nil
	}

	status := args[0]
	err := c.tasks.ClearByStatus(status)
	if err != nil {
		return err
	}

	fmt.Fprintf(writer, "Cleared all tasks with status %q\n", status)
	return nil
}
