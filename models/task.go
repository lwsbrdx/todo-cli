package models

import (
	"fmt"

	"gorm.io/gorm"
)

type TaskStatus int

func (t TaskStatus) String() string {
	switch t {
	case StatusEmpty:
		return "Empty"
	case Todo:
		return "To do"
	case Wip:
		return "In progress"
	case Done:
		return "Done"
	}
	return "Unknown"
}

const (
	StatusEmpty TaskStatus = iota
	Todo
	Wip
	Done
)

type Task struct {
	gorm.Model
	Name   string
	Status TaskStatus
}

func (t *Task) String() string {
	task2String := fmt.Sprintf("%d\t%s\t%s", t.ID, t.Status, t.Name)
	return task2String
}
