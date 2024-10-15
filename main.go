package main

import (
	"log"
	"os"
	"todo/commands"
	"todo/flags"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "Todo",
		Commands: []*cli.Command{
			{
				Name:    "add",
				Usage:   "Add a task",
				Aliases: []string{"a"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Aliases:  []string{"n"},
						Usage:    "Name of the task",
						Required: true,
					},
					&cli.GenericFlag{
						// TODO: Voir pour l'autocomplete
						Name:  "status",
						Usage: "Set the task status (`todo`, `wip`, `done`)",
						Value: &flags.TaskStatusFlag{},
					},
				},
				Action: commands.Add,
			},
			{
				Name:    "list",
				Usage:   "List tasks",
				Aliases: []string{"l"},
				Action:  commands.List,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
