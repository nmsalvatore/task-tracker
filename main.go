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
	tasks := Tasks{}

	err := tasks.Add("", "make dinner", "go on a walk")
	if err != nil {
		return fmt.Errorf("failed to add tasks: %v", err)
	}

	err = tasks.Add("done", "drink coffee")
	if err != nil {
		return fmt.Errorf("failed to add tasks: %v", err)
	}

	tasks.List(os.Stdout)
	return nil
}
