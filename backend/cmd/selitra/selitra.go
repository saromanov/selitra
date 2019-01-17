package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var flags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "SELITRA_HOST",
		Name:   "selitra-host",
		Usage:  "host of selitra server",
	},
}

func setupServer(c *cli.Context) {

}
func run(c *cli.Context) {
	server, err := setupServer(c)
	if err != nil {
		panic(err)
	}
	fmt.Println(provider)
}

func main() {
	app := cli.NewApp()
	app.Name = "selitra"
	app.Action = run
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
