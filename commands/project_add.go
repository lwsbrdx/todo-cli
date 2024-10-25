package commands

import (
	"log"
	"todo/flags"
	"todo/models"

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
	if err := DbService.Db.Create(&project).Error; err != nil {
		return err
	}

	log.Printf("Project %s successfully created", name)
	return nil
}
