package hh

import "github.com/codegangsta/cli"

func App(args []string) {
	app := cli.NewApp()
	app.Name = "hh"
	app.Usage = "Hacking + Health = Happiness"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) {
		println("Hello Healthy Hacker!")
	}

	app.Run(args)
}
