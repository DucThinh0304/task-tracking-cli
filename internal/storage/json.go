package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"task-tracker-cli/internal/model"
)

func GetStoragePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "tasks.json", nil
	}

	dirPath := filepath.Join(homeDir, ".task-tracker-cli")
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return "", err
	}
	return filepath.Join(dirPath, "tasks.json"), nil

}
func WriteTasksToFile(tasks []model.Task) error {
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	storagePath, err := GetStoragePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(storagePath, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadTasksFromFile() ([]model.Task, error) {
	var tasks []model.Task
	storagePath, err := GetStoragePath()
	if err != nil {
		return nil, err
	}
	file, err := os.OpenFile(storagePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fileInfo, _ := file.Stat()

	if fileInfo.Size() == 0 {
		return nil, nil
	}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return tasks, err
	}
	return tasks, nil
}
