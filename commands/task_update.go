package commands

import (
	"errors"
	"strconv"
	"todo/flags"
	"todo/models"
	"todo/renderers"
	"todo/services"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm/clause"
)

var UpdateTask = cli.Command{
	Name:    "update",
	Usage:   "Update a task",
	Aliases: []string{"u"},
	Flags: []cli.Flag{
		flags.NewTaskNameFlag(false),
		flags.NewTaskStatusFlag(false),
		flags.NewProjectIDFlag(false),
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
	if err := services.DbInstance.Db.
		Preload(clause.Associations).
		First(&task, id).Error; err != nil {
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

	updatedTaskProject := cCtx.Generic("project").(*flags.ProjectIDFlagValue).Project
	if updatedTaskProject.ID != 0 && updatedTaskProject.ID != task.ProjectID {
		err := services.DbInstance.Db.
			Model(&task).
			Association("Project").
			Append(&updatedTaskProject)
		if err != nil {
			return err
		}
	}

	// Sauvegarder les modifications dans la base de données
	if err := services.DbInstance.Db.Save(&task).Error; err != nil {
		return err
	}

	tr := renderers.TasksRenderer{Tasks: []models.Task{task}}
	tr.Render()

	return nil
}
