package commands

import (
	"todo/flags"
	"todo/models"

	"github.com/urfave/cli/v2"
)

func Add(cCtx *cli.Context) error {

	name := cCtx.String("name")
	status := cCtx.Generic("status").(*flags.TaskStatusFlag)

	task := models.Task{Name: name, Status: status.Status}
	if err := DbService.Db.Create(&task).Error; err != nil {
		return err
	}

	return nil
}
