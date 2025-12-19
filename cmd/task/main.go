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
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("getting home directory: %v", err)
	}

	dir := home + "/.tasks"
	err = os.Mkdir(dir, 0750)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("creating data directory: %v", err)
	}

	cli := NewCLI(dir + "/tasks.json")
	err = cli.tasks.Load(cli.filename)
	if err != nil {
		return fmt.Errorf("loading tasks: %v", err)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "add":
		err = cli.Add(os.Stdout, args)
	case "clear":
		err = cli.Clear(os.Stdout, args)
	case "delete":
		err = cli.Delete(os.Stdout, args)
	case "help":
		err = cli.Help(os.Stdout, args)
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
