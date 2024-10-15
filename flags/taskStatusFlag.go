package flags

import (
	"errors"
	"strings"
	"todo/models"
)

// Parse une chaîne en TaskStatus
func ParseTaskStatusFromString(s string) (models.TaskStatus, error) {
	switch strings.ToLower(s) {
	case "todo":
		return models.Todo, nil
	case "wip":
		return models.Wip, nil
	case "done":
		return models.Done, nil
	}
	return 0, errors.New("invalid task status")
}

type TaskStatusFlag struct {
	Status models.TaskStatus
}

// Implementation de cli.Generic
func (t *TaskStatusFlag) Set(value string) error {
	status, err := ParseTaskStatusFromString(value)

	if err != nil {
		t.Status = models.Todo
		return nil
	}

	t.Status = status
	return nil
}

func (t *TaskStatusFlag) String() string {
	return t.Status.String()
}
