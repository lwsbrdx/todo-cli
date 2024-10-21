package commands

import (
	"errors"
	"fmt"
	"todo/flags"
	"todo/models"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var Delete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a task",
	Aliases: []string{"d"},
	Action:  delete,
	Flags: []cli.Flag{
		flags.NewTaskIDFlag(true),
	},
}

func delete(cCtx *cli.Context) error {
	id := cCtx.Int("id")

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

	if err := DbService.Db.Unscoped().Delete(task).Error; err != nil {
		return err
	}

	fmt.Println("Task successfully deleted")
	return list(cCtx)
}
