package cli

import (
	"fmt"

	"github.com/Distributed-Computing-Denver/lamport/config"
	"github.com/Distributed-Computing-Denver/lamport/node"
	"github.com/urfave/cli"
)

// Run executes the Lamport cli
func Run(args []string) error {
	app := getApp()
	return app.Run(args)
}

func getApp() *cli.App {
	app := cli.NewApp()
	app.Name = "lamport"
	app.Usage = "An academic exercise in building a distributed system"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "run a lamport node",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "config, c",
					Value: "lamport.toml",
					Usage: "lamport configuration `FILE`",
				},
			},
			Action: func(c *cli.Context) error {
				cf := c.String("config")
				return action(cf)
			},
		},
	}
	return app
}

func action(c string) error {
	config, err := config.ReadConfig(c)
	if err != nil {
		return fmt.Errorf("Error processing config file: %s", err)
	}
	return node.Run(node.New(config))
}
