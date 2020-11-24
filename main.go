package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
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
	app.Version = "1.1.0"

	app.Commands = []cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generate scss files from `assetsType`",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "path, p"},
			},
			Action:  func(c *cli.Context) error {
				create.Start("", c.String("path"), c.Args().First())
				return nil
			},
		},
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "create APB scss files",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "path, p"},
			},
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