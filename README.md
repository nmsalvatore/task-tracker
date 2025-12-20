# Task Tracker

A simple CLI task tracker built with Go. This is one of the roadmap.sh [Backend Projects](https://roadmap.sh/backend/projects). The project requirements are outlined [here](https://roadmap.sh/projects/task-tracker).

## Installation

Clone this repo and `cd` into the project's root directory. Build the binary with `make`. Move the binary to `/usr/local/bin` or wherever you like to keep your binaries. Do some tasks!

## Usage

```bash
# Add tasks
task add "Go to the moon"
task add "Do a little dance" "Make a little love" "Get down tonight"
task add "Drink coffee" --status done
task add "Build a task tracker" --status in-progress

# Update task
task update 1 "Get a real job"

# Delete task(s)
task delete 1
task delete 1 2 3

# Clear tasks
task clear
task clear done 
task clear in-progress
task clear todo

# Mark task status
task mark 5 done
task mark 3 4 5 done
task mark 6 in-progress
task mark 5 todo

# List tasks
task list
task list done
task list in-progress
task list todo

# Helpers
task version
task help
```
