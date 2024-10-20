package commands

import (
	"fmt"
	"todo/models"

	"github.com/urfave/cli/v2"
)

var Search = cli.Command{
	Name:    "search",
	Aliases: []string{"s"},
	Usage:   "Search tasks",
	Action:  search,
}

func search(cCtx *cli.Context) error {
	searchedName := cCtx.Args().First()

	var tasks []models.Task
	if err := DbService.Db.Where("name LIKE ?", "%"+searchedName+"%").Find(&tasks).Error; err != nil {
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
