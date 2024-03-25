package core

import (
	"fmt"
)

var (
	Tasks    TaskContainer
	filepath = "data/tasks.json"
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

func (t *TaskContainer) Complete(idx int) error {
	if idx < 0 || idx >= len(t.Data) {
		return fmt.Errorf("Task with number {%d} not exists", idx)
	}

	task := t.Data[idx]
	task.Complete()

	return nil
}

func (t *TaskContainer) Uncomplete(idx int) error {
	if idx < 0 || idx >= len(t.Data) {
		return fmt.Errorf("Task with number {%d} not exists", idx)
	}

	task := t.Data[idx]
	task.Uncomplete()

	return nil
}

func (t *TaskContainer) Remove(idx int) error {
	if idx < 0 || idx >= len(t.Data) {
		return fmt.Errorf("Task with number {%d} not exists", idx)
	}

	t.Data = append(t.Data[:idx], t.Data[idx+1:]...)

	return nil
}

func (t *TaskContainer) Clean() {
	t.Init()
	t.Save()
}
