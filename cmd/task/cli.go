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
		return fmt.Errorf("add: %v", err)
	}

	status, descriptions, err := parseAddArgs(args)
	if err != nil {
		return fmt.Errorf("add: %v", err)
	}

	err = c.tasks.Add(status, descriptions...)
	if err != nil {
		return fmt.Errorf("add: %v", err)
	}

	for _, task := range descriptions {
		if status != "" {
			fmt.Fprintf(writer, "added task %q, marked %s\n", task, status)
		} else {
			fmt.Fprintf(writer, "added task %q\n", task)
		}
	}

	return nil
}

func (c *CLI) Clear(writer io.Writer, args []string) error {
	if len(args) == 0 {
		c.tasks.Clear()
		fmt.Fprintln(writer, "cleared all tasks")
		return nil
	}

	status := args[0]

	err := c.tasks.ClearByStatus(status)
	if err != nil {
		return fmt.Errorf("clear by status: %v", err)
	}

	fmt.Fprintf(writer, "cleared all tasks marked %s\n", status)
	return nil
}

func (c *CLI) Delete(writer io.Writer, args []string) error {
	if len(args) == 0 {
		return errors.New("delete: no id provided")
	}

	ids, err := argsToInts(args)
	if err != nil {
		return fmt.Errorf("delete: %v", err)
	}

	err = c.tasks.Delete(ids...)
	if err != nil {
		return fmt.Errorf("delete: %v", err)
	}

	for _, id := range ids {
		fmt.Fprintf(writer, "task %d deleted\n", id)
	}
	return nil
}

func (c *CLI) List(writer io.Writer, args []string) error {
	if len(args) > 1 {
		return errors.New("list: too many arguments")
	}

	if len(args) == 0 {
		PrintTasks(writer, c.tasks.Get())
		return nil
	}

	tasks, err := c.tasks.GetByStatus(args[0])
	if err != nil {
		return fmt.Errorf("list: %v", err)
	}

	PrintTasks(writer, tasks)
	return nil
}

func (c *CLI) Mark(writer io.Writer, args []string) error {
	if len(args) < 2 {
		return errors.New("mark: missing arguments")
	}

	li := len(args) - 1
	status := args[li]

	ids, err := argsToInts(args[:li])
	if err != nil {
		return fmt.Errorf("mark: %v", err)
	}

	err = c.tasks.Mark(status, ids...)
	if err != nil {
		return fmt.Errorf("mark: %v", err)
	}

	for _, id := range ids {
		fmt.Fprintf(writer, "task %d marked %q\n", id, status)
	}
	return nil
}

func (c *CLI) Update(writer io.Writer, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("update: missing arguments")
	}

	if len(args) > 2 {
		return fmt.Errorf("update: too many arguments")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("update: %v", err)
	}

	description := args[1]
	err = c.tasks.Update(id, description)
	if err != nil {
		return fmt.Errorf("update: %v", err)
	}

	fmt.Fprintf(writer, "task %d description updated to %q\n", id, description)
	return nil
}

func argsToInts(args []string) ([]int, error) {
	nums := make([]int, len(args))

	for i := range args {
		num, err := strconv.Atoi(args[i])
		if err != nil {
			return nil, fmt.Errorf("converting arguments to ints: %v", err)
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
