package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	defaultPath := fmt.Sprintf("%s/.config/bliz", homeDir)

	bliz := NewBliz(defaultPath)
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "get",
				Usage: "get the value by key eg. get key",
				Action: func(c *cli.Context) error {
					if c.Args().First() == "" {
						return &KeyRequiredError{}
					}
					value := bliz.Get(c.Args().First())
					fmt.Println(value)
					return nil
				},
			},
			{
				Name:  "set",
				Usage: "set the value by key eg. set key value",
				Action: func(c *cli.Context) error {
					if c.Args().First() == "" {
						return &KeyRequiredError{}
					}
					if c.Args().Get(1) == "" {
						return &ValueRequiredError{}
					}
					bliz.Set(c.Args().First(), c.Args().Get(1))
					fmt.Println("value set")
					return nil
				},
			},
			{
				Name:  "list",
				Usage: "list all the keys, might be deleted in the future",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

//KeyRequiredError return error message for key validation
type KeyRequiredError struct{}

func (m *KeyRequiredError) Error() string {
	return "key is required"
}

//ValueRequiredError return error message for value validation
type ValueRequiredError struct{}

func (m *ValueRequiredError) Error() string {
	return "value is required"
}
