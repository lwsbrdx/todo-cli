package commands

import (
	"fmt"
	"sort"
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

	sort.SliceStable(tasks, func(i, j int) bool {
		return tasks[i].Status < tasks[j].Status
	})

	tr := renderers.TasksRenderer{Tasks: tasks}
	tr.Render()

	return nil
}
