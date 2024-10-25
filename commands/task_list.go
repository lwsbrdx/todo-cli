package commands

import (
	"fmt"
	"todo/helpers"
	"todo/models"

	"github.com/urfave/cli/v2"
)

var ListTasks = cli.Command{
	Name:    "list",
	Usage:   "List tasks",
	Aliases: []string{"l"},
	Action:  listTasks,
}

func listTasks(_ *cli.Context) error {
	var tasks []models.Task

	if err := DbService.Db.Find(&tasks).Error; err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	tr := helpers.TasksRenderer{Tasks: tasks}
	tr.Render()

	return nil
}
