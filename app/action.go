package hh

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

var ActionWriter = os.Stdout

func Action(c *cli.Context) {
	fmt.Fprintln(ActionWriter, "Hello Healthy Hacker!")
}
