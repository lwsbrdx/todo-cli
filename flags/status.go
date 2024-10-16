package flags

import (
	"errors"
	"strings"
	"todo/models"

	"github.com/urfave/cli/v2"
)

func NewTaskStatusFlag(required bool) cli.Flag {
	return &cli.GenericFlag{
		Name:     "status",
		Aliases:  []string{"s"},
		Usage:    "Set the task status (`todo`, `wip`, `done`)",
		Value:    &TaskStatusFlagValue{},
		Required: required,
	}
}

type TaskStatusFlagValue struct {
	Status models.TaskStatus
}

// Implementation de cli.Generic
func (t *TaskStatusFlagValue) Set(value string) error {
	switch strings.ToLower(value) {
	case "todo":
		t.Status = models.Todo
	case "wip":
		t.Status = models.Wip
	case "done":
		t.Status = models.Done
	default:
		return errors.New("invalid task status")
	}

	return nil
}

func (t *TaskStatusFlagValue) String() string {
	return t.Status.String()
}
