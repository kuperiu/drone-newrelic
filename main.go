package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

const apiURL = "https://api.newrelic.com/v2"

var (
	version     = "0.0.0"
	build       = "0"
	buildCommit = "00000"
)

func main() {
	app := cli.NewApp()
	app.Name = "new relic drone plugin"
	app.Usage = "new relic drone plugin"
	app.Action = run
	app.Version = fmt.Sprintf("%s+%s", version, build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "apiKey",
			Usage:    "new relic license key",
			EnvVar:   "NEW_RELIC_LICENSE_KEY",
			Required: true,
		},
		cli.StringFlag{
			Name:     "applicationName",
			Usage:    "application name",
			EnvVar:   "APPLICATION_NAME",
			Required: true,
		},
		cli.StringFlag{
			Name:     "revision",
			Usage:    "git sha",
			EnvVar:   "DRONE_COMMIT_SHA",
			Required: true,
		},
		cli.StringFlag{
			Name:     "changelog",
			Usage:    "change log",
			EnvVar:   "DRONE_COMMIT_MESSAGE",
			Required: true,
		},
		cli.StringFlag{
			Name:     "description",
			Usage:    "description",
			EnvVar:   "DRONE_COMMIT_MESSAGE",
			Required: true,
		},
		cli.StringFlag{
			Name:     "user",
			Usage:    "user",
			EnvVar:   "DRONE_COMMIT_AUTHOR",
			Required: true,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	applicationName := c.String("applicationName")
	apiKey := c.String("apiKey")
	applicationID, err := getApplicationID(applicationName, apiKey)
	if err != nil {
		red := color.New(color.BgRed)
		red.Println(err)
		os.Exit(1)
	}

	revision := c.String("revision")
	changelog := c.String("changelog")
	description := c.String("description")
	user := c.String("user")
	err = recordDeployment(applicationID, revision, changelog, description, user, apiKey)
	if err != nil {
		red := color.New(color.BgRed)
		red.Println(err)
		os.Exit(1)
	}
	return nil
}
