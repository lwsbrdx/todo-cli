package commands

import (
	"fmt"
	"todo/models"
	"todo/renderers"

	"github.com/urfave/cli/v2"
)

var SearchTasks = cli.Command{
	Name:    "search",
	Aliases: []string{"s"},
	Usage:   "Search tasks",
	Action:  searchTasks,
}

func searchTasks(cCtx *cli.Context) error {
	searchedName := cCtx.Args().First()

	var tasks []models.Task
	if err := DbService.Db.Where("name LIKE ?", "%"+searchedName+"%").Find(&tasks).Error; err != nil {
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
