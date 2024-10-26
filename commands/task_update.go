package commands

import (
	"errors"
	"strconv"
	"todo/flags"
	"todo/models"
	"todo/renderers"
	"todo/services"

	"github.com/urfave/cli/v2"
)

var UpdateTask = cli.Command{
	Name:    "update",
	Usage:   "Update a task",
	Aliases: []string{"u"},
	Flags: []cli.Flag{
		flags.NewTaskNameFlag(false),
		flags.NewTaskStatusFlag(false),
	},
	Action: updateTask,
}

func updateTask(cCtx *cli.Context) error {
	id, err := strconv.Atoi(cCtx.Args().First())
	if err != nil {
		return errors.New("invalid task ID")
	}

	// Récupérer la task correspondante à l'ID
	var task models.Task
	if err := services.DbInstance.Db.First(&task, id).Error; err != nil {
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
	if err := services.DbInstance.Db.Save(&task).Error; err != nil {
		return err
	}

	tr := renderers.TasksRenderer{Tasks: []models.Task{task}}
	tr.Render()

	return nil
}
