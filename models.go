package main

import (
	"fmt"
	"os"
)

type Task struct {
	ID          int
	Description string
	Completed   bool
}

type TaskList = []*Task

func NewTask(description string) error {
	tasks, err := load()
	if err != nil {
		return fmt.Errorf("failed to load tasks: %v", err)
	}
	task := Task{ID: len(tasks) + 1, Description: description, Completed: false}
	tasks = append(tasks, &task)
	if err := save(tasks); err != nil {
		return fmt.Errorf("failed to save tasks: %v", err)
	}
	return nil
}

func GetTasks(filter TaskListType) TaskList {
	t, err := load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load tasks: %v", err)
	}
	switch filter {
	case CompletedTasks:
		return filterTasks(t, func(t *Task) bool { return t.Completed })
	case UncompletedTasks:
		return filterTasks(t, func(t *Task) bool { return !t.Completed })
	case AllTasks:
		fallthrough
	default:
		return t
	}
}

func DeleteTask(id int) error {
	tasks, err := load()
	if err != nil {
		return fmt.Errorf("failed to load tasks: %v", err)
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	if err := save(tasks); err != nil {
		return fmt.Errorf("failed to save tasks: %v", err)
	}
	return nil
}
func CompleteTask(id int) error {
	tasks, err := load()
	if err != nil {
		return fmt.Errorf("failed to load tasks: %v", err)
	}
	for _, task := range tasks {
		if task.ID == id {
			task.Completed = true
			break
		}
	}
	if err := save(tasks); err != nil {
		return fmt.Errorf("failed to save tasks: %v", err)
	}
	return nil
}
func filterTasks(tasks TaskList, filter func(*Task) bool) TaskList {
	var filteredTasks TaskList
	for _, task := range tasks {
		if filter(task) {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}
