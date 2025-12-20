package main

import (
	"fmt"
	"log"
	"os"
)

const (
	appName = "task"
	version = "0.1.0"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: task <command> [arguments]")
		os.Exit(1)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "version":
		PrintVersion()
	case "help":
		PrintHelp(os.Stdout, args)
	default:
		err := run(cmd, args)
		if err != nil {
			log.Fatalf("task: %v", err)
		}
	}
}

func run(cmd string, args []string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("getting home directory: %v", err)
	}

	dir := fmt.Sprintf("%s/.%s", home, appName)
	err = os.Mkdir(dir, 0750)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("creating data directory: %v", err)
	}

	cli := NewCLI(dir + "/tasks.json")
	err = cli.tasks.Load(cli.filename)
	if err != nil {
		return fmt.Errorf("loading tasks: %v", err)
	}

	switch cmd {
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
	default:
		err = fmt.Errorf("no command '%s'", cmd)
	}

	if err != nil {
		return err
	}

	err = cli.tasks.Save(cli.filename)
	if err != nil {
		return fmt.Errorf("saving tasks: %v", err)
	}

	return nil
}
