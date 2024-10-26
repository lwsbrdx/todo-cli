package commands

import (
	"errors"
	"fmt"
	"strconv"
	"todo/models"
	"todo/services"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var DeleteTask = cli.Command{
	Name:    "delete",
	Usage:   "Delete a task",
	Aliases: []string{"d"},
	Action:  deleteTask,
}

func deleteTask(cCtx *cli.Context) error {
	ids := cCtx.Args().Slice()
	if len(ids) == 0 {
		return errors.New("task ID is required")
	}

	tasksToDelete := []models.Task{}
	for _, id := range ids {
		taskID, err := strconv.Atoi(id)
		if err != nil {
			return err
		}

		// Check if task exists
		task := models.Task{
			Model: gorm.Model{
				ID: uint(taskID),
			},
		}
		if err := services.DbInstance.Db.First(&task).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New(fmt.Sprintf("task with ID %d not found", taskID))
			}
			return err
		}

		tasksToDelete = append(tasksToDelete, task)
	}

	// Soft deletes the tasks
	if err := services.DbInstance.Db.Delete(tasksToDelete).Error; err != nil {
		return err
	}

	return listTasks(cCtx)
}
