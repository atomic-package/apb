package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/atomic-package/apb/lib/create"
)

// These variables are set in build step
var (
	Version  = "unset"
	Revision = "unset"
)

func main() {
	app := cli.NewApp()
	app.Name = "apb-cli"
	app.Usage = "Command line tool for developing APB CSS (Atomic Parts Base CSS)"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name: "path",
			Aliases: []string{"p"},
			Usage: "Load configuration",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generate scss files from `assetsType`",
			Action:  func(c *cli.Context) error {
				create.Start("", c.String("path"), c.Args().First())
				return nil
			},
		},
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "create APB scss files",
			Action:  func(c *cli.Context) error {
				create.Start(c.Args().First(), c.String("path"), "")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}