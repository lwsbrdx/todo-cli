package commands

import (
	"errors"
	"strconv"
	"todo/models"
	"todo/services"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm/clause"
)

var DeleteProject = cli.Command{
	Name:    "delete",
	Usage:   "Delete a project",
	Aliases: []string{"d"},
	Action:  deleteProject,
}

func deleteProject(cCtx *cli.Context) error {
	id, err := strconv.Atoi(cCtx.Args().First())

	if err != nil {
		return err
	}

	if id == 0 {
		return errors.New("project ID is required")
	}

	var project models.Project
	if err := services.DbInstance.Db.
		Preload(clause.Associations).
		First(&project, id).
		Error; err != nil {
		return err
	}

	for _, task := range project.Tasks {
		if err := services.DbInstance.Db.
			Model(&task).
			Association("Project").
			Delete(&project); err != nil {
			return err
		}
	}

	if err := services.DbInstance.Db.Delete(&project).Error; err != nil {
		return err
	}

	return listProjects(cCtx)
}
