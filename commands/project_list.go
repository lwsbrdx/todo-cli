package commands

import (
	"fmt"
	"todo/flags"
	"todo/models"
	"todo/renderers"
	"todo/services"

	"github.com/urfave/cli/v2"
)

var ListProjects = cli.Command{
	Name:    "list",
	Aliases: []string{"l"},
	Usage:   "List projects",
	Flags: []cli.Flag{
		flags.NewTrashedFlag(),
	},
	Action: listProjects,
}

func listProjects(cCtx *cli.Context) error {
	if cCtx.Bool("trashed") {
		return listTrashedProjects(cCtx)
	}

	projects := []models.Project{}

	if err := services.DbInstance.Db.Find(&projects).Error; err != nil {
		return err
	}

	if len(projects) == 0 {
		fmt.Println("No projects found")
		return nil
	}

	pr := renderers.ProjectsRenderer{Projects: projects}
	pr.Render()

	return nil
}

func listTrashedProjects(_ *cli.Context) error {
	var projects []models.Project
	if err := services.DbInstance.Db.
		Unscoped().
		Where("deleted_at IS NOT NULL").
		Find(&projects).Error; err != nil {
		return err
	}

	if len(projects) == 0 {
		fmt.Println("Projects trash is empty")
		return nil
	}

	pr := renderers.ProjectsRenderer{Projects: projects}
	pr.Render()
	return nil
}
