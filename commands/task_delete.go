package commands

import (
	"errors"
	"fmt"
	"strconv"
	"todo/models"

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
	id, err := strconv.Atoi(cCtx.Args().First())
	if err != nil {
		return err
	}

	if id == 0 {
		return errors.New("task ID is required")
	}

	task := &models.Task{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	if err := DbService.Db.First(task).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("task not found")
		}
		return err
	}

	if err := DbService.Db.Delete(task).Error; err != nil {
		return err
	}

	fmt.Println("Task successfully deleted")
	return listTasks(cCtx)
}
