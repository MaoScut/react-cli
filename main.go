package main

import (
	"fmt"
	"github.com/madp/react-cli/components"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const templatePath = "./templates"

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					fmt.Println("added task: ", c.Args().First())
					return nil
				},
			},
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action: func(c *cli.Context) error {
					fmt.Println("completed task: ", c.Args().First())
					return nil
				},
			},
			{
				Name:    "generate",
				Aliases: []string{"g"},
				Usage:   "generate react components",
				Subcommands: []*cli.Command{
					{
						Name:  "fc",
						Usage: "add a new function component",
						Action: func(c *cli.Context) error {
							err := components.GenerateReactFc(
								"./templates/fc.tsx",
								".",
								c.Args().First(),
							)
							return err
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(c *cli.Context) error {
							fmt.Println("removed task template: ", c.Args().First())
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func generateReactFcTemplateStr(templateFilePath string, name string) (string, error) {
	f, err := ioutil.ReadFile(templateFilePath)
	if err != nil {
		fmt.Print(err)
		return "", err
	}
	result := strings.Replace(string(f), "App", name, -1)
	return result, nil
}
