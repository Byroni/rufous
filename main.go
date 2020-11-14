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
				Name:  "init",
				Usage: "Generate a new job project with a given template (DEFAULT: node)",
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
			{
				Name:  "setup",
				Usage: "RUN THIS for first time setup",
				Action: func(c *cli.Context) error {
					cyan.Println("Running first time setup...")
					return nil
				},
			},
		},
	}

	app.UsageText = "rufous command <command options> sub-command <sub-command options>"

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
