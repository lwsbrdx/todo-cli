package commands

import (
	"todo/flags"
	"todo/models"
	"todo/services"

	"github.com/urfave/cli/v2"
)

var GlobalCommand = cli.Command{
	Name:   "global",
	Action: global,
	Flags: []cli.Flag{
		flags.NewTrashedFlag(),
		flags.NewEmptyTrashFlag(),
	},
}

func global(cCtx *cli.Context) error {
	shouldEmptyTrash := cCtx.Bool("empty-trash")

	if shouldEmptyTrash {
		if err := services.DbInstance.Db.
			Unscoped().
			Where("deleted_at IS NOT NULL").
			Delete(&models.Task{}).Error; err != nil {
			return err
		}
	}

	if shouldEmptyTrash {
		cCtx.Set("trashed", "true")
	}

	return listTasks(cCtx)
}
