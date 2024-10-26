package commands

import (
	"fmt"
	"todo/models"
	"todo/renderers"
	"todo/services"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm/clause"
)

var ListTasks = cli.Command{
	Name:    "list",
	Usage:   "List tasks",
	Aliases: []string{"l"},
	Action:  listTasks,
}

func listTasks(_ *cli.Context) error {
	var tasks []models.Task

	if err := services.DbInstance.Db.
		Preload(clause.Associations).
		Find(&tasks).Error; err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	tr := renderers.TasksRenderer{Tasks: tasks}
	tr.Render()

	return nil
}
