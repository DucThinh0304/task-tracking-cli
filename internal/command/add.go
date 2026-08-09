package command

import (
	"fmt"
	"task-tracker-cli/internal/model"
	"task-tracker-cli/internal/storage"
)

func AddTask(name string) {
	tasks, err := storage.ReadTasksFromFile()
	if err != nil {
		fmt.Printf("Error reading tasks: %v\n", err)
		return
	}
	id := GenerateID(tasks)
	tasks = append(tasks, model.NewTask(id, name))
	err = storage.WriteTasksToFile(tasks)
	if err != nil {
		fmt.Printf("Error writing tasks: %v\n", err)
		return
	}
	fmt.Printf("Task added successfully (ID: %d)\n", id)
}

func GenerateID(tasks []model.Task) int {
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}
