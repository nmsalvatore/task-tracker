package main

import (
	"fmt"
	"log"
	"os"
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
		err = cli.List(os.Stdout, args)
	case "mark":
		err = cli.Mark(os.Stdout, args)
	case "update":
		err = cli.Update(os.Stdout, args)
	}

	err = cli.tasks.Save(cli.filename)
	if err != nil {
		return fmt.Errorf("save tasks: %v", err)
	}

	return nil
}
