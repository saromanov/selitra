package main

import (
	"os"

	"github.com/saromanov/selitra/backend/internal/app"
	"github.com/saromanov/selitra/backend/server"
	"github.com/urfave/cli"
)

var flags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "SELITRA_HOST",
		Name:   "selitra-host",
		Usage:  "host of selitra server",
	},
}

// setupServer provides initialization of server
func setupServer(a *app.App, c *cli.Context) {
	server.Create(&server.Config{
		App:     a,
		Address: c.String("selitra-host"),
	})
}

func run(c *cli.Context) {
	app := app.New()
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
