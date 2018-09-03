package main

import (
	"fmt"
	"github.com/urfave/cli"
)

func MyCommand(c *cli.Context) {
	fmt.Println("invoked: " + c.Command.Name)
}
