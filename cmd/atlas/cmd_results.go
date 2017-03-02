// cmd_results.go

package main

import (
	"github.com/urfave/cli"
)

// init injects our results-related commands (now just an alias for "measurement results")
func init() {
	cliCommands = append(cliCommands, cli.Command{
		Name:        "results",
		Aliases:     []string{"r", "res"},
		Usage:       "results for one measurement",
		Description: "returns results for one measurement group",
		Action:      measurementResults,
	})
}
