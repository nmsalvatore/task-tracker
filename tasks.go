package main

func AddTask(task string, status string) string {
	if status == "" {
		status = "todo"
	}
	return status + ": " + task
}
