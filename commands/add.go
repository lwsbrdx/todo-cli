package commands

import (
	"todo/flags"
	"todo/helpers"
	"todo/models"

	"github.com/urfave/cli/v2"
)

var Add = cli.Command{
	Name:    "add",
	Usage:   "Add a task",
	Aliases: []string{"a"},
	Flags: []cli.Flag{
		flags.NewTaskNameFlag(true),
		flags.NewTaskStatusFlag(false),
	},
	Action: add,
}

func add(cCtx *cli.Context) error {
	name := cCtx.String("name")
	status := cCtx.Generic("status").(*flags.TaskStatusFlagValue).Status

	if status == models.StatusEmpty {
		status = models.Todo
	}

	task := models.Task{Name: name, Status: status}
	if err := DbService.Db.Create(&task).Error; err != nil {
		return err
	}

	tr := helpers.TasksRenderer{Tasks: []models.Task{task}}
	tr.Render()

	return nil
}
