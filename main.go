package main

import (
	"os"
)

func main() {
	tasks := Tasks{
		items: []Task{
			{Description: "drink coffee"},
			{Description: "go on a walk"},
		},
	}
	tasks.List(os.Stdout)
}
