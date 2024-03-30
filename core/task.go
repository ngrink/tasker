package core

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	UUID        uuid.UUID
	Description string
	Scope       string
	IsDone      bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type CreateTaskDto struct {
	Description string
	Scope       string
}

func NewTask(data CreateTaskDto) Task {
	task := Task{
		UUID:        uuid.New(),
		Description: data.Description,
		Scope:       data.Scope,
		IsDone:      false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	return task
}

func (t *Task) Edit(description string) {
	t.Description = description
}

func (t *Task) Complete() {
	t.IsDone = true
	t.CompletedAt = time.Now()
}

func (t *Task) Uncomplete() {
	t.IsDone = false
	t.CompletedAt = time.Time{}
}

func (t *Task) GetAge() string {
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

func (t *Task) GetStatusMark() string {
	if t.IsDone {
		return "✅"
	} else {
		return ""
	}
}

func (t *Task) GetDesriptionWithStatus() string {
	if t.IsDone {
		return "✅ " + t.Description
	} else {
		return t.Description
	}
}
