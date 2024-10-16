package main

import (
	"log"
	"os"
	"todo/commands"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "Todo",
		Commands: []*cli.Command{
			&commands.Add,
			&commands.List,
			&commands.Update,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
