package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/saromanov/selitra/backend/server"
)

var flags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "SELITRA_HOST",
		Name:   "selitra-host",
		Usage:  "host of selitra server",
	},
}

// setupServer provides initialization of server
func setupServer(c *cli.Context) {
	server.Create(&server.Config{
		Address: c.String("selitra-host")
	})
}

func run(c *cli.Context) {
	setupServer(c)
}

func main() {
	app := cli.NewApp()
	app.Name = "selitra"
	app.Action = run
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
