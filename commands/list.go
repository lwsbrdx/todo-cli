package commands

import (
	"fmt"
	"todo/models"

	"github.com/urfave/cli/v2"
)

var List = cli.Command{
	Name:    "list",
	Usage:   "List tasks",
	Aliases: []string{"l"},
	Action:  list,
}

func list(_ *cli.Context) error {
	var tasks []models.Task

	if err := DbService.Db.Find(&tasks).Error; err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	for _, task := range tasks {
		fmt.Println(task.String())
	}

	return nil
}
