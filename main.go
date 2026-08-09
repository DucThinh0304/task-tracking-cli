package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"task-tracker-cli/internal/command"
	"task-tracker-cli/internal/model"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		return
	}

	action := os.Args[1]

	if strings.HasPrefix(action, "mark-") {
		rawStatus := strings.TrimPrefix(action, "mark-")
		status := model.TaskStatus(rawStatus)
		if !status.IsValid() {
			fmt.Printf("Invalid status: '%s'\n", status)
			return
		}
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide an id")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: id must be a number")
			return
		}
		command.MarkTaskStatus(id, status)

	} else {
		switch action {
		case "add":
			if len(os.Args) < 3 {
				fmt.Println("Error: Please provide a task name")
				return
			}
			name := os.Args[2]
			command.AddTask(name)
		case "list":
			if len(os.Args) < 3 {
				command.ListTasks()
			} else {
				status := model.TaskStatus(os.Args[2])
				if !status.IsValid() {
					fmt.Printf("Invalid status: '%s'\n", status)
					return
				}
				command.ListTasksWithStatus(status)
			}
		case "update":
			if len(os.Args) < 4 {
				fmt.Println("Error: Please provide a task id and a name")
				return
			}
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("Error: id must be a number")
				return
			}
			command.UpdateTask(id, os.Args[3])
		case "delete":
			if len(os.Args) < 3 {
				fmt.Println("Error: Please provide a task id")
				return
			}
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("Error: id must be a number")
				return
			}
			command.DeleteTask(id)
		default:
			fmt.Printf("Unknown command: %s\nUsage: task-cli <command> [arguments]\n", action)
		}
	}
}
