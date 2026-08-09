package command

import (
	"fmt"
	"task-tracker-cli/internal/model"
	"task-tracker-cli/internal/storage"
	"time"
)

func MarkTaskStatus(id int, status model.TaskStatus) {
	tasks, err := storage.ReadTasksFromFile()
	if err != nil {
		fmt.Printf("Error reading tasks: %v\n", err)
		return
	}
	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Task with ID %d not found.\n", id)
		return

	}
	err = storage.WriteTasksToFile(tasks)
	if err != nil {
		fmt.Printf("Error writing tasks: %v\n", err)
		return
	}
	fmt.Printf("Task marked as %s successfully\n", status)
}
