package main

import (
	"os"
	"io/ioutil"
	"github.com/saromanov/selitra/backend/internal/app"
	"github.com/saromanov/selitra/backend/server"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
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

// parseConfig provides parsing of the config .yml file
func parseConfig(path string) (*alerting.Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %v", err)
	}
	var c *alerting.Config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, fmt.Errorf("unable to parse .born.yml: %v", err)
	}

	return c, nil
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
