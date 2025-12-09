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

	err := tasks.Add("", "update fastrak", "go on a walk", "something unimportant", "drink whiskey")
	if err != nil {
		return fmt.Errorf("add tasks: %v", err)
	}

	err = tasks.Add("done", "drink coffee")
	if err != nil {
		return fmt.Errorf("add tasks: %v", err)
	}

	err = tasks.Mark("done", 2, 4)
	if err != nil {
		return fmt.Errorf("mark task status: %v", err)
	}

	err = tasks.Delete(3)
	if err != nil {
		return fmt.Errorf("delete task: %v", err)
	}

	err = tasks.Add("done", "go to work")
	if err != nil {
		return fmt.Errorf("add tasks: %v", err)
	}

	tasks.List(os.Stdout, "")
	return nil
}
