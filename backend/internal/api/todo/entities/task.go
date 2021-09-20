package entities

import (
	"errors"
	"fmt"
	"time"
)

type Task struct {
	Id          string
	Name        string
	Description string
	UpdatedAt   time.Time
}

func (task Task) Equal(dst Task) bool {
	return task.Id == dst.Id &&
		task.Name == dst.Name &&
		task.Description == dst.Description &&
		task.UpdatedAt.Equal(dst.UpdatedAt)
}

func (task Task) Validate() error {
	if len(task.Id) < 1 {
		return errors.New(fmt.Sprintf("invalid Id length %d", len(task.Id)))
	}
	if len(task.Name) < 1 || 50 < len(task.Name) {
		return errors.New(fmt.Sprintf("invalid Name length %d", len(task.Name)))
	}
	if 1000 < len(task.Description) {
		return errors.New(fmt.Sprintf("invalid Description length %d", len(task.Description)))
	}
	if (time.Time{}).Equal(task.UpdatedAt) || task.UpdatedAt.After(time.Now()) {
		return errors.New(fmt.Sprintf("invalid UpdatedAt %+v", task.UpdatedAt))
	}
	return nil
}
