package commands

import (
	"errors"
	"fmt"
	"strconv"
	"todo/models"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var Delete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a task",
	Aliases: []string{"d"},
	Action:  delete,
}

func delete(cCtx *cli.Context) error {
	id := cCtx.Args().First()
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	if idInt == 0 {
		return errors.New("task ID is required")
	}

	task := &models.Task{
		Model: gorm.Model{
			ID: uint(idInt),
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
