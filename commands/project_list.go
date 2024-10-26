package commands

import (
	"fmt"
	"todo/models"
	"todo/renderers"
	"todo/services"

	"github.com/urfave/cli/v2"
)

var ListProjects = cli.Command{
	Name:    "list",
	Aliases: []string{"l"},
	Usage:   "List projects",
	Action:  listProjects,
}

func listProjects(cCtx *cli.Context) error {
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
