package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("failure: %v", err)
	}
}

func run() error {
	tasks := Tasks{}

	err := tasks.Add("", "make dinner", "go on a walk", "drink whiskey")
	if err != nil {
		return fmt.Errorf("add tasks: %v", err)
	}

	err = tasks.Add("done", "drink coffee")
	if err != nil {
		return fmt.Errorf("add tasks: %v", err)
	}

	err = tasks.Mark(3, "in-progress")
	if err != nil {
		return fmt.Errorf("mark task status: %v", err)
	}

	tasks.List(os.Stdout)
	return nil
}
