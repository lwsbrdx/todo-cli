package commands

import (
	"fmt"
	"todo/flags"
	"todo/models"
	"todo/services"

	"github.com/urfave/cli/v2"
)

var AddProject = cli.Command{
	Name:    "add",
	Usage:   "Add a project",
	Aliases: []string{"a"},
	Action:  addProject,
	Flags: []cli.Flag{
		flags.NewProjectNameFlag(true),
	},
}

func addProject(cCtx *cli.Context) error {
	name := cCtx.String("name")

	project := models.Project{Name: name}
	if err := services.DbInstance.Db.Create(&project).Error; err != nil {
		return err
	}

	fmt.Printf("Project %s successfully created\n", name)
	return nil
}
