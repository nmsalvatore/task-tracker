package main

import (
	"os"
)

func main() {
	tasks := Tasks{}
	tasks.List(os.Stdout)
}
