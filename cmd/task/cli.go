package main

import (
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
	err := c.tasks.Add("", args...)
	if err != nil {
		return err
	}

	for _, task := range args {
		fmt.Fprintf(writer, "Added task %q\n", task)
	}
	return nil
}

func (c *CLI) Clear(args []string) error {
	if len(args) == 0 {
		c.tasks.Clear()
		fmt.Println("Cleared all tasks")
		return nil
	}

	status := args[0]
	err := c.tasks.ClearByStatus(status)
	if err != nil {
		return fmt.Errorf("clear task by status: %v", err)
	}

	fmt.Printf("Cleared all tasks with status %q\n", status)
	return nil
}
