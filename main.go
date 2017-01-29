package main

import (
	"fmt"
	"log"
	"os"

	"github.com/coreos/go-semver/semver"
	"github.com/urfave/cli"
	De "github.com/visionmedia/go-debug"
)

var debug = De.Debug("meshchain-util:main")

func main() {
	app := cli.NewApp()
	app.Name = "meshchain-util"
	app.Version = version()
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "identity-filepath, i",
			EnvVar: "MESHCHAIN_UTIL_IDENTITY_FILEPATH",
			Usage:  "Path to an RSA pem file",
			Value:  "~/.ssh/id_rsa",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "Create a new meshchain record",
			Action:  Create,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "data-filepath, d",
					EnvVar: "MESHCHAIN_UTIL_DATA_FILEPATH",
					Usage:  "Path to a data file",
					Value:  "~/.ssh/id_rsa",
				},
			},
		},
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "Print the current version",
			Action:  cli.ShowVersion,
		},
	}

	app.Run(os.Args)
}

func run(context *cli.Context) {
	cli.ShowAppHelp(context)
	os.Exit(1)
}

func version() string {
	version, err := semver.NewVersion(VERSION)
	if err != nil {
		errorMessage := fmt.Sprintf("Error with version number: %v", VERSION)
		log.Panicln(errorMessage, err.Error())
	}
	return version.String()
}
