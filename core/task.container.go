package core

import (
	"fmt"
	"sort"
)

var (
	Tasks TaskContainer
)

type TaskContainer struct {
	Data []Task
}

func (t *TaskContainer) Init() {
	t.Data = []Task{}
}

func (t *TaskContainer) Add(task Task) {
	t.Data = append(t.Data, task)
}

func (t *TaskContainer) Edit(idx int, description string) error {
	if idx < 0 || idx >= len(t.Data) {
		return fmt.Errorf("Task with number {%d} not exists", idx)
	}

	task := &t.Data[idx]
	task.Edit(description)

	return nil
}

func (t *TaskContainer) Complete(idx int) error {
	if idx < 0 || idx >= len(t.Data) {
		return fmt.Errorf("Task with number {%d} not exists", idx)
	}

	task := &t.Data[idx]
	task.Complete()
	t.Sort()

	return nil
}

func (t *TaskContainer) Uncomplete(idx int) error {
	if idx < 0 || idx >= len(t.Data) {
		return fmt.Errorf("Task with number {%d} not exists", idx)
	}

	task := &t.Data[idx]
	task.Uncomplete()
	t.Sort()

	return nil
}

func (t *TaskContainer) Remove(idx int) error {
	if idx < 0 || idx >= len(t.Data) {
		return fmt.Errorf("Task with number {%d} not exists", idx)
	}

	t.Data = append(t.Data[:idx], t.Data[idx+1:]...)

	return nil
}

func (t *TaskContainer) Purge() {
	new := []Task{}

	for _, task := range t.Data {
		if !task.IsDone {
			new = append(new, task)
		}
	}

	t.Data = new
}

func (t *TaskContainer) Clean() {
	t.Init()
	t.Save()
}

func (t *TaskContainer) Sort() {
	sort.Slice(t.Data, func(i, j int) bool {
		a, b := t.Data[i], t.Data[j]

		switch {
		case a.IsDone != b.IsDone:
			return a.IsDone
		case a.CreatedAt != b.CreatedAt:
			return a.CreatedAt.Before(b.CreatedAt)
		default:
			return true
		}
	})
}
