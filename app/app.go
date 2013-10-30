package hh

import "github.com/codegangsta/cli"

const (
	NAME    = "hh"
	USAGE   = "Hacking + Health = Happiness"
	VERSION = "0.0.1"
)

func App(args []string) {
	app := cli.NewApp()
	app.Name = NAME
	app.Usage = USAGE
	app.Version = VERSION
	app.Action = func(c *cli.Context) {
		println(Action(c))
	}
	app.Run(args)
}
