package flags

import "github.com/urfave/cli/v2"

func NewEmptyTrashFlag() cli.Flag {
	return &cli.BoolFlag{
		Name:     "empty-trash",
		Usage:    "Empty trashed tasks",
		Required: false,
	}
}
