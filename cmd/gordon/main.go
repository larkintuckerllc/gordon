package main

import (
	"log"
	"os"

	"github.com/larkintuckerllc/gordon/internal/gordon"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "gordon",
		Usage: "manages DNS A records",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "dns",
				Usage:    "dns name",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "project",
				Usage:    "project name",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "zone",
				Usage:    "zone name",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			dnsName := c.String("dns")
			projectId := c.String("project")
			zoneName := c.String("zone")
			err := gordon.Execute(projectId, zoneName, dnsName)
			return err
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
