package main

import (
	"log"
	"os"

	"github.com/byroni/rufous/awshelper"
	"github.com/urfave/cli/v2"
)

func main() {
	sess := awshelper.InitSessionCredentials()

	_, err := sess.Config.Credentials.Get()
	if err != nil {
		red.Printf("Please make sure to create your AWS credentials file in ~/.aws/credentials. Visit ")
		cyan.Printf("https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials")
		red.Printf(" for more information")

		log.Fatal(err)
	}

	app := &cli.App{
		Name:  "Rufous",
		Usage: "Lambda based job scheduler",
		Commands: []*cli.Command{
			{
				Name:  "setup",
				Usage: "[COMING SOON] RUN THIS for first time setup",
				Action: func(c *cli.Context) error {
					cyan.Println("Running first time setup...")
					return nil
				},
			},
			{
				Name:  "init",
				Usage: "[COMING SOON] Generate a new job project with a given template (DEFAULT: node)",
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
				Name:  "upload",
				Usage: "[COMING SOON] Upload a job as an AWS lambda function",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "Choose a name for your lambda function. (DEFAULT <folder name>)",
					},
				},
				Action: func(c *cli.Context) error {
					cyan.Println("Uploading project folder to lambda")
					return nil
				},
			},
			{
				Name:  "list",
				Usage: "[COMING SOON] See currently scheduled jobs",
				Action: func(c *cli.Context) error {
					cyan.Println("See currently scheduled jobs")
					return nil
				},
			},
			{
				Name:  "pause",
				Usage: "[COMING SOON] Pause a currently scheduled job",
				Action: func(c *cli.Context) error {
					cyan.Println("Pausing job")
					return nil
				},
			},
			{
				Name:  "start",
				Usage: "[COMING SOON] Start a currently paused job",
				Action: func(c *cli.Context) error {
					cyan.Println("Starting job")
					return nil
				},
			},
			{
				Name:  "history",
				Usage: "[COMING SOON] See job history",
				Action: func(c *cli.Context) error {
					cyan.Println("Listing job history")
					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "[COMING SOON] Delete a job from Rufous",
				Action: func(c *cli.Context) error {
					red.Println("Warning: This action will delete the job and all associated data & configuration from Rufous. You can re-upload a job anytime.")
					cyan.Println("Deleting job")
					return nil
				},
			},
		},
	}

	app.UsageText = "rufous command <command options> sub-command <sub-command options>"

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
