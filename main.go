package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func Commander() []cli.Command {
	return []cli.Command{
		{
			Name:        "krisz",
			Description: "testcommand",
			Usage:       "invoke the command and see what happens",
			Action:      MyCommand,
			Flags:       []cli.Flag{FlName},
			BashComplete: func(c *cli.Context) {
				printFlagCompletion(FlName)
			},
		},
	}
}

type StringFlag struct {
	cli.StringFlag
}

var FlName = StringFlag{
	StringFlag: cli.StringFlag{
		Name:  "name",
		Usage: "name of resource",
	},
}

func printFlagCompletion(f cli.Flag) {
	fmt.Printf("--%s\n", f.GetName())
}

type greeting string

func (g greeting) Greet() {
	fmt.Println("Hello Universe")
}

var Greeter greeting

func main() {
	app := cli.NewApp()
	app.Name = "dlm"
	app.HelpName = "Hortonworks Data Lifecycle command line tool"
	app.Author = "Hortonworks"
	app.EnableBashCompletion = true
	app.Commands = Commander()

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
