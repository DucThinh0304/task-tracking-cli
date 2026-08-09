package command

import (
	"fmt"
	"task-tracker-cli/internal/model"
	"task-tracker-cli/internal/storage"
)

func ListTasks() {

	tasks, err := storage.ReadTasksFromFile()
	if err != nil {
		fmt.Printf("Error reading tasks: %v\n", err)
		return
	}
	if len(tasks) == 0 {
		fmt.Println("There's no task left, add more using: task-cli add [arguments])")
	}
	for _, task := range tasks {
		fmt.Println(task.String())
	}
}

func ListTasksWithStatus(status model.TaskStatus) {
	tasks, err := storage.ReadTasksFromFile()
	if err != nil {
		fmt.Printf("Error reading tasks: %v\n", err)
		return
	}
	found := false
	for _, task := range tasks {
		if task.Status == status {
			found = true
			fmt.Println(task.String())
		}
	}
	if !found {
		fmt.Printf("There's no task match status %s\n", status.String())
	}
}
