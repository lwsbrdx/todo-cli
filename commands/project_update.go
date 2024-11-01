package commands

import (
	"errors"
	"strconv"
	"todo/flags"
	"todo/models"
	"todo/services"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var UpdateProject = cli.Command{
	Name:    "update",
	Usage:   "Update a project",
	Aliases: []string{"u"},
	Flags: []cli.Flag{
		flags.NewProjectNameFlag(false),
	},
	Action: updateProject,
}

func updateProject(cCtx *cli.Context) error {
	projectID, err := strconv.Atoi(cCtx.Args().Get(0))
	if err != nil || projectID == 0 {
		return errors.New("invalid project ID")
	}

	projectName := cCtx.String("name")
	if projectName == "" {
		return errors.New("name is required")
	}

	project := models.Project{
		Model: gorm.Model{ID: uint(projectID)},
		Name:  projectName,
	}

	if err := services.DbInstance.Db.
		Save(&project).
		Error; err != nil {
		return err
	}

	return listProjects(cCtx)
}
