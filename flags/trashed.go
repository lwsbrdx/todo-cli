package flags

import "github.com/urfave/cli/v2"

func NewTrashedFlag() cli.Flag {
	return &cli.BoolFlag{
		Name:     "trashed",
		Aliases:  []string{"t"},
		Usage:    "Show trashed tasks",
		Required: false,
	}
}
