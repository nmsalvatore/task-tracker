package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("run failed: %v", err)
	}
}

func run() error {
	cli := NewCLI("tasks.json")

	err := cli.tasks.Load(cli.filename)
	if err != nil {
		return fmt.Errorf("load tasks: %v", err)
	}

	args := os.Args[2:]

	switch os.Args[1] {
	case "add":
		err = cli.Add(os.Stdout, args)
	case "clear":
		err = cli.Clear(os.Stdout, args)
	case "delete":
		err = cli.Delete(os.Stdout, args)
	case "list":
		PrintTasks(os.Stdout, cli.tasks.Get())
	case "mark":
		lastIndex := len(os.Args) - 1
		status := os.Args[lastIndex]

		ids, err := argsToInts(os.Args[2:lastIndex])
		if err != nil {
			return fmt.Errorf("mark task: %v", err)
		}

		err = cli.tasks.Mark(status, ids...)
		if err != nil {
			return fmt.Errorf("mark task: %v", err)
		}
	case "update":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return fmt.Errorf("convert string to int: %v", err)
		}

		err = cli.tasks.Update(id, os.Args[3])
		if err != nil {
			return fmt.Errorf("update task: %v", err)
		}
	}

	err = cli.tasks.Save(cli.filename)
	if err != nil {
		return fmt.Errorf("save tasks: %v", err)
	}

	return nil
}
