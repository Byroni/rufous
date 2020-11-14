package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func generateTemplateFiles(dir string) {
	fmt.Print("Creating a new job project in ")
	cyan.Println(dir)

	// Generate project files
	index, err := os.Create(dir + "/" + "index.js")
	if err != nil {
		log.Fatal(err)
	}

	_, err = index.WriteString("exports.handler =  async function(event, context) {\n  console.log(\"EVENT: \\n\" + JSON.stringify(event, null, 2))\n  return context.logStreamName\n}")
	if err != nil {
		log.Fatal(err)
	}

	config, err := os.Create(dir + "/" + "rufousConfig.yml")
	if err != nil {
		log.Fatal(err)
	}

	_, err = config.WriteString("language: node\nversion: 12.19.0\ncron: 0 0 0 0 0 0 0")
	if err != nil {
		log.Fatal(err)
	}
}

func initProject(c *cli.Context) error {
	// Project name required
	if c.NArg() == 0 {
		fmt.Println("Please specify job directory")
		cyan.Println("rufous init <project-directory>\n")

		fmt.Println("For example:")
		cyan.Println("rufous init my-project-folder")

		printHelpStatement()

		return nil
	}

	// Get working directory name
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Get project name
	currentWDPaths := []string{
		".",
		"./",
	}

	arg := c.Args().Get(0)

	for _, b := range currentWDPaths {
		if arg == b {
			// Project is cwd, generate files here
			generateTemplateFiles(cwd)
			return nil
		}
	}

	// If folder exists, generate templates in folder
	_, err = os.Stat(arg)
	if os.IsNotExist(err) {
		// Create folder
		err = os.MkdirAll(arg, 0755)
		if err != nil {
			log.Fatalf("Failed to create folder %s in %s", arg, cwd)
		}
		// Generate files in newly created folder
		generateTemplateFiles(cwd + "/" + arg)
	} else {
		// Generate files in existing folder
		generateTemplateFiles(cwd + "/" + arg)
	}

	return nil
}
