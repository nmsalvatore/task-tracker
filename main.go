package main

import (
	"os"
)

func main() {
	tasks := Tasks{}
	tasks.Add("drink coffee", "go on a walk")
	tasks.List(os.Stdout)
}
