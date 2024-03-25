package core

import (
	"fmt"
	"time"
)

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

func (t *Task) getAge() string {
	duration := time.Since(t.CreatedAt)

	days := int64(duration.Hours()) / 24
	hours := int64(duration.Hours()) % 24
	minutes := int64(duration.Minutes()) % 60
	seconds := int64(duration.Seconds()) % 60

	switch {
	case days > 0:
		return fmt.Sprintf("%dd", days)
	case hours > 0:
		return fmt.Sprintf("%dh", hours)
	case minutes > 0:
		return fmt.Sprintf("%dm", minutes)
	default:
		return fmt.Sprintf("%ds", seconds)
	}
}
