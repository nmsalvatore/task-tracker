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

	err := tasks.Add("", "update fastrak", "update mosaic of planes", "go to the gym")
	if err != nil {
		return fmt.Errorf("add tasks: %v", err)
	}

	PrintTasks(os.Stdout, tasks.items)
	return nil
}
