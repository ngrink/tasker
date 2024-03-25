package core

import "time"

type Task struct {
	Description string
	Category    string
	IsDone      bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type CreateTaskDto struct {
	Description string
	Category    string
}

func NewTask(data CreateTaskDto) Task {
	task := Task{
		Description: data.Description,
		Category:    data.Category,
		IsDone:      false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	return task
}

func (t *Task) Complete() {
	t.IsDone = true
	t.CompletedAt = time.Now()
}

func (t *Task) Uncomplete() {
	t.IsDone = false
	t.CompletedAt = time.Time{}
}
