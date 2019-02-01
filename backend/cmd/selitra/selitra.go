package main

import (
	"os"

	"github.com/saromanov/selitra/backend/internal/app"
	"github.com/saromanov/selitra/backend/server"
	"github.com/urfave/cli"
)

// setupServer provides initialization of server
func setupServer(a *app.App, c *cli.Context) {
	server.Create(app, &server.Config{
		Address: c.String("selitra-host"),
	})
}

func run(c *cli.Context) {
	app := app.New()
	setupServer(app, c)
}

func main() {
	app := cli.NewApp()
	app.Name = "selitra"
	app.Usage = "log processing tool"
	app.Commands = []cli.Command{
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "path to .yml config",
			Action: func(c *cli.Context) error {
				configPath := c.Args().First()
				config, err := parseConfig(configPath)
				if err != nil {
					panic(err)
				}
				if err := setupApp(config); err != nil {
					panic(err)
				}
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
