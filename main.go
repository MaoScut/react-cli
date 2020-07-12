package main

import (
	"fmt"
	"github.com/madp/react-cli/components"
	"github.com/madp/react-cli/utils"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/user"
)

func main() {
	var err error
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	templatesDir := user.HomeDir + "/" + ".react-cli/templates"
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
							var err error
							dir, fileName := utils.GetDirAndFileName(c.Args().First())
							err = components.GenerateReactFc(
								templatesDir,
								dir,
								fileName,
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

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
