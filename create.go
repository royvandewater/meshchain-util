package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

// Create creates a new meshchain record
func Create(context *cli.Context) {

}

func getCreateOpts(context *cli.Context) (string, string) {
	identifyFilepath := context.GlobalString("identity-filepath")
	dataFilepath := context.String("data-filepath")

	if dataFilepath == "" {
		cli.ShowAppHelp(context)

		if dataFilepath == "" {
			color.Red("  Missing required flag --data-filepath or MESHCHAIN_UTIL_DATA_FILEPATH")
		}
		os.Exit(1)
	}

	return identifyFilepath, dataFilepath
}
