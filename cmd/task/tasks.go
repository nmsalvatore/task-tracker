package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
	"time"
)

type Task struct {
	ID          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Tasks struct {
	items []Task
}

func (t *Tasks) Add(status string, items ...string) error {
	if status == "" {
		status = "todo"
	}

	err := t.validateStatus(status)
	if err != nil {
		return fmt.Errorf("couldn't add task: %v", err)
	}

	for _, item := range items {
		now := time.Now()
		task := Task{
			ID:          t.getMaxID() + 1,
			Description: item,
			Status:      status,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		t.items = append(t.items, task)
	}

	return nil
}

func (t *Tasks) Clear() {
	t.items = []Task{}
}

func (t *Tasks) ClearByStatus(status string) error {
	err := t.validateStatus(status)
	if err != nil {
		return fmt.Errorf("couldn't clear tasks: %v", err)
	}

	t.items = slices.DeleteFunc(t.items, func(task Task) bool {
		return task.Status == status
	})

	return nil
}

func (t *Tasks) Delete(ids ...int) error {
	startLen := len(t.items)

	t.items = slices.DeleteFunc(t.items, func(task Task) bool {
		return slices.Contains(ids, task.ID)
	})

	if len(t.items) == startLen {
		return errors.New("task not found")
	}
	return nil
}

func (t *Tasks) Get() []Task {
	return t.items
}

func (t *Tasks) GetByStatus(status string) ([]Task, error) {
	err := t.validateStatus(status)
	if err != nil {
		return nil, fmt.Errorf("couldn't get tasks: %v", err)
	}

	var tasks []Task
	for _, item := range t.items {
		if item.Status == status {
			tasks = append(tasks, item)
		}
	}

	return tasks, nil
}

func (t *Tasks) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("read data file: %v", err)
	}

	err = json.Unmarshal(data, &t.items)
	if err != nil {
		return fmt.Errorf("decode json file: %v", err)
	}

	return nil
}

func (t *Tasks) Mark(status string, ids ...int) error {
	if status == "" {
		return errors.New("mark status empty")
	}

	err := t.validateStatus(status)
	if err != nil {
		return err
	}

	var updated int
	for i := range t.items {
		if slices.Contains(ids, t.items[i].ID) {
			t.items[i].Status = status
			t.items[i].UpdatedAt = time.Now()
			updated++
		}
	}

	if updated == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (t *Tasks) Save(filename string) error {
	data, err := json.MarshalIndent(t.items, "", "  ")
	if err != nil {
		return fmt.Errorf("encode json: %v", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("write data file: %v", err)
	}

	return nil
}

func (t *Tasks) Update(id int, description string) error {
	for i := range t.items {
		if t.items[i].ID == id {
			t.items[i].Description = description
			t.items[i].UpdatedAt = time.Now()
			return nil
		}
	}
	return errors.New("task not found")
}

func (t *Tasks) getMaxID() (max int) {
	for _, task := range t.items {
		if task.ID > max {
			max = task.ID
		}
	}
	return
}

func (t *Tasks) validateStatus(status string) error {
	valid := []string{"todo", "in-progress", "done"}
	if !slices.Contains(valid, status) {
		return fmt.Errorf("invalid status: %q (must be todo, in-progress, or done)", status)
	}
	return nil
}
