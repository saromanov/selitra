package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/saromanov/selitra/backend/internal/app"
	structs "github.com/saromanov/selitra/backend/internal/structs/v1"
	"github.com/saromanov/selitra/backend/server"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

// setupServer provides initialization of server
func setupServer(a *app.App, c *structs.Config) {
	server.Create(a, &server.Config{
		Address: c.Address,
	})
}

func run(c *structs.Config) error {
	app, err := app.New(c)
	if err != nil {
		return fmt.Errorf("unable to init app: %v", err)
	}
	setupServer(app, c)
	return nil
}

// parseConfig provides parsing of the config .yml file
func parseConfig(path string) (*structs.Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %v", err)
	}
	var c *structs.Config
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
				if err := run(config); err != nil {
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
