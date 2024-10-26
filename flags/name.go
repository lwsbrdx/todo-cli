package flags

import "github.com/urfave/cli/v2"

func NewTaskNameFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     "name",
		Aliases:  []string{"n"},
		Usage:    "Name of the task",
		Required: required,
	}
}

func NewProjectNameFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     "name",
		Aliases:  []string{"n"},
		Usage:    "Name of the project",
		Required: required,
	}
}
