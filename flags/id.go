package flags

import (
	"github.com/urfave/cli/v2"
)

func NewTaskIDFlag(required bool) cli.Flag {
	return &cli.IntFlag{
		Name:     "id",
		Aliases:  []string{"i"},
		Usage:    "ID of the task",
		Required: required,
	}
}
