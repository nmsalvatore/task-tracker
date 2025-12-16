package main

import (
	"errors"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
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
	err := validateAddFlags(args)
	if err != nil {
		return err
	}

	status, descriptions, err := parseAddArgs(args)
	if err != nil {
		return err
	}

	err = c.tasks.Add(status, descriptions...)
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

func (c *CLI) Delete(writer io.Writer, args []string) error {
	if len(args) == 0 {
		return errors.New("no task ID provided")
	}

	ids, err := argsToInts(args)
	if err != nil {
		return err
	}

	err = c.tasks.Delete(ids...)
	if err != nil {
		return err
	}

	for _, id := range ids {
		fmt.Fprintf(writer, "Deleted task %d\n", id)
	}
	return nil
}

func (c *CLI) List(writer io.Writer, args []string) error {
	if len(args) > 1 {
		return errors.New("too many arguments")
	}

	if len(args) == 0 {
		PrintTasks(writer, c.tasks.Get())
		return nil
	}

	tasks, err := c.tasks.GetByStatus(args[0])
	if err != nil {
		return err
	}

	PrintTasks(writer, tasks)
	return nil
}

func (c *CLI) Mark(writer io.Writer, args []string) error {
	li := len(args) - 1
	status := args[li]

	ids, err := argsToInts(args[:li])
	if err != nil {
		return err
	}

	err = c.tasks.Mark(status, ids...)
	if err != nil {
		return err
	}

	for _, id := range ids {
		fmt.Fprintf(writer, "Task %d marked %q\n", id, status)
	}
	return nil
}

func argsToInts(args []string) ([]int, error) {
	nums := make([]int, len(args))

	for i := range args {
		num, err := strconv.Atoi(args[i])
		if err != nil {
			return nil, fmt.Errorf("convert string to int: %v", err)
		}
		nums[i] = num
	}

	return nums, nil
}

func parseAddArgs(args []string) (status string, descriptions []string, err error) {
	for i := 0; i < len(args); i++ {
		if args[i] == "--status" {
			if i+1 >= len(args) {
				return "", nil, errors.New("--status flag requires a value")
			}
			status = args[i+1]
			i++
		} else {
			descriptions = append(descriptions, args[i])
		}
	}

	return
}

func validateAddFlags(args []string) error {
	if slices.ContainsFunc(args, func(arg string) bool {
		return strings.HasPrefix(arg, "--") && arg != "--status"
	}) {
		for _, arg := range args {
			if strings.HasPrefix(arg, "--") && arg != "--status" {
				return fmt.Errorf("unknown flag: %s", arg)
			}
		}
	}
	return nil
}
