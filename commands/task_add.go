package commands

import (
	"todo/flags"
	"todo/models"
	"todo/renderers"
	"todo/services"

	"github.com/urfave/cli/v2"
)

var AddTask = cli.Command{
	Name:    "add",
	Usage:   "Add a task",
	Aliases: []string{"a"},
	Flags: []cli.Flag{
		flags.NewTaskNameFlag(true),
		flags.NewTaskStatusFlag(false),
		flags.NewProjectIDFlag(false),
	},
	Action: addTask,
}

func addTask(cCtx *cli.Context) error {
	name := cCtx.String("name")
	status := cCtx.Generic("status").(*flags.TaskStatusFlagValue).Status
	projectFlag := cCtx.Generic("project").(*flags.ProjectIDFlagValue)

	if status == models.StatusEmpty {
		status = models.Todo
	}

	task := models.Task{
		Name:   name,
		Status: status,
	}
	if err := services.DbInstance.Db.Create(&task).Error; err != nil {
		return err
	}

	err := services.DbInstance.Db.
		Model(&task).
		Association("Project").
		Append(&projectFlag.Project)
	if err != nil {
		return err
	}

	tr := renderers.TasksRenderer{Tasks: []models.Task{task}}
	tr.Render()

	return nil
}
