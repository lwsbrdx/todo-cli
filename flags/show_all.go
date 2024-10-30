package flags

import "github.com/urfave/cli/v2"

func NewShowAllFlag() cli.Flag {
	return &cli.BoolFlag{
		Name:     "show-all",
		Aliases:  []string{"a"},
		Usage:    "Show all tasks",
		Required: false,
	}
}
