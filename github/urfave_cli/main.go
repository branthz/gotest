package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli"
)

func setValue(v string) error {
	fmt.Println("set ok", v)
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "s",
			Value: "localhost",
			Usage: "server address",
		},
	}
	app.HideHelp = true

	app.Action = func(c *cli.Context) error {
		fmt.Println("---------\n")
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name: "set",
			//Aliases: []string{"h"},
			Usage: "set value",
			Action: func(c *cli.Context) error {
				return setValue(c.Args().First())
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		return
	}
	time.Sleep(1e9 * 10)
}
