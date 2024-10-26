package models

import "gorm.io/gorm"

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
	ProjectID uint
	Project   Project
	Name      string
	Status    TaskStatus
}
