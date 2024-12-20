package main

import (
	"fmt"
	"os"
	"todo/commands"

	"github.com/urfave/cli/v2"
)

/**
 * Known issues :
 *	- The arguments of a command has to be after the flags
**/
func main() {
	app := &cli.App{
		Name:    "Todo",
		Version: "1.0.0",
		Action:  commands.GlobalCommand.Action,
		Flags:   commands.GlobalCommand.Flags,
		Commands: []*cli.Command{
			{
				Name:   "task",
				Action: commands.ListTasks.Action,
				Subcommands: []*cli.Command{
					&commands.ListTasks,
					&commands.AddTask,
					&commands.UpdateTask,
					&commands.DeleteTask,
					&commands.SearchTasks,
				},
			},
			{
				Name:   "project",
				Action: commands.ListProjects.Action,
				Subcommands: []*cli.Command{
					&commands.ListProjects,
					&commands.AddProject,
					&commands.UpdateProject,
					&commands.DeleteProject,
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
