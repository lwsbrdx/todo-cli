package commands

import (
	"github.com/urfave/cli/v2"
	"log"
)

func List(_ *cli.Context) error {
	log.Println("list command")
	return nil
}
