# Task Tracker

A simple CLI task tracker built with Go. This is one of the [roadmap.sh Backend Projects](https://roadmap.sh/backend/projects).

## Usage

```bash
# Add tasks
task add "Go to the moon"
task add "Do a little dance" "Make a little love" "Get down tonight"
task add "Drink coffee" --done
task add "Build a task tracker" --doing
task add "Get a real job" --todo

# Edit tasks
task update 1 "Go to Discovery Zone"

# Delete tasks
task delete 1
task delete --all
task delete --done
task delete --doing
task delete --todo

# Mark task status
task doing 5
task done 6
task todo 5

# List tasks
task list
task list --done
task list --doing
task list --todo

# Helpers
task version
task help
task tip
```
