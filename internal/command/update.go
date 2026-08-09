package command

import (
	"fmt"
	"task-tracker-cli/internal/storage"
	"time"
)

func UpdateTask(id int, name string) {
	tasks, err := storage.ReadTasksFromFile()
	if err != nil {
		fmt.Printf("Error reading tasks: %v\n", err)
		return
	}
	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = name
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Task with ID %d is not found.\n", id)
		return
	}
	err = storage.WriteTasksToFile(tasks)
	if err != nil {
		fmt.Printf("Error writing tasks: %v\n", err)
		return
	}
	fmt.Println("Task updated successfully")
}

func DeleteTask(id int) {
	tasks, err := storage.ReadTasksFromFile()
	if err != nil {
		fmt.Printf("Failed to read JSON file: %v\n", err)
		return
	}
	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Task with ID %d is not found.\n", id)
		return
	}
	err = storage.WriteTasksToFile(tasks)
	if err != nil {
		fmt.Printf("Error writing tasks: %v\n", err)
		return
	}
	fmt.Println("Task deleted successfully")
}
