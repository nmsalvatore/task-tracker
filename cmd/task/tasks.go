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
		return fmt.Errorf("clearing tasks: %v", err)
	}

	t.items = slices.DeleteFunc(t.items, func(task Task) bool {
		return task.Status == status
	})

	return nil
}

func (t *Tasks) Delete(ids ...int) error {
	err := t.validateIds(ids...)
	if err != nil {
		return err
	}

	t.items = slices.DeleteFunc(t.items, func(task Task) bool {
		return slices.Contains(ids, task.ID)
	})

	return nil
}

func (t *Tasks) Get() []Task {
	return t.items
}

func (t *Tasks) GetByStatus(status string) ([]Task, error) {
	err := t.validateStatus(status)
	if err != nil {
		return nil, err
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
		return fmt.Errorf("reading data file: %v", err)
	}

	err = json.Unmarshal(data, &t.items)
	if err != nil {
		return fmt.Errorf("decoding json file: %v", err)
	}

	return nil
}

func (t *Tasks) Mark(status string, ids ...int) error {
	if status == "" {
		return errors.New("no status provided")
	}

	err := t.validateStatus(status)
	if err != nil {
		return err
	}

	err = t.validateIds(ids...)
	if err != nil {
		return err
	}

	for i := range t.items {
		if slices.Contains(ids, t.items[i].ID) {
			t.items[i].Status = status
			t.items[i].UpdatedAt = time.Now()
		}
	}

	return nil
}

func (t *Tasks) Save(filename string) error {
	data, err := json.MarshalIndent(t.items, "", "  ")
	if err != nil {
		return fmt.Errorf("encoding json: %v", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("writing data file: %v", err)
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

func (t *Tasks) validateIds(ids ...int) error {
	if len(ids) == 0 {
		return errors.New("no id provided")
	}

	for _, id := range ids {
		if !slices.ContainsFunc(t.items, func(task Task) bool {
			return task.ID == id
		}) {
			return fmt.Errorf("invalid id: %d", id)
		}
	}
	return nil
}

func (t *Tasks) validateStatus(status string) error {
	valid := []string{"todo", "in-progress", "done"}
	if !slices.Contains(valid, status) {
		return fmt.Errorf("invalid status: %q (must be todo, in-progress, or done)", status)
	}
	return nil
}
