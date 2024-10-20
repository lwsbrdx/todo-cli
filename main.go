package main

import (
	"fmt"
	"os"
	"todo/commands"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "Todo",
		Action: commands.List.Action,
		Commands: []*cli.Command{
			&commands.Add,
			&commands.List,
			&commands.Update,
			&commands.Delete,
			&commands.Search,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
