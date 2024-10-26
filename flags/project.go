package flags

import (
	"errors"
	"fmt"
	"strconv"
	"todo/models"
	"todo/services"

	"github.com/urfave/cli/v2"
)

type ProjectIDFlagValue struct {
	Project models.Project
}

func NewProjectIDFlag(required bool) cli.Flag {
	return &cli.GenericFlag{
		Name:     "project",
		Aliases:  []string{"p"},
		Usage:    "ID of the project",
		Value:    &ProjectIDFlagValue{},
		Required: required,
	}
}

func (p *ProjectIDFlagValue) Set(value string) error {
	id, err := strconv.Atoi(value)
	if err != nil {
		return errors.New("invalid project ID")
	}

	var project models.Project
	if err := services.DbInstance.Db.First(&project, id).Error; err != nil {
		return err
	}

	p.Project = project

	return nil
}

func (p *ProjectIDFlagValue) String() string {
	return fmt.Sprintf("%d\t%s", p.Project.ID, p.Project.Name)
}
