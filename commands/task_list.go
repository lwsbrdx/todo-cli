package commands

import (
	"errors"
	"fmt"
	"sort"
	"todo/flags"
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
	Flags: []cli.Flag{
		flags.NewTrashedFlag(),
		flags.NewShowAllFlag(),
	},
	Action: listTasks,
}

func listTasks(cCtx *cli.Context) error {
	showTrashed := cCtx.Bool("trashed")
	showAll := cCtx.Bool("show-all")

	if showTrashed {
		return listTrashedTasks(cCtx)
	}

	request := services.DbInstance.Db.Preload(clause.Associations)
	if !showAll {
		request = request.Where(
			"status IN ?",
			[]models.TaskStatus{
				models.Todo,
				models.Wip,
			},
		)
	}

	var tasks []models.Task
	if err := request.Find(&tasks).Error; err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	sort.SliceStable(tasks, func(i, j int) bool {
		if tasks[i].Status == models.Wip {
			return true
		}

		return tasks[i].Status < tasks[j].Status
	})

	tr := renderers.TasksRenderer{Tasks: tasks}
	tr.Render()

	return nil
}

func listTrashedTasks(_ *cli.Context) error {
	var tasks []models.Task
	if err := services.DbInstance.Db.
		Unscoped().
		Preload(clause.Associations).
		Where("deleted_at IS NOT NULL").
		Find(&tasks).Error; err != nil {
		return err
	}

	if len(tasks) == 0 {
		return errors.New("Trash is empty")
	}

	tr := renderers.TasksRenderer{Tasks: tasks}
	tr.Render()
	return nil
}
