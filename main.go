package main

import (
	"log"
	"os"
)

func main() {
	tasks := Tasks{}

	err := tasks.Add("", "make dinner", "go on a walk")
	if err != nil {
		log.Fatalf("failed to add tasks: %v", err)
	}

	tasks.Add("done", "drink coffee")
	if err != nil {
		log.Fatalf("failed to add tasks: %v", err)
	}

	tasks.List(os.Stdout)
}
