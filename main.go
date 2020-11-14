package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Rufous",
		Usage: "Lambda based job scheduler",
		Commands: []*cli.Command{
			{
				Name: "init",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "template",
						Aliases: []string{"t"},
						Value:   "node",
						Usage:   "Choose your language template",
					},
				},
				Action: func(c *cli.Context) error {
					err := initProject(c)
					if err != nil {
						log.Fatal(err)
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
