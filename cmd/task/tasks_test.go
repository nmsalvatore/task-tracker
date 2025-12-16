package main

import (
	"slices"
	"testing"
)

func TestTasks_Add(t *testing.T) {
	t.Run("single task", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "drink coffee")

		got := tasks.Get()

		if len(got) != 1 {
			t.Errorf("got %d tasks, want 1", len(got))
		}
	})

	t.Run("multiple tasks", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "do a little dance", "make a little love", "get down tonight")

		got := tasks.Get()

		if len(got) != 3 {
			t.Errorf("got %d tasks, want 3", len(got))
		}
	})

	t.Run("id", func(t *testing.T) {
		tasks := Tasks{}
		description := "fly to the moon"
		tasks.Add("", description)

		got := tasks.Get()

		if got[0].ID != 1 {
			t.Errorf("got %d, want %d", got[0].ID, 1)
		}
	})

	t.Run("id after delete", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "one", "two", "three", "four")
		tasks.Delete(2)
		tasks.Add("", "five")

		got := tasks.Get()

		if got[3].ID != 5 {
			t.Errorf("got %d, want %d", got[3].ID, 5)
		}
	})

	t.Run("description", func(t *testing.T) {
		tasks := Tasks{}
		description := "drink coffee"
		tasks.Add("", description)

		got := tasks.Get()

		if got[0].Description != description {
			t.Errorf("got %q, want %q", got[0].Description, description)
		}
	})

	t.Run("status default", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "go on a walk")

		got := tasks.Get()

		if got[0].Status != "todo" {
			t.Errorf("got %q, want %q", got[0].Status, "todo")
		}
	})

	t.Run("status specified", func(t *testing.T) {
		tasks := Tasks{}
		status := "in-progress"
		tasks.Add(status, "go on a walk")

		got := tasks.Get()

		if got[0].Status != status {
			t.Errorf("got %q, want %q", got[0].Status, "in-progress")
		}
	})

	t.Run("status on multiples", func(t *testing.T) {
		tasks := Tasks{}
		status := "in-progress"
		tasks.Add(status, "go on a walk", "live the dream")

		got := tasks.Get()

		for i := range got {
			if got[i].Status != status {
				t.Errorf("got %q, want %q", got[i].Status, "in-progress")
			}
		}
	})

	t.Run("status rejected", func(t *testing.T) {
		tasks := Tasks{}
		err := tasks.Add("dude", "go on a walk")
		if err == nil {
			t.Error("wanted error but didn't get one")
		}
	})

	t.Run("created at", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first")
		tasks.Add("", "second")

		got := tasks.Get()

		if !got[0].CreatedAt.Before(got[1].CreatedAt) {
			t.Errorf("first not created before second")
		}
	})
}

func TestTasks_Clear(t *testing.T) {
	tasks := Tasks{}
	tasks.Add("", "first", "second", "third")
	tasks.Add("done", "fourth", "fifth")
	tasks.Clear()

	got := tasks.Get()

	if len(got) != 0 {
		t.Errorf("got length %d, want 0", len(got))
	}
}

func TestTasks_ClearByStatus(t *testing.T) {
	t.Run("valid status", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first", "second", "third")
		tasks.Add("done", "fourth", "fifth")
		tasks.ClearByStatus("done")

		got := tasks.Get()

		if len(got) != 3 {
			t.Fatalf("got length %d, want 3", len(got))
		}

		want := []string{"first", "second", "third"}
		for i, task := range got {
			if task.Description != want[i] {
				t.Errorf("got %q, want %q", task.Description, want[i])
			}
		}
	})

	t.Run("invalid status", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first", "second", "third")
		tasks.Add("done", "fourth", "fifth")

		err := tasks.ClearByStatus("beans")
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})
}

func TestTasks_Delete(t *testing.T) {
	t.Run("single task", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first", "second", "third")
		tasks.Delete(2)

		got := tasks.Get()

		for _, task := range got {
			if task.Description == "second" {
				t.Error("task not deleted")
			}
		}

		if len(got) == 3 {
			t.Error("list length unchanged")
		}
	})

	t.Run("multiple tasks", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first", "second", "third")
		tasks.Delete(2, 3)

		got := tasks.Get()

		deleted := []string{"second", "third"}
		for _, task := range got {
			if slices.Contains(deleted, task.Description) {
				t.Errorf("%q not deleted", task.Description)
			}
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		tasks := Tasks{}
		err := tasks.Delete(1)
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})

	t.Run("some valid ids, some invalid", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first", "second", "third")

		err := tasks.Delete(2, 3, 4)
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}

		got := tasks.Get()
		if len(got) != 3 {
			t.Errorf("got length %d, want 3", len(got))
		}
	})
}

func TestTasks_GetByStatus(t *testing.T) {
	t.Run("done", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first", "second", "third", "fourth")
		tasks.Mark("done", 3, 4)

		got, _ := tasks.GetByStatus("done")

		if len(got) != 2 {
			t.Fatalf("got %d tasks, want 2", len(got))
		}

		want := []string{"third", "fourth"}
		for i, task := range got {
			if task.Description != want[i] {
				t.Errorf("got %q, want %q", task.Description, want[i])
			}
		}
	})

	t.Run("in-progress", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first", "second", "third", "fourth")
		tasks.Mark("in-progress", 3, 4)

		got, _ := tasks.GetByStatus("in-progress")

		if len(got) != 2 {
			t.Fatalf("got %d tasks, want 2", len(got))
		}

		want := []string{"third", "fourth"}
		for i, task := range got {
			if task.Description != want[i] {
				t.Errorf("got %q, want %q", task.Description, want[i])
			}
		}
	})

	t.Run("todo", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first", "second", "third", "fourth")
		tasks.Mark("done", 3, 4)

		got, _ := tasks.GetByStatus("todo")

		if len(got) != 2 {
			t.Fatalf("got %d tasks, want 2", len(got))
		}

		want := []string{"first", "second"}
		for i, task := range got {
			if task.Description != want[i] {
				t.Errorf("got %q, want %q", task.Description, want[i])
			}
		}
	})

	t.Run("invalid status", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("done", "first", "second", "third", "fourth")
		tasks.Mark("todo", 3, 4)

		_, err := tasks.GetByStatus("beans")

		if err == nil {
			t.Errorf("wanted error, but didn't get one")
		}
	})
}

func TestTasks_Mark(t *testing.T) {
	t.Run("single task", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first", "second")
		tasks.Add("done", "third")

		tasks.Mark("in-progress", 1)
		tasks.Mark("done", 2)
		tasks.Mark("todo", 3)

		got := tasks.Get()

		for i, status := range []string{"in-progress", "done", "todo"} {
			if got[i].Status != status {
				t.Errorf("got %q, want %q", got[i].Status, status)
			}
		}
	})

	t.Run("multiple tasks", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first", "second", "third")

		status := "in-progress"
		tasks.Mark(status, 1, 2)

		got := tasks.Get()

		for i := range 2 {
			if got[i].Status != status {
				t.Errorf("got %q, want %q", got[i].Status, status)
			}
		}
	})

	t.Run("empty status", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first")

		err := tasks.Mark("", 1)
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})

	t.Run("invalid status", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first")

		err := tasks.Mark("gettin' it", 1)
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})

	t.Run("update time", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first")
		tasks.Mark("in-progress", 1)

		got := tasks.Get()

		if got[0].UpdatedAt.Equal(got[0].CreatedAt) {
			t.Error("time not updated")
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "sip tea")

		err := tasks.Mark("in-progress", 2)
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}

	})
}

func TestTasks_Update(t *testing.T) {
	t.Run("single task", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first", "third")
		tasks.Update(2, "second")

		got := tasks.Get()

		if got[1].Description != "second" {
			t.Errorf("got %q, want %q", got, "second")
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first", "third")

		err := tasks.Update(3, "second")
		if err == nil {
			t.Error("wanted error, but didn't get one")
		}
	})

	t.Run("timestamp updated", func(t *testing.T) {
		tasks := Tasks{}
		tasks.Add("", "first", "third")
		tasks.Update(2, "second")

		got := tasks.Get()

		if got[1].UpdatedAt.Equal(got[1].CreatedAt) {
			t.Error("timestamp not updated")
		}
	})
}
