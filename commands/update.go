package commands

import (
	"todo/flags"
	"todo/helpers"
	"todo/models"

	"github.com/urfave/cli/v2"
)

var Update = cli.Command{
	Name:    "update",
	Usage:   "Update a task",
	Aliases: []string{"u"},
	Flags: []cli.Flag{
		flags.NewTaskIDFlag(true),
		flags.NewTaskNameFlag(false),
		flags.NewTaskStatusFlag(false),
	},
	Action: update,
}

func update(cCtx *cli.Context) error {
	id := cCtx.Int("id")

	// Récupérer la task correspondante à l'ID
	var task models.Task
	if err := DbService.Db.First(&task, id).Error; err != nil {
		return err
	}

	// Mettre à jour les informations de la task
	updatedTaskName := cCtx.String("name")
	if updatedTaskName != "" {
		task.Name = updatedTaskName
	}

	updatedTaskStatus := cCtx.Generic("status").(*flags.TaskStatusFlagValue).Status
	if updatedTaskStatus != models.StatusEmpty && updatedTaskStatus != task.Status {
		task.Status = updatedTaskStatus
	}

	// Sauvegarder les modifications dans la base de données
	if err := DbService.Db.Save(&task).Error; err != nil {
		return err
	}

	tr := helpers.TasksRenderer{Tasks: []models.Task{task}}
	tr.Render()

	return nil
}
