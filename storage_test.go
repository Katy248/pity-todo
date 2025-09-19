package main

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	err := os.WriteFile("tasks.json", []byte(`[{"id":1,"name":"Task 1","completed":false},{"id":2,"name":"Task 2","completed":true}]`), 0644)
	if err != nil {
		t.Fatal(err)
	}
	tasks, err := load()
	if err != nil {
		t.Fatal(err)
	}
	if len(tasks) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(tasks))
	}
}
