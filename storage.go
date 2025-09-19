package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	TaskFile = "tasks.json"
)

func load() (TaskList, error) {
	data, err := os.ReadFile(TaskFile)
	if err != nil {
		return nil, fmt.Errorf("failed read tasks: %s", err.Error())
	}
	var t []*Task
	err = json.Unmarshal(data, &t)
	if err != nil {
		return nil, fmt.Errorf("failed unmarshal tasks: %s", err.Error())
	}
	return t, nil
}
func save(t TaskList) error {
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Errorf("failed marshal tasks: %s", err.Error())
	}
	if err := os.WriteFile("tasks.json", data, 0644); err != nil {
		return fmt.Errorf("failed write tasks: %s", err.Error())
	}
	return nil
}
