package hh

import (
	"fmt"
	"github.com/codegangsta/cli"
	"io"
	"os"
)

var ActionWriter io.Writer = os.Stdout

func Action(c *cli.Context) {
	fmt.Fprintln(ActionWriter, "Hello Healthy Hacker!")
}
