package main

import (
	"fmt"
	"os"
	"todo/commands"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "Todo",
		Version: "1.0.0",
		Commands: []*cli.Command{
			{
				Name: "task",
				Subcommands: []*cli.Command{
					&commands.AddTask,
					&commands.ListTasks,
					&commands.UpdateTask,
					&commands.DeleteTask,
					&commands.SearchTasks,
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
