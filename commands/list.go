package commands

import (
	"fmt"
	"todo/models"

	"github.com/urfave/cli/v2"
)

func List(_ *cli.Context) error {
	var tasks []models.Task

	if err := DbService.Db.Find(&tasks).Error; err != nil {
		return err
	}

	for _, task := range tasks {
		fmt.Println(task.String())
	}

	return nil
}
